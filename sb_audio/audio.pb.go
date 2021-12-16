// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: audio.proto

package sb_audio

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Channels int32

const (
	Channels_unknown Channels = 0
	Channels_mono    Channels = 1
	Channels_stereo  Channels = 2
)

// Enum value maps for Channels.
var (
	Channels_name = map[int32]string{
		0: "unknown",
		1: "mono",
		2: "stereo",
	}
	Channels_value = map[string]int32{
		"unknown": 0,
		"mono":    1,
		"stereo":  2,
	}
)

func (x Channels) Enum() *Channels {
	p := new(Channels)
	*p = x
	return p
}

func (x Channels) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Channels) Descriptor() protoreflect.EnumDescriptor {
	return file_audio_proto_enumTypes[0].Descriptor()
}

func (Channels) Type() protoreflect.EnumType {
	return &file_audio_proto_enumTypes[0]
}

func (x Channels) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Channels.Descriptor instead.
func (Channels) EnumDescriptor() ([]byte, []int) {
	return file_audio_proto_rawDescGZIP(), []int{0}
}

type Codec int32

const (
	Codec_none Codec = 0
	Codec_opus Codec = 1
	Codec_pcm  Codec = 2
)

// Enum value maps for Codec.
var (
	Codec_name = map[int32]string{
		0: "none",
		1: "opus",
		2: "pcm",
	}
	Codec_value = map[string]int32{
		"none": 0,
		"opus": 1,
		"pcm":  2,
	}
)

func (x Codec) Enum() *Codec {
	p := new(Codec)
	*p = x
	return p
}

func (x Codec) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Codec) Descriptor() protoreflect.EnumDescriptor {
	return file_audio_proto_enumTypes[1].Descriptor()
}

func (Codec) Type() protoreflect.EnumType {
	return &file_audio_proto_enumTypes[1]
}

func (x Codec) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Codec.Descriptor instead.
func (Codec) EnumDescriptor() ([]byte, []int) {
	return file_audio_proto_rawDescGZIP(), []int{1}
}

type None struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *None) Reset() {
	*x = None{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *None) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*None) ProtoMessage() {}

func (x *None) ProtoReflect() protoreflect.Message {
	mi := &file_audio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use None.ProtoReflect.Descriptor instead.
func (*None) Descriptor() ([]byte, []int) {
	return file_audio_proto_rawDescGZIP(), []int{0}
}

type Capabilities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                                                            // name of the server
	RxStreamAddress     string `protobuf:"bytes,2,opt,name=rx_stream_address,json=rxStreamAddress,proto3" json:"rx_stream_address,omitempty"`             // where the Server publishes audio from the radio
	TxStreamAddress     string `protobuf:"bytes,3,opt,name=tx_stream_address,json=txStreamAddress,proto3" json:"tx_stream_address,omitempty"`             // where the Server listens for audio to be transmitted on the radio
	StateUpdatesAddress string `protobuf:"bytes,4,opt,name=state_updates_address,json=stateUpdatesAddress,proto3" json:"state_updates_address,omitempty"` // where the Server listens for audio to be transmitted on the radio
	Index               int32  `protobuf:"varint,5,opt,name=index,proto3" json:"index,omitempty"`                                                         // static index for displaying several servers consistently in a GUI
}

func (x *Capabilities) Reset() {
	*x = Capabilities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Capabilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Capabilities) ProtoMessage() {}

func (x *Capabilities) ProtoReflect() protoreflect.Message {
	mi := &file_audio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Capabilities.ProtoReflect.Descriptor instead.
func (*Capabilities) Descriptor() ([]byte, []int) {
	return file_audio_proto_rawDescGZIP(), []int{1}
}

func (x *Capabilities) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Capabilities) GetRxStreamAddress() string {
	if x != nil {
		return x.RxStreamAddress
	}
	return ""
}

func (x *Capabilities) GetTxStreamAddress() string {
	if x != nil {
		return x.TxStreamAddress
	}
	return ""
}

func (x *Capabilities) GetStateUpdatesAddress() string {
	if x != nil {
		return x.StateUpdatesAddress
	}
	return ""
}

func (x *Capabilities) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type PingPong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping int64 `protobuf:"varint,1,opt,name=ping,proto3" json:"ping,omitempty"` // unix timestamp
}

func (x *PingPong) Reset() {
	*x = PingPong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingPong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingPong) ProtoMessage() {}

