// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RICsubscriptionDeleteResponse.proto

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
type RICsubscriptionDeleteResponseT struct {
	ProtocolIEs          *ProtocolIE_Container_1544P4T `protobuf:"bytes,1,opt,name=protocolIEs,proto3" json:"protocolIEs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *RICsubscriptionDeleteResponseT) Reset()         { *m = RICsubscriptionDeleteResponseT{} }
func (m *RICsubscriptionDeleteResponseT) String() string { return proto.CompactTextString(m) }
func (*RICsubscriptionDeleteResponseT) ProtoMessage()    {}
func (*RICsubscriptionDeleteResponseT) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6829107780d54e7, []int{0}
}

func (m *RICsubscriptionDeleteResponseT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RICsubscriptionDeleteResponseT.Unmarshal(m, b)
}
func (m *RICsubscriptionDeleteResponseT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RICsubscriptionDeleteResponseT.Marshal(b, m, deterministic)
}
func (m *RICsubscriptionDeleteResponseT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RICsubscriptionDeleteResponseT.Merge(m, src)
}
func (m *RICsubscriptionDeleteResponseT) XXX_Size() int {
	return xxx_messageInfo_RICsubscriptionDeleteResponseT.Size(m)
}
func (m *RICsubscriptionDeleteResponseT) XXX_DiscardUnknown() {
	xxx_messageInfo_RICsubscriptionDeleteResponseT.DiscardUnknown(m)
}

var xxx_messageInfo_RICsubscriptionDeleteResponseT proto.InternalMessageInfo

func (m *RICsubscriptionDeleteResponseT) GetProtocolIEs() *ProtocolIE_Container_1544P4T {
	if m != nil {
		return m.ProtocolIEs
	}
	return nil
}

func init() {
	proto.RegisterType((*RICsubscriptionDeleteResponseT)(nil), "e2ctypes.RICsubscriptionDeleteResponse_t")
}

func init() {
	proto.RegisterFile("RICsubscriptionDeleteResponse.proto", fileDescriptor_c6829107780d54e7)
}

var fileDescriptor_c6829107780d54e7 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x0e, 0xf2, 0x74, 0x2e,
	0x2e, 0x4d, 0x2a, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0xcc, 0xcf, 0x73, 0x49, 0xcd, 0x49, 0x2d,
	0x49, 0x0d, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x48, 0x35, 0x4a, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x96, 0x92, 0x0a, 0x00, 0x09, 0x24, 0xe7,
	0xe7, 0x78, 0xba, 0xea, 0x3a, 0xe7, 0xe7, 0x95, 0x24, 0x66, 0xe6, 0xa5, 0x16, 0x41, 0x54, 0x29,
	0x15, 0x70, 0xc9, 0xe3, 0x35, 0x2c, 0xbe, 0x44, 0xc8, 0x93, 0x8b, 0xbb, 0x00, 0x6e, 0x40, 0xb1,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xba, 0x1e, 0xcc, 0x78, 0x3d, 0x84, 0xe9, 0xf1, 0x70,
	0xd3, 0xe3, 0x0d, 0x4d, 0x4d, 0x4c, 0x02, 0x4c, 0xe2, 0x4b, 0x82, 0x90, 0xf5, 0x5a, 0x31, 0x49,
	0x30, 0x26, 0xb1, 0x81, 0x05, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x61, 0xc8, 0x31, 0xc2,
	0xc5, 0x00, 0x00, 0x00,
}