// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data.proto

package main

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

type ProtoImageRow struct {
	Values               []float32 `protobuf:"fixed32,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ProtoImageRow) Reset()         { *m = ProtoImageRow{} }
func (m *ProtoImageRow) String() string { return proto.CompactTextString(m) }
func (*ProtoImageRow) ProtoMessage()    {}
func (*ProtoImageRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}

func (m *ProtoImageRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoImageRow.Unmarshal(m, b)
}
func (m *ProtoImageRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoImageRow.Marshal(b, m, deterministic)
}
func (m *ProtoImageRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoImageRow.Merge(m, src)
}
func (m *ProtoImageRow) XXX_Size() int {
	return xxx_messageInfo_ProtoImageRow.Size(m)
}
func (m *ProtoImageRow) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoImageRow.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoImageRow proto.InternalMessageInfo

func (m *ProtoImageRow) GetValues() []float32 {
	if m != nil {
		return m.Values
	}
	return nil
}

type ProtoImageData struct {
	Rows                 []*ProtoImageRow `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ProtoImageData) Reset()         { *m = ProtoImageData{} }
func (m *ProtoImageData) String() string { return proto.CompactTextString(m) }
func (*ProtoImageData) ProtoMessage()    {}
func (*ProtoImageData) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{1}
}

func (m *ProtoImageData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoImageData.Unmarshal(m, b)
}
func (m *ProtoImageData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoImageData.Marshal(b, m, deterministic)
}
func (m *ProtoImageData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoImageData.Merge(m, src)
}
func (m *ProtoImageData) XXX_Size() int {
	return xxx_messageInfo_ProtoImageData.Size(m)
}
func (m *ProtoImageData) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoImageData.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoImageData proto.InternalMessageInfo

func (m *ProtoImageData) GetRows() []*ProtoImageRow {
	if m != nil {
		return m.Rows
	}
	return nil
}

type ProtoDWT struct {
	Mode                 string          `protobuf:"bytes,1,opt,name=mode,proto3" json:"mode,omitempty"`
	Data                 *ProtoImageData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ProtoDWT) Reset()         { *m = ProtoDWT{} }
func (m *ProtoDWT) String() string { return proto.CompactTextString(m) }
func (*ProtoDWT) ProtoMessage()    {}
func (*ProtoDWT) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{2}
}

func (m *ProtoDWT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoDWT.Unmarshal(m, b)
}
func (m *ProtoDWT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoDWT.Marshal(b, m, deterministic)
}
func (m *ProtoDWT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoDWT.Merge(m, src)
}
func (m *ProtoDWT) XXX_Size() int {
	return xxx_messageInfo_ProtoDWT.Size(m)
}
func (m *ProtoDWT) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoDWT.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoDWT proto.InternalMessageInfo

func (m *ProtoDWT) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *ProtoDWT) GetData() *ProtoImageData {
	if m != nil {
		return m.Data
	}
	return nil
}

type FileImageLayer struct {
	Rows                 [][]byte `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileImageLayer) Reset()         { *m = FileImageLayer{} }
func (m *FileImageLayer) String() string { return proto.CompactTextString(m) }
func (*FileImageLayer) ProtoMessage()    {}
func (*FileImageLayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{3}
}

func (m *FileImageLayer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileImageLayer.Unmarshal(m, b)
}
func (m *FileImageLayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileImageLayer.Marshal(b, m, deterministic)
}
func (m *FileImageLayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileImageLayer.Merge(m, src)
}
func (m *FileImageLayer) XXX_Size() int {
	return xxx_messageInfo_FileImageLayer.Size(m)
}
func (m *FileImageLayer) XXX_DiscardUnknown() {
	xxx_messageInfo_FileImageLayer.DiscardUnknown(m)
}

var xxx_messageInfo_FileImageLayer proto.InternalMessageInfo

func (m *FileImageLayer) GetRows() [][]byte {
	if m != nil {
		return m.Rows
	}
	return nil
}

type FileImageHeader struct {
	Width                uint32   `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileImageHeader) Reset()         { *m = FileImageHeader{} }
func (m *FileImageHeader) String() string { return proto.CompactTextString(m) }
func (*FileImageHeader) ProtoMessage()    {}
func (*FileImageHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{4}
}

func (m *FileImageHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileImageHeader.Unmarshal(m, b)
}
func (m *FileImageHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileImageHeader.Marshal(b, m, deterministic)
}
func (m *FileImageHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileImageHeader.Merge(m, src)
}
func (m *FileImageHeader) XXX_Size() int {
	return xxx_messageInfo_FileImageHeader.Size(m)
}
func (m *FileImageHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_FileImageHeader.DiscardUnknown(m)
}

var xxx_messageInfo_FileImageHeader proto.InternalMessageInfo

func (m *FileImageHeader) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *FileImageHeader) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