func (x *PingPong) ProtoReflect() protoreflect.Message {
	mi := &file_audio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingPong.ProtoReflect.Descriptor instead.
func (*PingPong) Descriptor() ([]byte, []int) {
	return file_audio_proto_rawDescGZIP(), []int{2}
}

func (x *PingPong) GetPing() int64 {
	if x != nil {
		return x.Ping
	}
	return 0
}

// Audio frame consisting of the raw audio byte array + metadata
type Frame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Codec        Codec    `protobuf:"varint,1,opt,name=codec,proto3,enum=shackbus.audio.Codec" json:"codec,omitempty"`
	Channels     Channels `protobuf:"varint,2,opt,name=channels,proto3,enum=shackbus.audio.Channels" json:"channels,omitempty"` // Number of channels
	FrameLength  int32    `protobuf:"varint,3,opt,name=frame_length,json=frameLength,proto3" json:"frame_length,omitempty"`     // Audio frame length (in bytes)
	SamplingRate int32    `protobuf:"varint,4,opt,name=sampling_rate,json=samplingRate,proto3" json:"sampling_rate,omitempty"`  // Audio sampling rate
	BitDepth     int32    `protobuf:"varint,5,opt,name=bit_depth,json=bitDepth,proto3" json:"bit_depth,omitempty"`              // Audio bit depth (8...16 bit typically)
	Data         []byte   `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`                                       // Audio packets as raw byte array
	UserId       string   `protobuf:"bytes,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Frame) Reset() {
	*x = Frame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Frame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Frame) ProtoMessage() {}

func (x *Frame) ProtoReflect() protoreflect.Message {
	mi := &file_audio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Frame.ProtoReflect.Descriptor instead.
func (*Frame) Descriptor() ([]byte, []int) {
	return file_audio_proto_rawDescGZIP(), []int{3}
}

func (x *Frame) GetCodec() Codec {
	if x != nil {
		return x.Codec
	}
	return Codec_none
}

func (x *Frame) GetChannels() Channels {
	if x != nil {
		return x.Channels
	}
	return Channels_unknown
}

func (x *Frame) GetFrameLength() int32 {
	if x != nil {
		return x.FrameLength
	}
	return 0
}

func (x *Frame) GetSamplingRate() int32 {
	if x != nil {
		return x.SamplingRate
	}
	return 0
}

func (x *Frame) GetBitDepth() int32 {
	if x != nil {
		return x.BitDepth
	}
	return 0
}

func (x *Frame) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Frame) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RxOn   bool   `protobuf:"varint,1,opt,name=rx_on,json=rxOn,proto3" json:"rx_on,omitempty"`
	TxUser string `protobuf:"bytes,3,opt,name=tx_user,json=txUser,proto3" json:"tx_user,omitempty"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_audio_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_audio_proto_rawDescGZIP(), []int{4}
}

func (x *State) GetRxOn() bool {
	if x != nil {
		return x.RxOn
	}
	return false
}

func (x *State) GetTxUser() string {
	if x != nil {
		return x.TxUser
	}
	return ""
}

var File_audio_proto protoreflect.FileDescriptor

