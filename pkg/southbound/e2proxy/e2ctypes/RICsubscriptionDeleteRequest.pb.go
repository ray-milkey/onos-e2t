// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RICsubscriptionDeleteRequest.proto

package e2ctypes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Deprecated: Do not use.
type RICsubscriptionDeleteRequestT struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RICsubscriptionDeleteRequestT) Reset()         { *m = RICsubscriptionDeleteRequestT{} }
func (m *RICsubscriptionDeleteRequestT) String() string { return proto.CompactTextString(m) }
func (*RICsubscriptionDeleteRequestT) ProtoMessage()    {}
func (*RICsubscriptionDeleteRequestT) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee1bbab4eb622ffd, []int{0}
}

func (m *RICsubscriptionDeleteRequestT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RICsubscriptionDeleteRequestT.Unmarshal(m, b)
}
func (m *RICsubscriptionDeleteRequestT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RICsubscriptionDeleteRequestT.Marshal(b, m, deterministic)
}
func (m *RICsubscriptionDeleteRequestT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RICsubscriptionDeleteRequestT.Merge(m, src)
}
func (m *RICsubscriptionDeleteRequestT) XXX_Size() int {
	return xxx_messageInfo_RICsubscriptionDeleteRequestT.Size(m)
}
func (m *RICsubscriptionDeleteRequestT) XXX_DiscardUnknown() {
	xxx_messageInfo_RICsubscriptionDeleteRequestT.DiscardUnknown(m)
}

var xxx_messageInfo_RICsubscriptionDeleteRequestT proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RICsubscriptionDeleteRequestT)(nil), "e2ctypes.RICsubscriptionDeleteRequest_t")
}

func init() { proto.RegisterFile("RICsubscriptionDeleteRequest.proto", fileDescriptor_ee1bbab4eb622ffd) }

var fileDescriptor_ee1bbab4eb622ffd = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x0a, 0xf2, 0x74, 0x2e,
	0x2e, 0x4d, 0x2a, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0xcc, 0xcf, 0x73, 0x49, 0xcd, 0x49, 0x2d,
	0x49, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x48, 0x35, 0x4a, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x96, 0x92, 0x0a, 0x00, 0x09, 0x24, 0xe7, 0xe7,
	0x78, 0xba, 0xea, 0x3a, 0xe7, 0xe7, 0x95, 0x24, 0x66, 0xe6, 0xa5, 0x16, 0x41, 0x54, 0x29, 0xa9,
	0x70, 0xc9, 0xe1, 0x33, 0x2b, 0xbe, 0xc4, 0x8a, 0x49, 0x82, 0x31, 0x89, 0x0d, 0xac, 0xd8, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x02, 0xbf, 0x79, 0x4f, 0x78, 0x00, 0x00, 0x00,
}