// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/site.proto

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

type Site struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Page                 string   `protobuf:"bytes,3,opt,name=page,proto3" json:"page,omitempty"`
	Section              string   `protobuf:"bytes,4,opt,name=section,proto3" json:"section,omitempty"`
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Image                string   `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	Author               string   `protobuf:"bytes,7,opt,name=author,proto3" json:"author,omitempty"`
	User                 *User    `protobuf:"bytes,8,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Site) Reset()         { *m = Site{} }
func (m *Site) String() string { return proto.CompactTextString(m) }
func (*Site) ProtoMessage()    {}
func (*Site) Descriptor() ([]byte, []int) {
	return fileDescriptor_f64ecc3ac1e1b642, []int{0}
}

func (m *Site) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Site.Unmarshal(m, b)
}
func (m *Site) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Site.Marshal(b, m, deterministic)
}
func (m *Site) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Site.Merge(m, src)
}
func (m *Site) XXX_Size() int {
	return xxx_messageInfo_Site.Size(m)
}
func (m *Site) XXX_DiscardUnknown() {
	xxx_messageInfo_Site.DiscardUnknown(m)
}

var xxx_messageInfo_Site proto.InternalMessageInfo

func (m *Site) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Site) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Site) GetPage() string {
	if m != nil {
		return m.Page
	}
	return ""
}

func (m *Site) GetSection() string {
	if m != nil {
		return m.Section
	}
	return ""
}

func (m *Site) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Site) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Site) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Site) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*Site)(nil), "pb.Site")
}

func init() { proto.RegisterFile("pb/site.proto", fileDescriptor_f64ecc3ac1e1b642) }

var fileDescriptor_f64ecc3ac1e1b642 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xcf, 0x4d, 0x6a, 0xc3, 0x30,
	0x10, 0x05, 0x60, 0xe4, 0xca, 0x3f, 0x1d, 0xd3, 0xcd, 0x50, 0x8a, 0x28, 0x5d, 0x98, 0xae, 0xbc,
	0x72, 0x20, 0xb9, 0x89, 0x43, 0x0e, 0x60, 0x39, 0x43, 0x22, 0x88, 0x2d, 0x21, 0xc9, 0xb7, 0xcc,
	0xa1, 0x82, 0x46, 0x09, 0x64, 0x37, 0xef, 0x1b, 0xd0, 0xd3, 0xc0, 0x97, 0xd3, 0xbb, 0x60, 0x22,
	0x0d, 0xce, 0xdb, 0x68, 0xb1, 0x70, 0xfa, 0x37, 0xd1, 0x16, 0xc8, 0x67, 0xfa, 0xbf, 0x0b, 0x90,
	0x47, 0x13, 0x09, 0xbf, 0xa1, 0x8c, 0x26, 0xde, 0x48, 0x89, 0x4e, 0xf4, 0x9f, 0x63, 0x0e, 0x88,
	0x20, 0xd7, 0x69, 0x21, 0x55, 0x30, 0xf2, 0x9c, 0xcc, 0x4d, 0x17, 0x52, 0x1f, 0xd9, 0xd2, 0x8c,
	0x0a, 0xea, 0x40, 0x73, 0x34, 0x76, 0x55, 0x92, 0xf9, 0x15, 0xb1, 0x83, 0xf6, 0x4c, 0x61, 0xf6,
	0xc6, 0xf1, 0xb6, 0xe4, 0xed, 0x3b, 0xa5, 0x66, 0xb3, 0xa4, 0x07, 0xab, 0xdc, 0xcc, 0x01, 0x7f,
	0xa0, 0x9a, 0xb6, 0x78, 0xb5, 0x5e, 0xd5, 0xcc, 0xcf, 0x84, 0x7f, 0x20, 0xd3, 0xf7, 0x55, 0xd3,
	0x89, 0xbe, 0xdd, 0x37, 0x83, 0xd3, 0xc3, 0x29, 0x90, 0x1f, 0x59, 0x75, 0xc5, 0x57, 0x1d, 0x1e,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x33, 0x49, 0xdb, 0xf9, 0x00, 0x00, 0x00,
}