// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RICserviceUpdate.proto

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
type RICserviceUpdateT struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RICserviceUpdateT) Reset()         { *m = RICserviceUpdateT{} }
func (m *RICserviceUpdateT) String() string { return proto.CompactTextString(m) }
func (*RICserviceUpdateT) ProtoMessage()    {}
func (*RICserviceUpdateT) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd482fdaff02def2, []int{0}
}

func (m *RICserviceUpdateT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RICserviceUpdateT.Unmarshal(m, b)
}
func (m *RICserviceUpdateT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RICserviceUpdateT.Marshal(b, m, deterministic)
}
func (m *RICserviceUpdateT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RICserviceUpdateT.Merge(m, src)
}
func (m *RICserviceUpdateT) XXX_Size() int {
	return xxx_messageInfo_RICserviceUpdateT.Size(m)
}
func (m *RICserviceUpdateT) XXX_DiscardUnknown() {
	xxx_messageInfo_RICserviceUpdateT.DiscardUnknown(m)
}

var xxx_messageInfo_RICserviceUpdateT proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RICserviceUpdateT)(nil), "e2ctypes.RICserviceUpdate_t")
}

func init() { proto.RegisterFile("RICserviceUpdate.proto", fileDescriptor_fd482fdaff02def2) }

var fileDescriptor_fd482fdaff02def2 = []byte{
	// 97 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0b, 0xf2, 0x74, 0x2e,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x0d, 0x2d, 0x48, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x35, 0x4a, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x96, 0x92, 0x0a, 0x00,
	0x09, 0x24, 0xe7, 0xe7, 0x78, 0xba, 0xea, 0x3a, 0xe7, 0xe7, 0x95, 0x24, 0x66, 0xe6, 0xa5, 0x16,
	0x41, 0x54, 0x29, 0x49, 0x70, 0x09, 0xa1, 0xeb, 0x8f, 0x2f, 0xb1, 0x62, 0x92, 0x60, 0x4c, 0x62,
	0x03, 0x2b, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x24, 0xca, 0xd9, 0xc5, 0x60, 0x00, 0x00,
	0x00,
}