var file_audio_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73,
	0x68, 0x61, 0x63, 0x6b, 0x62, 0x75, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x22, 0x06, 0x0a,
	0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x22, 0xc4, 0x01, 0x0a, 0x0c, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x78,
	0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x78, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x78, 0x5f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x74, 0x78, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x73, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x13, 0x73, 0x74, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x1e, 0x0a, 0x08,
	0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x22, 0xfc, 0x01, 0x0a,
	0x05, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x68, 0x61, 0x63, 0x6b, 0x62, 0x75, 0x73,
	0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x52, 0x05, 0x63, 0x6f,
	0x64, 0x65, 0x63, 0x12, 0x34, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73, 0x68, 0x61, 0x63, 0x6b, 0x62, 0x75, 0x73,
	0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x52,
	0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x23, 0x0a, 0x0d,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x74, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x62, 0x69, 0x74, 0x44, 0x65, 0x70, 0x74, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x72, 0x78, 0x5f, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x72, 0x78, 0x4f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x55, 0x73,
	0x65, 0x72, 0x2a, 0x2d, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x0b,
	0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x6d,
	0x6f, 0x6e, 0x6f, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x73, 0x74, 0x65, 0x72, 0x65, 0x6f, 0x10,
	0x02, 0x2a, 0x24, 0x0a, 0x05, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x12, 0x08, 0x0a, 0x04, 0x6e, 0x6f,
	0x6e, 0x65, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x6f, 0x70, 0x75, 0x73, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x70, 0x63, 0x6d, 0x10, 0x02, 0x32, 0xb9, 0x02, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x45, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x73, 0x68, 0x61, 0x63, 0x6b, 0x62, 0x75, 0x73,
	0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x4e, 0x6f, 0x6e, 0x65, 0x1a, 0x1c, 0x2e, 0x73, 0x68,
	0x61, 0x63, 0x6b, 0x62, 0x75, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x43, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x68, 0x61, 0x63, 0x6b, 0x62, 0x75, 0x73,
	0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x4e, 0x6f, 0x6e, 0x65, 0x1a, 0x15, 0x2e, 0x73, 0x68,
	0x61, 0x63, 0x6b, 0x62, 0x75, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x14, 0x2e, 0x73, 0x68, 0x61, 0x63, 0x6b, 0x62, 0x75, 0x73, 0x2e, 0x61, 0x75, 0x64,
	0x69, 0x6f, 0x2e, 0x4e, 0x6f, 0x6e, 0x65, 0x1a, 0x14, 0x2e, 0x73, 0x68, 0x61, 0x63, 0x6b, 0x62,
	0x75, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x4e, 0x6f, 0x6e, 0x65, 0x12, 0x38, 0x0a,
	0x0a, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x14, 0x2e, 0x73, 0x68,
	0x61, 0x63, 0x6b, 0x62, 0x75, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x4e, 0x6f, 0x6e,
	0x65, 0x1a, 0x14, 0x2e, 0x73, 0x68, 0x61, 0x63, 0x6b, 0x62, 0x75, 0x73, 0x2e, 0x61, 0x75, 0x64,
	0x69, 0x6f, 0x2e, 0x4e, 0x6f, 0x6e, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12,
	0x18, 0x2e, 0x73, 0x68, 0x61, 0x63, 0x6b, 0x62, 0x75, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x1a, 0x18, 0x2e, 0x73, 0x68, 0x61, 0x63,
	0x6b, 0x62, 0x75, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x50,
	0x6f, 0x6e, 0x67, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x73, 0x62, 0x5f, 0x61, 0x75, 0x64, 0x69,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_audio_proto_rawDescOnce sync.Once
	file_audio_proto_rawDescData = file_audio_proto_rawDesc
)

func file_audio_proto_rawDescGZIP() []byte {
	file_audio_proto_rawDescOnce.Do(func() {
		file_audio_proto_rawDescData = protoimpl.X.CompressGZIP(file_audio_proto_rawDescData)
	})
	return file_audio_proto_rawDescData
}

var file_audio_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_audio_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_audio_proto_goTypes = []interface{}{
	(Channels)(0),        // 0: shackbus.audio.Channels
	(Codec)(0),           // 1: shackbus.audio.Codec
	(*None)(nil),         // 2: shackbus.audio.None
	(*Capabilities)(nil), // 3: shackbus.audio.Capabilities
	(*PingPong)(nil),     // 4: shackbus.audio.PingPong
	(*Frame)(nil),        // 5: shackbus.audio.Frame
	(*State)(nil),        // 6: shackbus.audio.State
}
var file_audio_proto_depIdxs = []int32{
	1, // 0: shackbus.audio.Frame.codec:type_name -> shackbus.audio.Codec
	0, // 1: shackbus.audio.Frame.channels:type_name -> shackbus.audio.Channels
	2, // 2: shackbus.audio.Server.GetCapabilities:input_type -> shackbus.audio.None
	2, // 3: shackbus.audio.Server.GetState:input_type -> shackbus.audio.None
	2, // 4: shackbus.audio.Server.StartStream:input_type -> shackbus.audio.None
	2, // 5: shackbus.audio.Server.StopStream:input_type -> shackbus.audio.None
	4, // 6: shackbus.audio.Server.Ping:input_type -> shackbus.audio.PingPong
	3, // 7: shackbus.audio.Server.GetCapabilities:output_type -> shackbus.audio.Capabilities
	6, // 8: shackbus.audio.Server.GetState:output_type -> shackbus.audio.State
	2, // 9: shackbus.audio.Server.StartStream:output_type -> shackbus.audio.None
	2, // 10: shackbus.audio.Server.StopStream:output_type -> shackbus.audio.None
	4, // 11: shackbus.audio.Server.Ping:output_type -> shackbus.audio.PingPong
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_audio_proto_init() }
func file_audio_proto_init() {
	if File_audio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_audio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*None); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_audio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Capabilities); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_audio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingPong); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_audio_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Frame); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_audio_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_audio_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_audio_proto_goTypes,
		DependencyIndexes: file_audio_proto_depIdxs,
		EnumInfos:         file_audio_proto_enumTypes,
		MessageInfos:      file_audio_proto_msgTypes,
	}.Build()
	File_audio_proto = out.File
	file_audio_proto_rawDesc = nil
	file_audio_proto_goTypes = nil
	file_audio_proto_depIdxs = nil
}
