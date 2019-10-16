// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data.proto

package data

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

type Subsampling int32

const (
	Subsampling_SUBSAMPLING_410 Subsampling = 0
	Subsampling_SUBSAMPLING_420 Subsampling = 1
	Subsampling_SUBSAMPLING_422 Subsampling = 2
	Subsampling_SUBSAMPLING_444 Subsampling = 3
)

var Subsampling_name = map[int32]string{
	0: "SUBSAMPLING_410",
	1: "SUBSAMPLING_420",
	2: "SUBSAMPLING_422",
	3: "SUBSAMPLING_444",
}

var Subsampling_value = map[string]int32{
	"SUBSAMPLING_410": 0,
	"SUBSAMPLING_420": 1,
	"SUBSAMPLING_422": 2,
	"SUBSAMPLING_444": 3,
}

func (x Subsampling) String() string {
	return proto.EnumName(Subsampling_name, int32(x))
}

func (Subsampling) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}

type ImageRow struct {
	Values               []float32 `protobuf:"fixed32,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ImageRow) Reset()         { *m = ImageRow{} }
func (m *ImageRow) String() string { return proto.CompactTextString(m) }
func (*ImageRow) ProtoMessage()    {}
func (*ImageRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}

func (m *ImageRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageRow.Unmarshal(m, b)
}
func (m *ImageRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageRow.Marshal(b, m, deterministic)
}
func (m *ImageRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageRow.Merge(m, src)
}
func (m *ImageRow) XXX_Size() int {
	return xxx_messageInfo_ImageRow.Size(m)
}
func (m *ImageRow) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageRow.DiscardUnknown(m)
}

var xxx_messageInfo_ImageRow proto.InternalMessageInfo

func (m *ImageRow) GetValues() []float32 {
	if m != nil {
		return m.Values
	}
	return nil
}

type ImageData struct {
	Rows                 []*ImageRow `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ImageData) Reset()         { *m = ImageData{} }
func (m *ImageData) String() string { return proto.CompactTextString(m) }
func (*ImageData) ProtoMessage()    {}
func (*ImageData) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{1}
}

func (m *ImageData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageData.Unmarshal(m, b)
}
func (m *ImageData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageData.Marshal(b, m, deterministic)
}
func (m *ImageData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageData.Merge(m, src)
}
func (m *ImageData) XXX_Size() int {
	return xxx_messageInfo_ImageData.Size(m)
}
func (m *ImageData) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageData.DiscardUnknown(m)
}

var xxx_messageInfo_ImageData proto.InternalMessageInfo

func (m *ImageData) GetRows() []*ImageRow {
	if m != nil {
		return m.Rows
	}
	return nil
}

type PythonDWT struct {
	Mode                 string     `protobuf:"bytes,1,opt,name=mode,proto3" json:"mode,omitempty"`
	Data                 *ImageData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PythonDWT) Reset()         { *m = PythonDWT{} }
func (m *PythonDWT) String() string { return proto.CompactTextString(m) }
func (*PythonDWT) ProtoMessage()    {}
func (*PythonDWT) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{2}
}

func (m *PythonDWT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PythonDWT.Unmarshal(m, b)
}
func (m *PythonDWT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PythonDWT.Marshal(b, m, deterministic)
}
func (m *PythonDWT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PythonDWT.Merge(m, src)
}
func (m *PythonDWT) XXX_Size() int {
	return xxx_messageInfo_PythonDWT.Size(m)
}
func (m *PythonDWT) XXX_DiscardUnknown() {
	xxx_messageInfo_PythonDWT.DiscardUnknown(m)
}

var xxx_messageInfo_PythonDWT proto.InternalMessageInfo

func (m *PythonDWT) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *PythonDWT) GetData() *ImageData {
	if m != nil {
		return m.Data
	}
	return nil
}

