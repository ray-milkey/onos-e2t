// Code generated by protoc-gen-go. DO NOT EDIT.
// Criticality.proto is a deprecated file.

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

type CriticalityT int32

const (
	CriticalityT_Criticality_reject CriticalityT = 0 // Deprecated: Do not use.
	CriticalityT_Criticality_ignore CriticalityT = 1 // Deprecated: Do not use.
	CriticalityT_Criticality_notify CriticalityT = 2 // Deprecated: Do not use.
)

var CriticalityT_name = map[int32]string{
	0: "Criticality_reject",
	1: "Criticality_ignore",
	2: "Criticality_notify",
}

var CriticalityT_value = map[string]int32{
	"Criticality_reject": 0,
	"Criticality_ignore": 1,
	"Criticality_notify": 2,
}

func (x CriticalityT) String() string {
	return proto.EnumName(CriticalityT_name, int32(x))
}

func (CriticalityT) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_08953a4dae088f7f, []int{0}
}

func init() {
	proto.RegisterEnum("e2ctypes.CriticalityT", CriticalityT_name, CriticalityT_value)
}

func init() { proto.RegisterFile("Criticality.proto", fileDescriptor_08953a4dae088f7f) }

var fileDescriptor_08953a4dae088f7f = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x74, 0x2e, 0xca, 0x2c,
	0xc9, 0x4c, 0x4e, 0xcc, 0xc9, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48,
	0x35, 0x4a, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4a, 0xe6, 0xe2, 0x45, 0x92, 0x8e, 0x2f, 0x11,
	0x92, 0xe2, 0x12, 0x42, 0x16, 0x28, 0x4a, 0xcd, 0x4a, 0x4d, 0x2e, 0x11, 0x60, 0x90, 0x62, 0xe2,
	0x60, 0x44, 0x97, 0xcb, 0x4c, 0xcf, 0xcb, 0x2f, 0x4a, 0x15, 0x60, 0xc4, 0x26, 0x97, 0x97, 0x5f,
	0x92, 0x99, 0x56, 0x29, 0xc0, 0x04, 0x92, 0x73, 0x62, 0xde, 0xc1, 0xc8, 0x98, 0xc4, 0x06, 0xb6,
	0xda, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x1c, 0x52, 0x8e, 0x8f, 0x00, 0x00, 0x00,
}