type FileImageData struct {
	Y                    *FileImageLayer `protobuf:"bytes,1,opt,name=y,proto3" json:"y,omitempty"`
	U                    *FileImageLayer `protobuf:"bytes,2,opt,name=u,proto3" json:"u,omitempty"`
	V                    *FileImageLayer `protobuf:"bytes,3,opt,name=v,proto3" json:"v,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FileImageData) Reset()         { *m = FileImageData{} }
func (m *FileImageData) String() string { return proto.CompactTextString(m) }
func (*FileImageData) ProtoMessage()    {}
func (*FileImageData) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{5}
}

func (m *FileImageData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileImageData.Unmarshal(m, b)
}
func (m *FileImageData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileImageData.Marshal(b, m, deterministic)
}
func (m *FileImageData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileImageData.Merge(m, src)
}
func (m *FileImageData) XXX_Size() int {
	return xxx_messageInfo_FileImageData.Size(m)
}
func (m *FileImageData) XXX_DiscardUnknown() {
	xxx_messageInfo_FileImageData.DiscardUnknown(m)
}

var xxx_messageInfo_FileImageData proto.InternalMessageInfo

func (m *FileImageData) GetY() *FileImageLayer {
	if m != nil {
		return m.Y
	}
	return nil
}

func (m *FileImageData) GetU() *FileImageLayer {
	if m != nil {
		return m.U
	}
	return nil
}

func (m *FileImageData) GetV() *FileImageLayer {
	if m != nil {
		return m.V
	}
	return nil
}

type FileImage struct {
	Header               *FileImageHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data                 *FileImageData   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *FileImage) Reset()         { *m = FileImage{} }
func (m *FileImage) String() string { return proto.CompactTextString(m) }
func (*FileImage) ProtoMessage()    {}
func (*FileImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{6}
}

func (m *FileImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileImage.Unmarshal(m, b)
}
func (m *FileImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileImage.Marshal(b, m, deterministic)
}
func (m *FileImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileImage.Merge(m, src)
}
func (m *FileImage) XXX_Size() int {
	return xxx_messageInfo_FileImage.Size(m)
}
func (m *FileImage) XXX_DiscardUnknown() {
	xxx_messageInfo_FileImage.DiscardUnknown(m)
}

var xxx_messageInfo_FileImage proto.InternalMessageInfo

func (m *FileImage) GetHeader() *FileImageHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *FileImage) GetData() *FileImageData {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtoImageRow)(nil), "main.ProtoImageRow")
	proto.RegisterType((*ProtoImageData)(nil), "main.ProtoImageData")
	proto.RegisterType((*ProtoDWT)(nil), "main.ProtoDWT")
	proto.RegisterType((*FileImageLayer)(nil), "main.FileImageLayer")
	proto.RegisterType((*FileImageHeader)(nil), "main.FileImageHeader")
	proto.RegisterType((*FileImageData)(nil), "main.FileImageData")
	proto.RegisterType((*FileImage)(nil), "main.FileImage")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor_871986018790d2fd) }

var fileDescriptor_871986018790d2fd = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xc9, 0x56, 0x87, 0x7b, 0xb3, 0x13, 0xb2, 0x29, 0x3d, 0x96, 0x20, 0xac, 0x17, 0x7b,
	0x98, 0x27, 0x4f, 0x5e, 0x86, 0x4c, 0xf0, 0x20, 0x41, 0xf0, 0x1c, 0x6d, 0x58, 0x0b, 0xab, 0x91,
	0x9a, 0xb6, 0x14, 0xff, 0x79, 0x79, 0xaf, 0x21, 0xb3, 0x3d, 0x78, 0xeb, 0x7b, 0x7c, 0xf2, 0xfd,
	0x91, 0x06, 0x20, 0x53, 0x56, 0xa5, 0x5f, 0x95, 0xb1, 0x86, 0x07, 0xa5, 0x2a, 0x3e, 0xc5, 0x06,
	0xc2, 0x17, 0x1c, 0x9f, 0x4a, 0x75, 0xd0, 0xd2, 0xb4, 0xfc, 0x1a, 0x66, 0x8d, 0x3a, 0xd6, 0xfa,
	0x3b, 0x62, 0xf1, 0x34, 0x99, 0x48, 0x37, 0x89, 0x7b, 0x58, 0x9e, 0xc0, 0x9d, 0xb2, 0x8a, 0x6f,
	0x20, 0xa8, 0x4c, 0xdb, 0x73, 0x8b, 0xed, 0x2a, 0x45, 0xbd, 0x74, 0x20, 0x26, 0x09, 0x10, 0x7b,
	0x38, 0xa7, 0xf5, 0xee, 0xed, 0x95, 0x73, 0x08, 0x4a, 0x93, 0xe9, 0x88, 0xc5, 0x2c, 0x99, 0x4b,
	0xfa, 0xe6, 0x09, 0x04, 0x98, 0x2b, 0x9a, 0xc4, 0x2c, 0x59, 0x6c, 0xd7, 0x63, 0x21, 0x34, 0x93,
	0x44, 0x88, 0x1b, 0x58, 0x3e, 0x16, 0x47, 0x4d, 0xeb, 0x67, 0xd5, 0xe9, 0x0a, 0xf5, 0x7c, 0x88,
	0x0b, 0xe7, 0xf7, 0x00, 0x97, 0x9e, 0xda, 0x6b, 0x95, 0xe9, 0x8a, 0xaf, 0xe1, 0xac, 0x2d, 0x32,
	0x9b, 0x93, 0x6f, 0x28, 0xfb, 0x01, 0xbb, 0xe6, 0xba, 0x38, 0xe4, 0x96, 0xac, 0x43, 0xe9, 0x26,
	0xf1, 0x03, 0xa1, 0x17, 0xa0, 0xaa, 0x02, 0x58, 0x47, 0x47, 0x7d, 0xbc, 0x61, 0x0c, 0xc9, 0x3a,
	0x64, 0xea, 0x61, 0x85, 0x31, 0x53, 0x23, 0xd3, 0x44, 0xd3, 0xff, 0x98, 0x46, 0x7c, 0xc0, 0xdc,
	0x2f, 0xf9, 0x2d, 0x26, 0xc4, 0x06, 0xce, 0xfd, 0x6a, 0x74, 0xaa, 0xaf, 0x27, 0x1d, 0x84, 0xbf,
	0xe4, 0xcf, 0x4d, 0xae, 0x46, 0xf0, 0xe9, 0x22, 0xdf, 0x67, 0xf4, 0x06, 0xee, 0x7e, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x0a, 0x27, 0x35, 0x8b, 0x11, 0x02, 0x00, 0x00,
}
