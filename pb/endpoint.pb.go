// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/endpoint.proto

package pb

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Endpoint struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Port                 int64    `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef2d3517b7d1a853, []int{0}
}

func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Endpoint) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Endpoint) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Endpoint) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "pb.Endpoint")
}

func init() { proto.RegisterFile("pb/endpoint.proto", fileDescriptor_ef2d3517b7d1a853) }

var fileDescriptor_ef2d3517b7d1a853 = []byte{
	// 120 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x48, 0xd2, 0x4f,
	0xcd, 0x4b, 0x29, 0xc8, 0xcf, 0xcc, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a,
	0x48, 0x52, 0x4a, 0xe2, 0xe2, 0x70, 0x85, 0x8a, 0x0a, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6,
	0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x20, 0xb1, 0x8c, 0xfc, 0xe2, 0x12, 0x09,
	0x26, 0x88, 0x18, 0x88, 0x2d, 0x24, 0xc5, 0xc5, 0x51, 0x90, 0x58, 0x5c, 0x5c, 0x9e, 0x5f, 0x94,
	0x22, 0xc1, 0x0c, 0x16, 0x87, 0xf3, 0x41, 0xea, 0x0b, 0xf2, 0x8b, 0x4a, 0x24, 0x58, 0x14, 0x18,
	0x35, 0x98, 0x83, 0xc0, 0xec, 0x24, 0x36, 0xb0, 0x75, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xfd, 0x6e, 0x40, 0xe1, 0x83, 0x00, 0x00, 0x00,
}
