// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: embed_types.proto

package protocol

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

type EmbedImage struct {
	Url                  *string  `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Width                *uint32  `protobuf:"varint,2,opt,name=width" json:"width,omitempty"`
	Height               *uint32  `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmbedImage) Reset()         { *m = EmbedImage{} }
func (m *EmbedImage) String() string { return proto.CompactTextString(m) }
func (*EmbedImage) ProtoMessage()    {}
func (*EmbedImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a65fcf161f3afa9f, []int{0}
}

func (m *EmbedImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmbedImage.Unmarshal(m, b)
}
func (m *EmbedImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmbedImage.Marshal(b, m, deterministic)
}
func (m *EmbedImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmbedImage.Merge(m, src)
}
func (m *EmbedImage) XXX_Size() int {
	return xxx_messageInfo_EmbedImage.Size(m)
}
func (m *EmbedImage) XXX_DiscardUnknown() {
	xxx_messageInfo_EmbedImage.DiscardUnknown(m)
}

var xxx_messageInfo_EmbedImage proto.InternalMessageInfo

func (m *EmbedImage) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *EmbedImage) GetWidth() uint32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *EmbedImage) GetHeight() uint32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

type Provider struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Provider) Reset()         { *m = Provider{} }
func (m *Provider) String() string { return proto.CompactTextString(m) }
func (*Provider) ProtoMessage()    {}
func (*Provider) Descriptor() ([]byte, []int) {
	return fileDescriptor_a65fcf161f3afa9f, []int{1}
}

func (m *Provider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Provider.Unmarshal(m, b)
}
func (m *Provider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Provider.Marshal(b, m, deterministic)
}
func (m *Provider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Provider.Merge(m, src)
}
func (m *Provider) XXX_Size() int {
	return xxx_messageInfo_Provider.Size(m)
}
func (m *Provider) XXX_DiscardUnknown() {
	xxx_messageInfo_Provider.DiscardUnknown(m)
}

var xxx_messageInfo_Provider proto.InternalMessageInfo

