// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: import_public/b.proto

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

type Local struct {
	M                    *sub.M   `protobuf:"bytes,1,opt,name=m" json:"m,omitempty" pg:"m"`
	E                    *sub.E   `protobuf:"varint,2,opt,name=e,enum=goproto.test.import_public.sub.E" json:"e,omitempty" pg:"e"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" pg:"-"`
	XXX_unrecognized     []byte   `json:"-" pg:"-"`
	XXX_sizecache        int32    `json:"-" pg:"-"`
}

func (m *Local) Reset()         { *m = Local{} }
func (m *Local) String() string { return proto.CompactTextString(m) }
func (*Local) ProtoMessage()    {}
func (*Local) Descriptor() ([]byte, []int) {
	return fileDescriptor_84995586b3d09710, []int{0}
}
func (m *Local) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Local.Unmarshal(m, b)
}
func (m *Local) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Local.Marshal(b, m, deterministic)
}
func (m *Local) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Local.Merge(m, src)
}
func (m *Local) XXX_Size() int {
	return xxx_messageInfo_Local.Size(m)
}
func (m *Local) XXX_DiscardUnknown() {
	xxx_messageInfo_Local.DiscardUnknown(m)
}

var xxx_messageInfo_Local proto.InternalMessageInfo

func (m *Local) GetM() *sub.M {
	if m != nil {
		return m.M
	}
	return nil
}

func (m *Local) GetE() sub.E {
	if m != nil && m.E != nil {
		return *m.E
	}
	return sub.E_ZERO
}

func init() {
	proto.RegisterType((*Local)(nil), "goproto.test.import_public.Local")
}

func init() { proto.RegisterFile("import_public/b.proto", fileDescriptor_84995586b3d09710) }

var fileDescriptor_84995586b3d09710 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xcc, 0x31, 0xcb, 0xc2, 0x30,
	0x10, 0xc6, 0x71, 0xf2, 0xc2, 0xbb, 0x54, 0x70, 0x28, 0x08, 0xb5, 0x53, 0x75, 0xea, 0xd2, 0x3b,
	0xf0, 0x23, 0x14, 0xdd, 0x74, 0xe9, 0xe8, 0x22, 0x49, 0x8c, 0x31, 0xd0, 0xf4, 0x4a, 0x73, 0xf9,
	0xfe, 0xd2, 0x76, 0xca, 0x20, 0x6e, 0xcf, 0xf0, 0xfc, 0xfe, 0xd9, 0xce, 0xf9, 0x91, 0x26, 0x7e,
	0x8c, 0x51, 0xf5, 0x4e, 0xa3, 0x82, 0x71, 0x22, 0xa6, 0xbc, 0xb4, 0xb4, 0x0c, 0x60, 0x13, 0x18,
	0x92, 0x4f, 0xb9, 0x4f, 0x49, 0x88, 0x0a, 0xe5, 0xca, 0x8e, 0x2e, 0xfb, 0xbf, 0x92, 0x96, 0x7d,
	0x8e, 0x99, 0xf0, 0x85, 0xa8, 0x44, 0xbd, 0x39, 0x1d, 0xe0, 0x7b, 0x0b, 0x42, 0x54, 0x70, 0xeb,
	0x84, 0x9f, 0x81, 0x29, 0xfe, 0x2a, 0x51, 0x6f, 0x7f, 0x83, 0x4b, 0x27, 0x4c, 0x7b, 0xbe, 0xb7,
	0xd6, 0xf1, 0x3b, 0x2a, 0xd0, 0xe4, 0x91, 0x27, 0x1a, 0x9a, 0x18, 0x70, 0x71, 0x2a, 0xbe, 0xd6,
	0xa1, 0x1b, 0x6b, 0x86, 0xc6, 0x92, 0x25, 0x9c, 0x53, 0x4f, 0xc9, 0x12, 0x93, 0xdc, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x8c, 0x49, 0x40, 0x30, 0xff, 0x00, 0x00, 0x00,
}
