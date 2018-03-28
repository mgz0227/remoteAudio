package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	// _ "net/http/pprof"

	"github.com/dh1tw/remoteAudio/audio"
	"github.com/dh1tw/remoteAudio/audio/pbReader"
	"github.com/dh1tw/remoteAudio/audio/pbWriter"
	"github.com/dh1tw/remoteAudio/audio/scReader"
	"github.com/dh1tw/remoteAudio/audio/scWriter"
	"github.com/dh1tw/remoteAudio/audio/wavReader"
	"github.com/dh1tw/remoteAudio/audiocodec/opus"
	"github.com/gordonklaus/portaudio"
	"github.com/nats-io/go-nats"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serverMqttCmd represents the mqtt command
var natsClientCmd = &cobra.Command{
	Use:   "natsclient",
	Short: "nats client",
	Long:  `nats client`,
	Run:   natsAudioClient,
}

func init() {
	serverCmd.AddCommand(natsClientCmd)
	natsClientCmd.Flags().StringP("broker-url", "u", "localhost", "Broker URL")
	natsClientCmd.Flags().IntP("broker-port", "p", 4222, "Broker Port")
	natsClientCmd.Flags().StringP("password", "P", "", "NATS Password")
	natsClientCmd.Flags().StringP("username", "U", "", "NATS Username")
	natsClientCmd.Flags().StringP("radio", "Y", "myradio", "Radio ID")
}