func (m *Provider) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type Favicon struct {
	Url                  *string  `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Favicon) Reset()         { *m = Favicon{} }
func (m *Favicon) String() string { return proto.CompactTextString(m) }
func (*Favicon) ProtoMessage()    {}
func (*Favicon) Descriptor() ([]byte, []int) {
	return fileDescriptor_a65fcf161f3afa9f, []int{2}
}

func (m *Favicon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Favicon.Unmarshal(m, b)
}
func (m *Favicon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Favicon.Marshal(b, m, deterministic)
}
func (m *Favicon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Favicon.Merge(m, src)
}
func (m *Favicon) XXX_Size() int {
	return xxx_messageInfo_Favicon.Size(m)
}
func (m *Favicon) XXX_DiscardUnknown() {
	xxx_messageInfo_Favicon.DiscardUnknown(m)
}

var xxx_messageInfo_Favicon proto.InternalMessageInfo

func (m *Favicon) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

type EmbedHTML struct {
	Content              *string  `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
	Width                *uint32  `protobuf:"varint,2,opt,name=width" json:"width,omitempty"`
	Height               *uint32  `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmbedHTML) Reset()         { *m = EmbedHTML{} }
func (m *EmbedHTML) String() string { return proto.CompactTextString(m) }
func (*EmbedHTML) ProtoMessage()    {}
func (*EmbedHTML) Descriptor() ([]byte, []int) {
	return fileDescriptor_a65fcf161f3afa9f, []int{3}
}

func (m *EmbedHTML) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmbedHTML.Unmarshal(m, b)
}
func (m *EmbedHTML) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmbedHTML.Marshal(b, m, deterministic)
}
func (m *EmbedHTML) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmbedHTML.Merge(m, src)
}
func (m *EmbedHTML) XXX_Size() int {
	return xxx_messageInfo_EmbedHTML.Size(m)
}
func (m *EmbedHTML) XXX_DiscardUnknown() {
	xxx_messageInfo_EmbedHTML.DiscardUnknown(m)
}

var xxx_messageInfo_EmbedHTML proto.InternalMessageInfo

func (m *EmbedHTML) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func (m *EmbedHTML) GetWidth() uint32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *EmbedHTML) GetHeight() uint32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

type EmbedInfo struct {
	Title                *string     `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Type                 *string     `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	OriginalUrl          *string     `protobuf:"bytes,3,opt,name=original_url,json=originalUrl" json:"original_url,omitempty"`
	Thumbnail            *EmbedImage `protobuf:"bytes,4,opt,name=thumbnail" json:"thumbnail,omitempty"`
	Provider             *Provider   `protobuf:"bytes,5,opt,name=provider" json:"provider,omitempty"`
	Description          *string     `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	Favicon              *Favicon    `protobuf:"bytes,7,opt,name=favicon" json:"favicon,omitempty"`
	Html                 *EmbedHTML  `protobuf:"bytes,8,opt,name=html" json:"html,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EmbedInfo) Reset()         { *m = EmbedInfo{} }
func (m *EmbedInfo) String() string { return proto.CompactTextString(m) }
func (*EmbedInfo) ProtoMessage()    {}
func (*EmbedInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a65fcf161f3afa9f, []int{4}
}

func (m *EmbedInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmbedInfo.Unmarshal(m, b)
}
func (m *EmbedInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmbedInfo.Marshal(b, m, deterministic)
}
func (m *EmbedInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmbedInfo.Merge(m, src)
}
func (m *EmbedInfo) XXX_Size() int {
	return xxx_messageInfo_EmbedInfo.Size(m)
}
func (m *EmbedInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_EmbedInfo.DiscardUnknown(m)
}

var xxx_messageInfo_EmbedInfo proto.InternalMessageInfo

func (m *EmbedInfo) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *EmbedInfo) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *EmbedInfo) GetOriginalUrl() string {
	if m != nil && m.OriginalUrl != nil {
		return *m.OriginalUrl
	}
	return ""
}

func (m *EmbedInfo) GetThumbnail() *EmbedImage {
	if m != nil {
		return m.Thumbnail
	}
	return nil
}

func (m *EmbedInfo) GetProvider() *Provider {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (m *EmbedInfo) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *EmbedInfo) GetFavicon() *Favicon {
	if m != nil {
		return m.Favicon
	}
	return nil
}

func (m *EmbedInfo) GetHtml() *EmbedHTML {
	if m != nil {
		return m.Html
	}
	return nil
}

func init() {
	proto.RegisterType((*EmbedImage)(nil), "bgs.protocol.EmbedImage")
	proto.RegisterType((*Provider)(nil), "bgs.protocol.Provider")
	proto.RegisterType((*Favicon)(nil), "bgs.protocol.Favicon")
	proto.RegisterType((*EmbedHTML)(nil), "bgs.protocol.EmbedHTML")
	proto.RegisterType((*EmbedInfo)(nil), "bgs.protocol.EmbedInfo")
}

func init() { proto.RegisterFile("embed_types.proto", fileDescriptor_a65fcf161f3afa9f) }

var fileDescriptor_a65fcf161f3afa9f = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xdd, 0x4a, 0xeb, 0x40,
	0x14, 0x85, 0x49, 0xff, 0xd2, 0xec, 0xf6, 0xc0, 0x39, 0xc3, 0xb1, 0x0e, 0x08, 0x12, 0x73, 0x55,
	0x10, 0x92, 0x52, 0xc4, 0x07, 0x10, 0x14, 0x0b, 0x15, 0x24, 0xea, 0x8d, 0x37, 0x25, 0x3f, 0xd3,
	0x64, 0x20, 0x99, 0x09, 0x93, 0x9d, 0x8a, 0xef, 0xe7, 0x83, 0x49, 0x26, 0x89, 0x6d, 0xd1, 0x1b,
	0xef, 0xf6, 0xda, 0x59, 0x7b, 0x65, 0xf1, 0x0d, 0xfc, 0x63, 0x79, 0xc8, 0xe2, 0x0d, 0xbe, 0x17,
	0xac, 0x74, 0x0b, 0x25, 0x51, 0x92, 0x69, 0x98, 0xb4, 0x63, 0x24, 0x33, 0x67, 0x0d, 0x70, 0x5b,
	0x5b, 0x56, 0x79, 0x90, 0x30, 0xf2, 0x17, 0xfa, 0x95, 0xca, 0xa8, 0x61, 0x1b, 0x73, 0xcb, 0xaf,
	0x47, 0xf2, 0x1f, 0x86, 0x6f, 0x3c, 0xc6, 0x94, 0xf6, 0x6c, 0x63, 0xfe, 0xc7, 0x6f, 0x04, 0x99,
	0xc1, 0x28, 0x65, 0x3c, 0x49, 0x91, 0xf6, 0xf5, 0xba, 0x55, 0xce, 0x39, 0x8c, 0x1f, 0x95, 0xdc,
	0xf1, 0x98, 0x29, 0x42, 0x60, 0x20, 0x82, 0x9c, 0xb5, 0x61, 0x7a, 0x76, 0xce, 0xc0, 0xbc, 0x0b,
	0x76, 0x3c, 0x92, 0xe2, 0xfb, 0xaf, 0x9c, 0x27, 0xb0, 0x74, 0x95, 0xfb, 0xe7, 0x87, 0x35, 0xa1,
	0x60, 0x46, 0x52, 0x20, 0x13, 0xd8, 0x5a, 0x3a, 0xf9, 0xcb, 0x46, 0x1f, 0xbd, 0x36, 0x75, 0x25,
	0xb6, 0xb2, 0xbe, 0x45, 0x8e, 0x59, 0x57, 0xaa, 0x11, 0x75, 0xd3, 0x1a, 0x90, 0x0e, 0xb4, 0x7c,
	0x3d, 0x93, 0x0b, 0x98, 0x4a, 0xc5, 0x13, 0x2e, 0x82, 0x6c, 0x53, 0xf7, 0xec, 0xeb, 0x6f, 0x93,
	0x6e, 0xf7, 0xa2, 0x32, 0x72, 0x0d, 0x16, 0xa6, 0x55, 0x1e, 0x8a, 0x80, 0x67, 0x74, 0x60, 0x1b,
	0xf3, 0xc9, 0x92, 0xba, 0x87, 0x70, 0xdd, 0x3d, 0x59, 0x7f, 0x6f, 0x25, 0x4b, 0x18, 0x17, 0x2d,
	0x24, 0x3a, 0xd4, 0x67, 0xb3, 0xe3, 0xb3, 0x0e, 0xa1, 0xff, 0xe5, 0x23, 0x36, 0x4c, 0x62, 0x56,
	0x46, 0x8a, 0x17, 0xc8, 0xa5, 0xa0, 0xa3, 0xa6, 0xcd, 0xc1, 0x8a, 0x78, 0x60, 0x6e, 0x1b, 0xb4,
	0xd4, 0xd4, 0xa1, 0x27, 0xc7, 0xa1, 0x2d, 0x77, 0xbf, 0x73, 0x91, 0x4b, 0x18, 0xa4, 0x98, 0x67,
	0x74, 0xac, 0xdd, 0xa7, 0x3f, 0x34, 0xaf, 0x1f, 0xc2, 0xd7, 0xa6, 0x9b, 0xab, 0xd7, 0x65, 0xc2,
	0x31, 0xad, 0x42, 0x37, 0x92, 0xb9, 0x57, 0x56, 0x05, 0x53, 0xc5, 0x62, 0x81, 0x5e, 0x22, 0x8b,
	0x94, 0xa9, 0x48, 0x05, 0x5b, 0xf4, 0x42, 0xc1, 0xd0, 0x0b, 0x93, 0xd2, 0xeb, 0x52, 0x3e, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xf4, 0x4e, 0x68, 0x46, 0x7e, 0x02, 0x00, 0x00,
}