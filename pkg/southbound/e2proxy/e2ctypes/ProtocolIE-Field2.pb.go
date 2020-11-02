// Code generated by protoc-gen-go. DO NOT EDIT.
// ProtocolIE-Field2.proto is a deprecated file.

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
type RICaction_ToBeSetup_ItemIEsT struct {
	Id          ProtocolIE_IDT `protobuf:"varint,1,opt,name=id,proto3,enum=e2ctypes.ProtocolIE_IDT" json:"id,omitempty"`                 // Deprecated: Do not use.
	Criticality CriticalityT   `protobuf:"varint,2,opt,name=criticality,proto3,enum=e2ctypes.CriticalityT" json:"criticality,omitempty"` // Deprecated: Do not use.
	// Types that are valid to be assigned to Choice:
	//	*RICaction_ToBeSetup_ItemIEsT_RICaction_ToBeSetup_Item
	Choice               isRICaction_ToBeSetup_ItemIEsT_Choice `protobuf_oneof:"choice"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *RICaction_ToBeSetup_ItemIEsT) Reset()         { *m = RICaction_ToBeSetup_ItemIEsT{} }
func (m *RICaction_ToBeSetup_ItemIEsT) String() string { return proto.CompactTextString(m) }
func (*RICaction_ToBeSetup_ItemIEsT) ProtoMessage()    {}
func (*RICaction_ToBeSetup_ItemIEsT) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f996f3f8c5383e0, []int{0}
}

func (m *RICaction_ToBeSetup_ItemIEsT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RICaction_ToBeSetup_ItemIEsT.Unmarshal(m, b)
}
func (m *RICaction_ToBeSetup_ItemIEsT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RICaction_ToBeSetup_ItemIEsT.Marshal(b, m, deterministic)
}
func (m *RICaction_ToBeSetup_ItemIEsT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RICaction_ToBeSetup_ItemIEsT.Merge(m, src)
}
func (m *RICaction_ToBeSetup_ItemIEsT) XXX_Size() int {
	return xxx_messageInfo_RICaction_ToBeSetup_ItemIEsT.Size(m)
}
func (m *RICaction_ToBeSetup_ItemIEsT) XXX_DiscardUnknown() {
	xxx_messageInfo_RICaction_ToBeSetup_ItemIEsT.DiscardUnknown(m)
}

var xxx_messageInfo_RICaction_ToBeSetup_ItemIEsT proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *RICaction_ToBeSetup_ItemIEsT) GetId() ProtocolIE_IDT {
	if m != nil {
		return m.Id
	}
	return ProtocolIE_IDT_ProtocolIE_ID_id_Dummy
}

// Deprecated: Do not use.
func (m *RICaction_ToBeSetup_ItemIEsT) GetCriticality() CriticalityT {
	if m != nil {
		return m.Criticality
	}
	return CriticalityT_Criticality_reject
}

type isRICaction_ToBeSetup_ItemIEsT_Choice interface {
	isRICaction_ToBeSetup_ItemIEsT_Choice()
}

type RICaction_ToBeSetup_ItemIEsT_RICaction_ToBeSetup_Item struct {
	RICaction_ToBeSetup_Item *RICaction_ToBeSetup_ItemT `protobuf:"bytes,3,opt,name=RICaction_ToBeSetup_Item,json=RICactionToBeSetupItem,proto3,oneof"`
}

func (*RICaction_ToBeSetup_ItemIEsT_RICaction_ToBeSetup_Item) isRICaction_ToBeSetup_ItemIEsT_Choice() {
}

func (m *RICaction_ToBeSetup_ItemIEsT) GetChoice() isRICaction_ToBeSetup_ItemIEsT_Choice {
	if m != nil {
		return m.Choice
	}
	return nil
}

// Deprecated: Do not use.
func (m *RICaction_ToBeSetup_ItemIEsT) GetRICaction_ToBeSetup_Item() *RICaction_ToBeSetup_ItemT {
	if x, ok := m.GetChoice().(*RICaction_ToBeSetup_ItemIEsT_RICaction_ToBeSetup_Item); ok {
		return x.RICaction_ToBeSetup_Item
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RICaction_ToBeSetup_ItemIEsT) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RICaction_ToBeSetup_ItemIEsT_RICaction_ToBeSetup_Item)(nil),
	}
}

func init() {
	proto.RegisterType((*RICaction_ToBeSetup_ItemIEsT)(nil), "e2ctypes.RICaction_ToBeSetup_ItemIEs_t")
}

func init() { proto.RegisterFile("ProtocolIE-Field2.proto", fileDescriptor_0f996f3f8c5383e0) }

var fileDescriptor_0f996f3f8c5383e0 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0f, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0xf1, 0x74, 0xd5, 0x75, 0xcb, 0x4c, 0xcd, 0x49, 0x31, 0xd2, 0x2b, 0x00,
	0x89, 0x08, 0x71, 0xa4, 0x1a, 0x25, 0x97, 0x54, 0x16, 0xa4, 0x16, 0x4b, 0x09, 0x23, 0x29, 0xf1,
	0x74, 0x81, 0x48, 0x4b, 0x09, 0x3a, 0x17, 0x65, 0x96, 0x64, 0x26, 0x27, 0xe6, 0x64, 0x96, 0x54,
	0x42, 0x85, 0xe4, 0x82, 0x3c, 0x9d, 0x13, 0x93, 0x4b, 0x32, 0xf3, 0xf3, 0x74, 0x43, 0xf2, 0x9d,
	0x52, 0x83, 0x53, 0x4b, 0x4a, 0x0b, 0x74, 0x3d, 0x4b, 0x52, 0x73, 0x21, 0xf2, 0x4a, 0x8d, 0x4c,
	0x5c, 0xb2, 0x70, 0x25, 0xf1, 0x70, 0x25, 0xf1, 0x20, 0x25, 0x9e, 0xae, 0xc5, 0xf1, 0x25, 0x42,
	0xba, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x7c, 0x46, 0x92, 0x7a, 0x30, 0x07,
	0xe8, 0x21, 0xec, 0x8f, 0xf7, 0x74, 0x89, 0x2f, 0x71, 0x62, 0x92, 0x60, 0x0c, 0x62, 0xca, 0x4c,
	0x11, 0xb2, 0xe5, 0xe2, 0x4e, 0x46, 0xb8, 0x42, 0x82, 0x09, 0xac, 0x4f, 0x1c, 0xa1, 0x0f, 0xc9,
	0x89, 0x50, 0x5d, 0xc8, 0xea, 0x85, 0x52, 0xb8, 0x24, 0x70, 0x39, 0x47, 0x82, 0x59, 0x81, 0x51,
	0x83, 0xdb, 0x48, 0x05, 0x61, 0x16, 0x2e, 0x95, 0x10, 0x83, 0x3d, 0x18, 0x82, 0xc4, 0xe0, 0x2a,
	0xe0, 0x0a, 0x40, 0xf2, 0x56, 0x4c, 0x12, 0x8c, 0x4e, 0x1c, 0x5c, 0x6c, 0xc9, 0x19, 0xf9, 0x99,
	0xc9, 0xa9, 0x4e, 0xcc, 0x3b, 0x18, 0x19, 0x93, 0xd8, 0xc0, 0xe1, 0x61, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xe1, 0x79, 0x2c, 0x1d, 0x7c, 0x01, 0x00, 0x00,
}