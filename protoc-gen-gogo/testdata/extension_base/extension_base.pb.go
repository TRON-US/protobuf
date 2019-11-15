// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: extension_base/extension_base.proto

package extension_base

import (
	fmt "fmt"
	proto "github.com/tron-us/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type BaseMessage struct {
	Height                       *int32   `protobuf:"varint,1,opt,name=height" json:"height,omitempty" pg:"height"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *BaseMessage) Reset()         { *m = BaseMessage{} }
func (m *BaseMessage) String() string { return proto.CompactTextString(m) }
func (*BaseMessage) ProtoMessage()    {}
func (*BaseMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbd53bac0b7ca8a, []int{0}
}

var extRange_BaseMessage = []proto.ExtensionRange{
	{Start: 4, End: 9},
	{Start: 16, End: 536870911},
}

func (*BaseMessage) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_BaseMessage
}

func (m *BaseMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseMessage.Unmarshal(m, b)
}
func (m *BaseMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseMessage.Marshal(b, m, deterministic)
}
func (m *BaseMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseMessage.Merge(m, src)
}
func (m *BaseMessage) XXX_Size() int {
	return xxx_messageInfo_BaseMessage.Size(m)
}
func (m *BaseMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BaseMessage proto.InternalMessageInfo

func (m *BaseMessage) GetHeight() int32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

// Another message that may be extended, using message_set_wire_format.
type OldStyleMessage struct {
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `protobuf_messageset:"1" json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *OldStyleMessage) Reset()         { *m = OldStyleMessage{} }
func (m *OldStyleMessage) String() string { return proto.CompactTextString(m) }
func (*OldStyleMessage) ProtoMessage()    {}
func (*OldStyleMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbd53bac0b7ca8a, []int{1}
}

var extRange_OldStyleMessage = []proto.ExtensionRange{
	{Start: 100, End: 2147483646},
}

func (*OldStyleMessage) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_OldStyleMessage
}

func (m *OldStyleMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OldStyleMessage.Unmarshal(m, b)
}
func (m *OldStyleMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OldStyleMessage.Marshal(b, m, deterministic)
}
func (m *OldStyleMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OldStyleMessage.Merge(m, src)
}
func (m *OldStyleMessage) XXX_Size() int {
	return xxx_messageInfo_OldStyleMessage.Size(m)
}
func (m *OldStyleMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_OldStyleMessage.DiscardUnknown(m)
}

var xxx_messageInfo_OldStyleMessage proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BaseMessage)(nil), "extension_base.BaseMessage")
	proto.RegisterType((*OldStyleMessage)(nil), "extension_base.OldStyleMessage")
}

func init() {
	proto.RegisterFile("extension_base/extension_base.proto", fileDescriptor_2fbd53bac0b7ca8a)
}

var fileDescriptor_2fbd53bac0b7ca8a = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xad, 0x28, 0x49,
	0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x8b, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x47, 0xe5, 0xea, 0x15, 0x14,
	0xe5, 0x97, 0xe4, 0x0b, 0xf1, 0xa1, 0x8a, 0x2a, 0x99, 0x72, 0x71, 0x3b, 0x25, 0x16, 0xa7, 0xfa,
	0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x0a, 0x89, 0x71, 0xb1, 0x65, 0xa4, 0x66, 0xa6, 0x67, 0x94,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x41, 0x79, 0x5a, 0x2c, 0x1c, 0x2c, 0x02, 0x5c, 0x5a,
	0x1c, 0x1c, 0x02, 0x02, 0x0d, 0x0d, 0x0d, 0x0d, 0x4c, 0x4a, 0xf2, 0x5c, 0xfc, 0xfe, 0x39, 0x29,
	0xc1, 0x25, 0x95, 0x39, 0x30, 0xad, 0x5a, 0x1c, 0x1c, 0x29, 0x02, 0xff, 0xff, 0xff, 0xff, 0xcf,
	0x6e, 0xc5, 0xc4, 0xc1, 0xe8, 0xe4, 0x1a, 0xe5, 0x9c, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97,
	0x9c, 0x9f, 0xab, 0x5f, 0x52, 0x94, 0x9f, 0xa7, 0x5b, 0x5a, 0xac, 0x0f, 0x76, 0x43, 0x52, 0x69,
	0x1a, 0x84, 0x91, 0xac, 0x9b, 0x9e, 0x9a, 0xa7, 0x9b, 0x9e, 0x9f, 0x9e, 0xaf, 0x5f, 0x92, 0x5a,
	0x5c, 0x92, 0x92, 0x58, 0x92, 0x88, 0xe6, 0x68, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xa7,
	0x06, 0xdf, 0xd4, 0x00, 0x00, 0x00,
}
