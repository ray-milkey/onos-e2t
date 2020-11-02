// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RICcontrolStatus.proto

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

type RICcontrolStatusT int32 // Deprecated: Do not use.
const (
	RICcontrolStatusT_RICcontrolStatus_success  RICcontrolStatusT = 0
	RICcontrolStatusT_RICcontrolStatus_rejected RICcontrolStatusT = 1
	RICcontrolStatusT_RICcontrolStatus_failed   RICcontrolStatusT = 2
)

var RICcontrolStatusT_name = map[int32]string{
	0: "RICcontrolStatus_success",
	1: "RICcontrolStatus_rejected",
	2: "RICcontrolStatus_failed",
}

var RICcontrolStatusT_value = map[string]int32{
	"RICcontrolStatus_success":  0,
	"RICcontrolStatus_rejected": 1,
	"RICcontrolStatus_failed":   2,
}

func (x RICcontrolStatusT) String() string {
	return proto.EnumName(RICcontrolStatusT_name, int32(x))
}

func (RICcontrolStatusT) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4c1454b9b047398b, []int{0}
}

func init() {
	proto.RegisterEnum("e2ctypes.RICcontrolStatusT", RICcontrolStatusT_name, RICcontrolStatusT_value)
}

func init() { proto.RegisterFile("RICcontrolStatus.proto", fileDescriptor_4c1454b9b047398b) }

var fileDescriptor_4c1454b9b047398b = []byte{
	// 120 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0b, 0xf2, 0x74, 0x4e,
	0xce, 0xcf, 0x2b, 0x29, 0xca, 0xcf, 0x09, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x35, 0x4a, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2a, 0xe2, 0x12,
	0x42, 0x57, 0x13, 0x5f, 0x22, 0x24, 0xc3, 0x25, 0x81, 0x21, 0x5a, 0x5c, 0x9a, 0x9c, 0x9c, 0x5a,
	0x5c, 0x2c, 0xc0, 0x20, 0x24, 0xcb, 0x25, 0x89, 0x21, 0x5b, 0x94, 0x9a, 0x95, 0x9a, 0x5c, 0x92,
	0x9a, 0x22, 0xc0, 0x28, 0x24, 0xcd, 0x25, 0x8e, 0x21, 0x9d, 0x96, 0x98, 0x99, 0x93, 0x9a, 0x22,
	0xc0, 0x24, 0xc5, 0x24, 0xc1, 0x98, 0xc4, 0x06, 0x76, 0x84, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0xf0, 0x40, 0x1e, 0xd8, 0x9e, 0x00, 0x00, 0x00,
}