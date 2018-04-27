// Code generated by protoc-gen-go. DO NOT EDIT.
// source: EventTypePb.proto

package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EventTypePb int32

const (
	EventTypePb_REGISTER   EventTypePb = 0
	EventTypePb_UNREGISTER EventTypePb = 1
)

var EventTypePb_name = map[int32]string{
	0: "REGISTER",
	1: "UNREGISTER",
}
var EventTypePb_value = map[string]int32{
	"REGISTER":   0,
	"UNREGISTER": 1,
}

func (x EventTypePb) String() string {
	return proto.EnumName(EventTypePb_name, int32(x))
}
func (EventTypePb) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func init() {
	proto.RegisterEnum("EventTypePb", EventTypePb_name, EventTypePb_value)
}

func init() { proto.RegisterFile("EventTypePb.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x74, 0x2d, 0x4b, 0xcd,
	0x2b, 0x09, 0xa9, 0x2c, 0x48, 0x0d, 0x48, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xd2, 0xe6,
	0xe2, 0x46, 0x12, 0x14, 0xe2, 0xe1, 0xe2, 0x08, 0x72, 0x75, 0xf7, 0x0c, 0x0e, 0x71, 0x0d, 0x12,
	0x60, 0x10, 0xe2, 0xe3, 0xe2, 0x0a, 0xf5, 0x83, 0xf3, 0x19, 0x9d, 0x0c, 0xb8, 0x34, 0x92, 0xf3,
	0x73, 0xf5, 0x12, 0x73, 0x32, 0x0b, 0x12, 0x2b, 0xf5, 0x8a, 0xf3, 0xd3, 0x12, 0xf5, 0x8a, 0x52,
	0xd3, 0x33, 0x8b, 0x4b, 0x8a, 0x2a, 0xf5, 0x8a, 0x53, 0x8b, 0xca, 0x52, 0x8b, 0xf4, 0x72, 0xf3,
	0x53, 0x52, 0x73, 0xf4, 0x0a, 0x92, 0x02, 0x18, 0xa3, 0x98, 0x0a, 0x92, 0x92, 0xd8, 0xc0, 0xb6,
	0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x20, 0xb8, 0x6e, 0x5b, 0x7a, 0x00, 0x00, 0x00,
}