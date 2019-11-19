// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: import_public/a.proto

package import_public

import (
	fmt "fmt"
	proto "github.com/tron-us/protobuf/proto"
	sub "github.com/tron-us/protobuf/protoc-gen-gogo/testdata/import_public/sub"
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

const Default_M_DefaultField = sub.Default_M_DefaultField

// M from public import import_public/sub/a.proto
type M = sub.M
type M_OneofInt32 = sub.M_OneofInt32
type M_OneofInt64 = sub.M_OneofInt64

// M_Grouping from public import import_public/sub/a.proto
type M_Grouping = sub.M_Grouping

// M_Submessage from public import import_public/sub/a.proto
type M_Submessage = sub.M_Submessage
type M_Submessage_SubmessageOneofInt32 = sub.M_Submessage_SubmessageOneofInt32
type M_Submessage_SubmessageOneofInt64 = sub.M_Submessage_SubmessageOneofInt64

// E from public import import_public/sub/a.proto
type E = sub.E

var E_name = sub.E_name
var E_value = sub.E_value

const E_ZERO = E(sub.E_ZERO)

// M_Subenum from public import import_public/sub/a.proto
type M_Subenum = sub.M_Subenum

var M_Subenum_name = sub.M_Subenum_name
var M_Subenum_value = sub.M_Subenum_value

const M_M_ZERO = M_Subenum(sub.M_M_ZERO)

// M_Submessage_Submessage_Subenum from public import import_public/sub/a.proto
type M_Submessage_Submessage_Subenum = sub.M_Submessage_Submessage_Subenum

var M_Submessage_Submessage_Subenum_name = sub.M_Submessage_Submessage_Subenum_name
var M_Submessage_Submessage_Subenum_value = sub.M_Submessage_Submessage_Subenum_value

const M_Submessage_M_SUBMESSAGE_ZERO = M_Submessage_Submessage_Subenum(sub.M_Submessage_M_SUBMESSAGE_ZERO)

var E_ExtensionField = sub.E_ExtensionField

type Public struct {
	M                    *sub.M   `protobuf:"bytes,1,opt,name=m" json:"m,omitempty" pg:"m"`
	E                    *sub.E   `protobuf:"varint,2,opt,name=e,enum=goproto.test.import_public.sub.E" json:"e,omitempty" pg:"e"`
	Local                *Local   `protobuf:"bytes,3,opt,name=local" json:"local,omitempty" pg:"local"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" pg:"-"`
	XXX_unrecognized     []byte   `json:"-" pg:"-"`
	XXX_sizecache        int32    `json:"-" pg:"-"`
}

func (m *Public) Reset()         { *m = Public{} }
func (m *Public) String() string { return proto.CompactTextString(m) }
func (*Public) ProtoMessage()    {}
func (*Public) Descriptor() ([]byte, []int) {
	return fileDescriptor_73b7577c95fa6b70, []int{0}
}
func (m *Public) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Public.Unmarshal(m, b)
}
func (m *Public) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Public.Marshal(b, m, deterministic)
}
func (m *Public) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Public.Merge(m, src)
}
func (m *Public) XXX_Size() int {
	return xxx_messageInfo_Public.Size(m)
}
func (m *Public) XXX_DiscardUnknown() {
	xxx_messageInfo_Public.DiscardUnknown(m)
}

var xxx_messageInfo_Public proto.InternalMessageInfo

func (m *Public) GetM() *sub.M {
	if m != nil {
		return m.M
	}
	return nil
}

func (m *Public) GetE() sub.E {
	if m != nil && m.E != nil {
		return *m.E
	}
	return sub.E_ZERO
}

func (m *Public) GetLocal() *Local {
	if m != nil {
		return m.Local
	}
	return nil
}

func init() {
	proto.RegisterType((*Public)(nil), "goproto.test.import_public.Public")
}

func init() { proto.RegisterFile("import_public/a.proto", fileDescriptor_73b7577c95fa6b70) }

var fileDescriptor_73b7577c95fa6b70 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8d, 0x31, 0xcb, 0xc2, 0x30,
	0x18, 0x84, 0xbf, 0xf7, 0x13, 0x1d, 0x22, 0x38, 0x14, 0x84, 0xda, 0xa9, 0x3a, 0x75, 0x69, 0x02,
	0x2e, 0xee, 0x45, 0x37, 0x85, 0xd2, 0xd1, 0x45, 0x92, 0x18, 0x63, 0xa1, 0xe9, 0x5b, 0x9a, 0xe4,
	0x17, 0xf9, 0x47, 0xa5, 0x2d, 0x0e, 0x15, 0xa4, 0xdb, 0x71, 0xdc, 0xf3, 0x1c, 0x59, 0x97, 0xa6,
	0xc1, 0xd6, 0xdd, 0x1a, 0x2f, 0xaa, 0x52, 0x32, 0x4e, 0x9b, 0x16, 0x1d, 0x06, 0x91, 0xc6, 0x3e,
	0x50, 0xa7, 0xac, 0xa3, 0xa3, 0x4d, 0xb4, 0x19, 0x23, 0xd6, 0x8b, 0x0f, 0x16, 0x7d, 0xd9, 0xc4,
	0x50, 0xef, 0x5e, 0x40, 0x16, 0x79, 0x5f, 0x05, 0x8c, 0x80, 0x09, 0x21, 0x86, 0x64, 0xb9, 0xdf,
	0xd2, 0xdf, 0x27, 0xd4, 0x7a, 0x41, 0x2f, 0x05, 0x98, 0x0e, 0x50, 0xe1, 0x7f, 0x0c, 0xc9, 0x6a,
	0x1a, 0x38, 0x15, 0xa0, 0x82, 0x03, 0x99, 0x57, 0x28, 0x79, 0x15, 0xce, 0xa6, 0x5f, 0xce, 0xdd,
	0xb0, 0x18, 0xf6, 0xd9, 0xf1, 0x9a, 0xe9, 0xd2, 0x3d, 0xbd, 0xa0, 0x12, 0x0d, 0x73, 0x2d, 0xd6,
	0xa9, 0xb7, 0xac, 0x67, 0x85, 0x7f, 0x0c, 0x41, 0xa6, 0x5a, 0xd5, 0xa9, 0x46, 0x8d, 0xac, 0xd3,
	0xdd, 0xb9, 0xe3, 0x6c, 0xa4, 0xcc, 0xff, 0x72, 0x78, 0x07, 0x00, 0x00, 0xff, 0xff, 0x69, 0xe2,
	0x5d, 0xf0, 0x55, 0x01, 0x00, 0x00,
}