func natsAudioClient(cmd *cobra.Command, args []string) {

	// Try to read config file
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		if strings.Contains(err.Error(), "Not Found in") {
			fmt.Println("no config file found")
		} else {
			fmt.Println("Error parsing config file", viper.ConfigFileUsed())
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// check if values from config file / pflags are valid
	if !checkAudioParameterValues() {
		os.Exit(-1)
	}

	// bind the pflags to viper settings
	viper.BindPFlag("nats.broker-url", cmd.Flags().Lookup("broker-url"))
	viper.BindPFlag("nats.broker-port", cmd.Flags().Lookup("broker-port"))
	viper.BindPFlag("nats.password", cmd.Flags().Lookup("password"))
	viper.BindPFlag("nats.username", cmd.Flags().Lookup("username"))
	viper.BindPFlag("nats.radio", cmd.Flags().Lookup("radio"))

	// profiling server
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	// viper settings need to be copied in local variables
	// since viper lookups allocate of each lookup a copy
	// and are quite inperformant

	audioFramesPerBuffer := viper.GetInt("audio.frame-length")

	oDeviceName := viper.GetString("output-device.device-name")
	oSamplerate := viper.GetFloat64("output-device.samplerate")
	oLatency := viper.GetDuration("output-device.latency")
	oChannels := viper.GetInt("output-device.channels")
	oRingBufferSize := viper.GetInt("audio.rx-buffer-length")

	iDeviceName := viper.GetString("input-device.device-name")
	iSamplerate := viper.GetFloat64("input-device.samplerate")
	iLatency := viper.GetDuration("input-device.latency")
	iChannels := viper.GetInt("input-device.channels")

	opusBitrate := viper.GetInt("opus.bitrate")
	opusComplexity := viper.GetInt("opus.complexity")

	portaudio.Initialize()
	defer portaudio.Terminate()

	// Setup receiving path
	fromRadioSinks, err := audio.NewDefaultRouter()
	if err != nil {
		log.Fatal(err)
	}

	fromRadioSources, err := audio.NewDefaultSelector()
	if err != nil {
		log.Fatal(err)
	}

	// Setup transmitting path
	toRadioSinks, err := audio.NewDefaultRouter()
	if err != nil {
		log.Fatal(err)
	}

	toRadioSources, err := audio.NewDefaultSelector()
	if err != nil {
		log.Fatal(err)
	}

	nc := &natsClient{
		fromRadioSources: fromRadioSources,
		fromRadioSinks:   fromRadioSinks,
		toRadioSources:   toRadioSources,
		toRadioSinks:     toRadioSinks,
	}

	fromRadioAudio, err := scWriter.NewScWriter(
		scWriter.DeviceName(oDeviceName),
		scWriter.Channels(oChannels),
		scWriter.Samplerate(oSamplerate),
		scWriter.Latency(oLatency),
		scWriter.RingBufferSize(oRingBufferSize),
		scWriter.FramesPerBuffer(audioFramesPerBuffer),
	)
	if err != nil {
		log.Fatal(err)
	}

	// wavRec, err := wavWriter.NewWavWriter("test_rec.wav",
	// 	wavWriter.BitDepth(8),
	// 	wavWriter.Channels(1),
	// 	wavWriter.Samplerate(22000))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// nc.txRouter.AddSink("wavFile", wavRec, false)

	wav, err := wavReader.NewWavReader("test.wav")
	if err != nil {
		log.Fatal(err)
	}

	toRadioAudio, err := scReader.NewScReader(
		scReader.Callback(nc.toTxSinksCb),
		scReader.DeviceName(iDeviceName),
		scReader.Channels(iChannels),
		scReader.Samplerate(iSamplerate),
		scReader.Latency(iLatency),
		scReader.FramesPerBuffer(audioFramesPerBuffer),
	)
	if err != nil {
		log.Fatal(err)
	}

	fromNetwork, err := pbReader.NewPbReader()
	if err != nil {
		log.Fatal(err)
	}

	natsc, err := nats.Connect("nats://195.201.117.206:4222")
	if err != nil {
		log.Fatal(err)
	}

	// subscribe to the audio topic and enqueue the raw data into the pbReader
	natsc.Subscribe("fromRadio", func(m *nats.Msg) {
		err := fromNetwork.Enqueue(m.Data)
		if err != nil {
			log.Println(err)
		}
	})

	// Callback which is called by pbWriter to push the audio
	// packets to the network
	toWireCb := func(data []byte) {
		err := natsc.Publish("toRadio", data)
		if err != nil {
			log.Println(err)
		}
	}

	// opus Encoder for PbWriter
	opusEncoder, err := opus.NewEncoder(
		opus.Bitrate(opusBitrate),
		opus.Complexity(opusComplexity),
		opus.Channels(iChannels),
		opus.Samplerate(iSamplerate),
		// opus.MaxBandwidth(opusMaxBandwidth),
	)

	toNetwork, err := pbWriter.NewPbWriter(toWireCb,
		pbWriter.Encoder(opusEncoder),
		pbWriter.Samplerate(iSamplerate),
		pbWriter.Channels(iChannels),
		pbWriter.FramesPerBuffer(audioFramesPerBuffer),
	)
	if err != nil {
		log.Fatal(err)
	}

	nc.fromRadioSources.AddSource("file", wav)
	nc.fromRadioSources.AddSource("fromNetwork", fromNetwork)
	nc.fromRadioSinks.AddSink("fromRadioAudio", fromRadioAudio, false)

	nc.toRadioSources.AddSource("file", wav)
	nc.toRadioSources.AddSource("toRadioAudio", toRadioAudio)
	nc.toRadioSinks.AddSink("toNetwork", toNetwork, false)

	// Channel to handle OS signals
	osSignals := make(chan os.Signal, 1)

	//subscribe to os.Interrupt (CTRL-C signal)
	signal.Notify(osSignals, os.Interrupt)

	// set callback to process audio fro
	nc.fromRadioSources.SetCb(nc.toRxSinksCb)
	// start streaming to the network immediately
	nc.fromRadioSinks.EnableSink("fromRadioAudio", true)
	nc.fromRadioSources.SetSource("fromNetwork")

	// set callback to process audio to be send to the radio
	nc.toRadioSources.SetCb(nc.toTxSinksCb)

	// stream immediately audio from the network to the radio
	nc.toRadioSources.SetSource("toRadioAudio")
	nc.toRadioSinks.EnableSink("toNetwork", true)

	keyb := make(chan string, 10)

	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			keyb <- strings.TrimSuffix(text, "\n")
		}
	}()

	for {
		select {
		case input := <-keyb:
			switch input {
			// case "a":
			// 	if err := n.router.EnableSink("wavFile", true); err != nil {
			// 		log.Println(err)
			// 	}
			// case "b":
			// 	if err := n.router.EnableSink("wavFile", false); err != nil {
			// 		log.Println(err)
			// 	}
			case "f":
				nc.fromRadioSinks.Flush()
				if err := nc.fromRadioSources.SetSource("file"); err != nil {
					log.Println(err)
				}
				nc.toRadioSinks.Flush()
				if err := nc.toRadioSources.SetSource("file"); err != nil {
					log.Println(err)
				}

			case "m":
				nc.toRadioSinks.Flush()
				if err := nc.fromRadioSources.SetSource("fromNetwork"); err != nil {
					log.Println(err)
				}
			case "i":
				fromRadioAudio.SetVolume(fromRadioAudio.Volume() + 0.5)
			case "d":
				fromRadioAudio.SetVolume(fromRadioAudio.Volume() - 0.5)
			}
		case sig := <-osSignals:
			if sig == os.Interrupt {
				// TBD: close also router (and all sinks)
				toRadioAudio.Close()
				fromRadioAudio.Close()
				return
			}
		}
	}
}