type WaveletHaar struct {
	Level                uint32   `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WaveletHaar) Reset()         { *m = WaveletHaar{} }
func (m *WaveletHaar) String() string { return proto.CompactTextString(m) }
func (*WaveletHaar) ProtoMessage()    {}
func (*WaveletHaar) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{3}
}

func (m *WaveletHaar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WaveletHaar.Unmarshal(m, b)
}
func (m *WaveletHaar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WaveletHaar.Marshal(b, m, deterministic)
}
func (m *WaveletHaar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaveletHaar.Merge(m, src)
}
func (m *WaveletHaar) XXX_Size() int {
	return xxx_messageInfo_WaveletHaar.Size(m)
}
func (m *WaveletHaar) XXX_DiscardUnknown() {
	xxx_messageInfo_WaveletHaar.DiscardUnknown(m)
}

var xxx_messageInfo_WaveletHaar proto.InternalMessageInfo

func (m *WaveletHaar) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type WaveletDaubechies struct {
	Level                uint32   `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	Coefficient          uint32   `protobuf:"varint,2,opt,name=coefficient,proto3" json:"coefficient,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WaveletDaubechies) Reset()         { *m = WaveletDaubechies{} }
func (m *WaveletDaubechies) String() string { return proto.CompactTextString(m) }
func (*WaveletDaubechies) ProtoMessage()    {}
func (*WaveletDaubechies) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{4}
}

func (m *WaveletDaubechies) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WaveletDaubechies.Unmarshal(m, b)
}
func (m *WaveletDaubechies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WaveletDaubechies.Marshal(b, m, deterministic)
}
func (m *WaveletDaubechies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaveletDaubechies.Merge(m, src)
}
func (m *WaveletDaubechies) XXX_Size() int {
	return xxx_messageInfo_WaveletDaubechies.Size(m)
}
func (m *WaveletDaubechies) XXX_DiscardUnknown() {
	xxx_messageInfo_WaveletDaubechies.DiscardUnknown(m)
}

var xxx_messageInfo_WaveletDaubechies proto.InternalMessageInfo

func (m *WaveletDaubechies) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *WaveletDaubechies) GetCoefficient() uint32 {
	if m != nil {
		return m.Coefficient
	}
	return 0
}

type WaveletConfig struct {
	// Types that are valid to be assigned to Data:
	//	*WaveletConfig_Haar
	//	*WaveletConfig_Daubechies
	Data                 isWaveletConfig_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *WaveletConfig) Reset()         { *m = WaveletConfig{} }
func (m *WaveletConfig) String() string { return proto.CompactTextString(m) }
func (*WaveletConfig) ProtoMessage()    {}
func (*WaveletConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{5}
}

func (m *WaveletConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WaveletConfig.Unmarshal(m, b)
}
func (m *WaveletConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WaveletConfig.Marshal(b, m, deterministic)
}
func (m *WaveletConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaveletConfig.Merge(m, src)
}
func (m *WaveletConfig) XXX_Size() int {
	return xxx_messageInfo_WaveletConfig.Size(m)
}
func (m *WaveletConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_WaveletConfig.DiscardUnknown(m)
}

var xxx_messageInfo_WaveletConfig proto.InternalMessageInfo

type isWaveletConfig_Data interface {
	isWaveletConfig_Data()
}

type WaveletConfig_Haar struct {
	Haar *WaveletHaar `protobuf:"bytes,1,opt,name=haar,proto3,oneof"`
}

type WaveletConfig_Daubechies struct {
	Daubechies *WaveletDaubechies `protobuf:"bytes,2,opt,name=daubechies,proto3,oneof"`
}

func (*WaveletConfig_Haar) isWaveletConfig_Data() {}

func (*WaveletConfig_Daubechies) isWaveletConfig_Data() {}

func (m *WaveletConfig) GetData() isWaveletConfig_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *WaveletConfig) GetHaar() *WaveletHaar {
	if x, ok := m.GetData().(*WaveletConfig_Haar); ok {
		return x.Haar
	}
	return nil
}

func (m *WaveletConfig) GetDaubechies() *WaveletDaubechies {
	if x, ok := m.GetData().(*WaveletConfig_Daubechies); ok {
		return x.Daubechies
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*WaveletConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*WaveletConfig_Haar)(nil),
		(*WaveletConfig_Daubechies)(nil),
	}
}

type QuantifierDeadZone struct {
	Width                uint32   `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Delta                uint32   `protobuf:"varint,2,opt,name=delta,proto3" json:"delta,omitempty"`
	Offset               float32  `protobuf:"fixed32,3,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuantifierDeadZone) Reset()         { *m = QuantifierDeadZone{} }
func (m *QuantifierDeadZone) String() string { return proto.CompactTextString(m) }
func (*QuantifierDeadZone) ProtoMessage()    {}
func (*QuantifierDeadZone) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{6}
}

func (m *QuantifierDeadZone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuantifierDeadZone.Unmarshal(m, b)
}
func (m *QuantifierDeadZone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuantifierDeadZone.Marshal(b, m, deterministic)
}
func (m *QuantifierDeadZone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuantifierDeadZone.Merge(m, src)
}
func (m *QuantifierDeadZone) XXX_Size() int {
	return xxx_messageInfo_QuantifierDeadZone.Size(m)
}
func (m *QuantifierDeadZone) XXX_DiscardUnknown() {
	xxx_messageInfo_QuantifierDeadZone.DiscardUnknown(m)
}

var xxx_messageInfo_QuantifierDeadZone proto.InternalMessageInfo

func (m *QuantifierDeadZone) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *QuantifierDeadZone) GetDelta() uint32 {
	if m != nil {
		return m.Delta
	}
	return 0
}

func (m *QuantifierDeadZone) GetOffset() float32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type QuantifierConfig struct {
	// Types that are valid to be assigned to Data:
	//	*QuantifierConfig_DeadZone
	Data                 isQuantifierConfig_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *QuantifierConfig) Reset()         { *m = QuantifierConfig{} }
func (m *QuantifierConfig) String() string { return proto.CompactTextString(m) }
func (*QuantifierConfig) ProtoMessage()    {}
func (*QuantifierConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{7}
}

func (m *QuantifierConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuantifierConfig.Unmarshal(m, b)
}
func (m *QuantifierConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuantifierConfig.Marshal(b, m, deterministic)
}
func (m *QuantifierConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuantifierConfig.Merge(m, src)
}
func (m *QuantifierConfig) XXX_Size() int {
	return xxx_messageInfo_QuantifierConfig.Size(m)
}
func (m *QuantifierConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_QuantifierConfig.DiscardUnknown(m)
}

var xxx_messageInfo_QuantifierConfig proto.InternalMessageInfo

type isQuantifierConfig_Data interface {
	isQuantifierConfig_Data()
}

type QuantifierConfig_DeadZone struct {
	DeadZone *QuantifierDeadZone `protobuf:"bytes,2,opt,name=dead_zone,json=deadZone,proto3,oneof"`
}

func (*QuantifierConfig_DeadZone) isQuantifierConfig_Data() {}

func (m *QuantifierConfig) GetData() isQuantifierConfig_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *QuantifierConfig) GetDeadZone() *QuantifierDeadZone {
	if x, ok := m.GetData().(*QuantifierConfig_DeadZone); ok {
		return x.DeadZone
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*QuantifierConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*QuantifierConfig_DeadZone)(nil),
	}
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
	return fileDescriptor_871986018790d2fd, []int{8}
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
	Width                uint32            `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32            `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Subsampling          Subsampling       `protobuf:"varint,3,opt,name=subsampling,proto3,enum=data.Subsampling" json:"subsampling,omitempty"`
	Wavelet              *WaveletConfig    `protobuf:"bytes,4,opt,name=wavelet,proto3" json:"wavelet,omitempty"`
	Quantifier           *QuantifierConfig `protobuf:"bytes,5,opt,name=quantifier,proto3" json:"quantifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FileImageHeader) Reset()         { *m = FileImageHeader{} }
func (m *FileImageHeader) String() string { return proto.CompactTextString(m) }
func (*FileImageHeader) ProtoMessage()    {}
func (*FileImageHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{9}
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

func (m *FileImageHeader) GetSubsampling() Subsampling {
	if m != nil {
		return m.Subsampling
	}
	return Subsampling_SUBSAMPLING_410
}

func (m *FileImageHeader) GetWavelet() *WaveletConfig {
	if m != nil {
		return m.Wavelet
	}
	return nil
}

func (m *FileImageHeader) GetQuantifier() *QuantifierConfig {
	if m != nil {
		return m.Quantifier
	}
	return nil
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
	return fileDescriptor_871986018790d2fd, []int{10}
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
	return fileDescriptor_871986018790d2fd, []int{11}
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
	proto.RegisterEnum("data.Subsampling", Subsampling_name, Subsampling_value)
	proto.RegisterType((*ImageRow)(nil), "data.ImageRow")
	proto.RegisterType((*ImageData)(nil), "data.ImageData")
	proto.RegisterType((*PythonDWT)(nil), "data.PythonDWT")
	proto.RegisterType((*WaveletHaar)(nil), "data.WaveletHaar")
	proto.RegisterType((*WaveletDaubechies)(nil), "data.WaveletDaubechies")
	proto.RegisterType((*WaveletConfig)(nil), "data.WaveletConfig")
	proto.RegisterType((*QuantifierDeadZone)(nil), "data.QuantifierDeadZone")
	proto.RegisterType((*QuantifierConfig)(nil), "data.QuantifierConfig")
	proto.RegisterType((*FileImageLayer)(nil), "data.FileImageLayer")
	proto.RegisterType((*FileImageHeader)(nil), "data.FileImageHeader")
	proto.RegisterType((*FileImageData)(nil), "data.FileImageData")
	proto.RegisterType((*FileImage)(nil), "data.FileImage")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor_871986018790d2fd) }

var fileDescriptor_871986018790d2fd = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x9b, 0x34, 0x34, 0x63, 0xf2, 0xd1, 0x4d, 0x09, 0x3e, 0x46, 0x2e, 0x12, 0x11, 0x52,
	0x4b, 0x49, 0x23, 0x10, 0x47, 0x4a, 0x04, 0xa9, 0x28, 0xa8, 0x6c, 0x40, 0x45, 0x5c, 0xaa, 0x4d,
	0x3c, 0x8e, 0x2d, 0x39, 0xde, 0xe2, 0xd8, 0x8e, 0x02, 0x3f, 0x95, 0x3f, 0x83, 0xf6, 0xc3, 0x8e,
	0x93, 0x90, 0xdb, 0xce, 0xf8, 0xed, 0xcc, 0x7b, 0x6f, 0x66, 0x0d, 0xe0, 0xb0, 0x98, 0x9d, 0x3f,
	0x44, 0x3c, 0xe6, 0xa4, 0x22, 0xce, 0xb6, 0x0d, 0x47, 0xd7, 0x73, 0x36, 0x43, 0xca, 0x97, 0xa4,
	0x03, 0xd5, 0x94, 0x05, 0x09, 0x2e, 0x2c, 0xa3, 0x5b, 0xee, 0x1d, 0x50, 0x1d, 0xd9, 0x2f, 0xa1,
	0x26, 0x31, 0x43, 0x16, 0x33, 0x62, 0x43, 0x25, 0xe2, 0x4b, 0x05, 0x31, 0xfb, 0x8d, 0x73, 0x59,
	0x31, 0x2b, 0x41, 0xe5, 0x37, 0x7b, 0x08, 0xb5, 0xdb, 0x55, 0xec, 0xf1, 0x70, 0x78, 0xf7, 0x8d,
	0x10, 0xa8, 0xcc, 0xb9, 0x83, 0x96, 0xd1, 0x35, 0x7a, 0x35, 0x2a, 0xcf, 0xe4, 0x14, 0x64, 0x77,
	0xeb, 0xa0, 0x6b, 0xf4, 0xcc, 0x7e, 0xb3, 0x50, 0x44, 0xf4, 0xa0, 0x8a, 0xda, 0x29, 0x98, 0x77,
	0x2c, 0xc5, 0x00, 0xe3, 0x11, 0x63, 0x11, 0x39, 0x81, 0xc3, 0x00, 0x53, 0x0c, 0x64, 0xa1, 0x3a,
	0x55, 0x81, 0xfd, 0x09, 0x8e, 0x35, 0x68, 0xc8, 0x92, 0x09, 0x4e, 0x3d, 0x1f, 0x17, 0xff, 0x87,
	0x92, 0x2e, 0x98, 0x53, 0x8e, 0xae, 0xeb, 0x4f, 0x7d, 0x0c, 0x63, 0xd9, 0xbb, 0x4e, 0x8b, 0x29,
	0xfb, 0x0f, 0xd4, 0x75, 0xb1, 0xf7, 0x3c, 0x74, 0xfd, 0x19, 0x79, 0x0e, 0x15, 0x8f, 0xb1, 0x48,
	0xd6, 0x31, 0xfb, 0xc7, 0x8a, 0x67, 0x81, 0xd4, 0xa8, 0x44, 0x25, 0x80, 0xbc, 0x15, 0xd6, 0x66,
	0xfd, 0xb5, 0xac, 0xa7, 0x1b, 0xf0, 0x35, 0xbd, 0x51, 0x89, 0x16, 0xc0, 0x57, 0x55, 0xe5, 0x85,
	0xfd, 0x03, 0xc8, 0xd7, 0x84, 0x85, 0xb1, 0xef, 0xfa, 0x18, 0x0d, 0x91, 0x39, 0x3f, 0x79, 0x88,
	0x42, 0xca, 0xd2, 0x77, 0x62, 0x2f, 0x93, 0x22, 0x03, 0x91, 0x75, 0x30, 0xd0, 0x06, 0xd6, 0xa9,
	0x0a, 0xc4, 0xfc, 0xb8, 0xeb, 0x2e, 0x30, 0xb6, 0xca, 0x5d, 0x43, 0xcc, 0x4f, 0x45, 0xf6, 0x18,
	0x5a, 0xeb, 0xca, 0x5a, 0xd9, 0x1b, 0xa8, 0x39, 0xc8, 0x9c, 0xfb, 0xdf, 0x3c, 0x44, 0xcd, 0xd7,
	0x52, 0x7c, 0x77, 0x49, 0x8c, 0x4a, 0xf4, 0xc8, 0xd1, 0xe7, 0x9c, 0xee, 0x33, 0x68, 0x7c, 0xf0,
	0x03, 0x94, 0x43, 0xbb, 0x61, 0x2b, 0x8c, 0xc4, 0xa0, 0xf3, 0xcd, 0x78, 0xac, 0x37, 0xe1, 0xaf,
	0x01, 0xcd, 0x1c, 0x36, 0x42, 0xe6, 0x60, 0xb4, 0x47, 0x52, 0x07, 0xaa, 0x1e, 0xfa, 0x33, 0x2f,
	0x1b, 0x8c, 0x8e, 0xc8, 0x25, 0x98, 0x8b, 0x64, 0xb2, 0x60, 0xf3, 0x87, 0xc0, 0x0f, 0x67, 0x52,
	0x59, 0x23, 0x9b, 0xc4, 0x78, 0xfd, 0x81, 0x16, 0x51, 0xe4, 0x0c, 0x1e, 0x2d, 0x95, 0xed, 0x56,
	0x45, 0x6a, 0x6b, 0x6f, 0xcc, 0x42, 0x79, 0x40, 0x33, 0x0c, 0x79, 0x0d, 0xf0, 0x2b, 0x57, 0x6d,
	0x1d, 0xca, 0x1b, 0x9d, 0x6d, 0x37, 0xf4, 0xa5, 0x02, 0x52, 0xec, 0x4b, 0x2e, 0x4e, 0x3f, 0x0e,
	0x63, 0xa5, 0x97, 0xe5, 0x44, 0xdd, 0xdf, 0xf4, 0x88, 0x1a, 0x2b, 0x81, 0x49, 0xb4, 0xe3, 0x7b,
	0x30, 0x89, 0xc0, 0xa4, 0x52, 0xea, 0x5e, 0x4c, 0x6a, 0x4f, 0xa1, 0x96, 0x27, 0xc9, 0x99, 0x70,
	0x4f, 0xb8, 0xab, 0xbb, 0x3f, 0xd9, 0xba, 0xa5, 0xac, 0xa7, 0x1a, 0x24, 0xf6, 0xba, 0xf0, 0xfe,
	0xda, 0x5b, 0xe0, 0xf5, 0x1b, 0x7c, 0xc1, 0xc0, 0x2c, 0x98, 0x4c, 0xda, 0xd0, 0x1c, 0x7f, 0xbf,
	0x1a, 0xbf, 0xfb, 0x7c, 0x7b, 0x73, 0xfd, 0xe5, 0xe3, 0xfd, 0xe0, 0xd5, 0x45, 0xab, 0xb4, 0x93,
	0xec, 0x5f, 0xb4, 0x8c, 0xdd, 0x64, 0xbf, 0x75, 0xb0, 0x93, 0x1c, 0x0c, 0x5a, 0xe5, 0x49, 0x55,
	0xfe, 0x8e, 0x2e, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x1f, 0x07, 0x1c, 0x9c, 0x04, 0x00,
	0x00,
}