type natsClient struct {
	fromRadioSinks   audio.Router   //rx path
	fromRadioSources audio.Selector //rx path
	toRadioSinks     audio.Router   //tx path
	toRadioSources   audio.Selector //tx path
	isPlaying        bool
	play             chan audio.Msg
}

func (nc *natsClient) toRxSinksCb(data audio.Msg) {
	token := nc.fromRadioSinks.Write(data)
	if token.Error != nil {
		// handle Error -> remove source
	}
	token.Wait()
	if data.EOF {
		// switch back to default source
		nc.fromRadioSinks.Flush()
		if err := nc.fromRadioSources.SetSource("fromNetwork"); err != nil {
			log.Println(err)
		}
	}
}

func (nc *natsClient) toTxSinksCb(data audio.Msg) {
	token := nc.toRadioSinks.Write(data)
	if token.Error != nil {
		// handle Error -> remove source
	}
	token.Wait()
	if data.EOF {
		// switch back to default source
		nc.toRadioSinks.Flush()
		if err := nc.toRadioSources.SetSource("toRadioAudio"); err != nil {
			log.Println(err)
		}
	}
}

// // GetOpusApplication returns the integer representation of a
// // Opus application value string (typically read from application settings)
// func GetOpusApplication(app string) (opus.Application, error) {
// 	switch app {
// 	case "audio":
// 		return opus.AppAudio, nil
// 	case "restricted_lowdelay":
// 		return opus.AppRestrictedLowdelay, nil
// 	case "voip":
// 		return opus.AppVoIP, nil
// 	}
// 	return 0, errors.New("unknown opus application value")
// }

// // GetOpusMaxBandwith returns the integer representation of an
// // Opus max bandwitdh value string (typically read from application settings)
// func GetOpusMaxBandwith(maxBw string) (opus.Bandwidth, error) {
// 	switch strings.ToLower(maxBw) {
// 	case "narrowband":
// 		return opus.Narrowband, nil
// 	case "mediumband":
// 		return opus.Mediumband, nil
// 	case "wideband":
// 		return opus.Wideband, nil
// 	case "superwideband":
// 		return opus.SuperWideband, nil
// 	case "fullband":
// 		return opus.Fullband, nil
// 	}

// 	return 0, errors.New("unknown opus max bandwidth value")
// }

// // GetCodec return the integer representation of a string containing the
// // name of an audio codec
// func GetCodec(codec string) (int, error) {
// 	switch strings.ToLower(codec) {
// 	case "pcm":
// 		return PCM, nil
// 	case "opus":
// 		return OPUS, nil
// 	}
// 	errMsg := fmt.Sprintf("unknown codec: %s", codec)
// 	return 0, errors.New(errMsg)
// }
