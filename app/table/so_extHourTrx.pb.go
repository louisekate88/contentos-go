// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_extHourTrx.proto

package table

import (
	fmt "fmt"
	prototype "github.com/coschain/contentos-go/prototype"
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

type SoExtHourTrx struct {
	Hour                 *prototype.TimePointSec `protobuf:"bytes,1,opt,name=hour,proto3" json:"hour,omitempty"`
	Count                uint64                  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoExtHourTrx) Reset()         { *m = SoExtHourTrx{} }
func (m *SoExtHourTrx) String() string { return proto.CompactTextString(m) }
func (*SoExtHourTrx) ProtoMessage()    {}
func (*SoExtHourTrx) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccee57d6353b6f3d, []int{0}
}

func (m *SoExtHourTrx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoExtHourTrx.Unmarshal(m, b)
}
func (m *SoExtHourTrx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoExtHourTrx.Marshal(b, m, deterministic)
}
func (m *SoExtHourTrx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoExtHourTrx.Merge(m, src)
}
func (m *SoExtHourTrx) XXX_Size() int {
	return xxx_messageInfo_SoExtHourTrx.Size(m)
}
func (m *SoExtHourTrx) XXX_DiscardUnknown() {
	xxx_messageInfo_SoExtHourTrx.DiscardUnknown(m)
}

var xxx_messageInfo_SoExtHourTrx proto.InternalMessageInfo

func (m *SoExtHourTrx) GetHour() *prototype.TimePointSec {
	if m != nil {
		return m.Hour
	}
	return nil
}

func (m *SoExtHourTrx) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type SoListExtHourTrxByHour struct {
	Hour                 *prototype.TimePointSec `protobuf:"bytes,1,opt,name=hour,proto3" json:"hour,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoListExtHourTrxByHour) Reset()         { *m = SoListExtHourTrxByHour{} }
func (m *SoListExtHourTrxByHour) String() string { return proto.CompactTextString(m) }
func (*SoListExtHourTrxByHour) ProtoMessage()    {}
func (*SoListExtHourTrxByHour) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccee57d6353b6f3d, []int{1}
}

func (m *SoListExtHourTrxByHour) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtHourTrxByHour.Unmarshal(m, b)
}
func (m *SoListExtHourTrxByHour) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtHourTrxByHour.Marshal(b, m, deterministic)
}
func (m *SoListExtHourTrxByHour) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtHourTrxByHour.Merge(m, src)
}
func (m *SoListExtHourTrxByHour) XXX_Size() int {
	return xxx_messageInfo_SoListExtHourTrxByHour.Size(m)
}
func (m *SoListExtHourTrxByHour) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtHourTrxByHour.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtHourTrxByHour proto.InternalMessageInfo

func (m *SoListExtHourTrxByHour) GetHour() *prototype.TimePointSec {
	if m != nil {
		return m.Hour
	}
	return nil
}

type SoListExtHourTrxByCount struct {
	Count                uint64                  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Hour                 *prototype.TimePointSec `protobuf:"bytes,2,opt,name=hour,proto3" json:"hour,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoListExtHourTrxByCount) Reset()         { *m = SoListExtHourTrxByCount{} }
func (m *SoListExtHourTrxByCount) String() string { return proto.CompactTextString(m) }
func (*SoListExtHourTrxByCount) ProtoMessage()    {}
func (*SoListExtHourTrxByCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccee57d6353b6f3d, []int{2}
}

func (m *SoListExtHourTrxByCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtHourTrxByCount.Unmarshal(m, b)
}
func (m *SoListExtHourTrxByCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtHourTrxByCount.Marshal(b, m, deterministic)
}
func (m *SoListExtHourTrxByCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtHourTrxByCount.Merge(m, src)
}
func (m *SoListExtHourTrxByCount) XXX_Size() int {
	return xxx_messageInfo_SoListExtHourTrxByCount.Size(m)
}
func (m *SoListExtHourTrxByCount) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtHourTrxByCount.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtHourTrxByCount proto.InternalMessageInfo

func (m *SoListExtHourTrxByCount) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *SoListExtHourTrxByCount) GetHour() *prototype.TimePointSec {
	if m != nil {
		return m.Hour
	}
	return nil
}

type SoUniqueExtHourTrxByHour struct {
	Hour                 *prototype.TimePointSec `protobuf:"bytes,1,opt,name=hour,proto3" json:"hour,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoUniqueExtHourTrxByHour) Reset()         { *m = SoUniqueExtHourTrxByHour{} }
func (m *SoUniqueExtHourTrxByHour) String() string { return proto.CompactTextString(m) }
func (*SoUniqueExtHourTrxByHour) ProtoMessage()    {}
func (*SoUniqueExtHourTrxByHour) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccee57d6353b6f3d, []int{3}
}

func (m *SoUniqueExtHourTrxByHour) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueExtHourTrxByHour.Unmarshal(m, b)
}
func (m *SoUniqueExtHourTrxByHour) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueExtHourTrxByHour.Marshal(b, m, deterministic)
}
func (m *SoUniqueExtHourTrxByHour) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueExtHourTrxByHour.Merge(m, src)
}
func (m *SoUniqueExtHourTrxByHour) XXX_Size() int {
	return xxx_messageInfo_SoUniqueExtHourTrxByHour.Size(m)
}
func (m *SoUniqueExtHourTrxByHour) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueExtHourTrxByHour.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueExtHourTrxByHour proto.InternalMessageInfo

func (m *SoUniqueExtHourTrxByHour) GetHour() *prototype.TimePointSec {
	if m != nil {
		return m.Hour
	}
	return nil
}

func init() {
	proto.RegisterType((*SoExtHourTrx)(nil), "table.so_extHourTrx")
	proto.RegisterType((*SoListExtHourTrxByHour)(nil), "table.so_list_extHourTrx_by_hour")
	proto.RegisterType((*SoListExtHourTrxByCount)(nil), "table.so_list_extHourTrx_by_count")
	proto.RegisterType((*SoUniqueExtHourTrxByHour)(nil), "table.so_unique_extHourTrx_by_hour")
}

func init() { proto.RegisterFile("app/table/so_extHourTrx.proto", fileDescriptor_ccee57d6353b6f3d) }

var fileDescriptor_ccee57d6353b6f3d = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0x41, 0x4b, 0xc4, 0x30,
	0x14, 0x84, 0xc9, 0xb2, 0xeb, 0x21, 0xe2, 0xa5, 0xec, 0x61, 0x5d, 0x15, 0x96, 0x9e, 0x8a, 0xd8,
	0x04, 0xf4, 0x1f, 0x78, 0x12, 0xc4, 0x4b, 0xe9, 0xc9, 0x4b, 0x68, 0x42, 0x68, 0x03, 0x6d, 0x5e,
	0x4c, 0x5e, 0xa0, 0xfd, 0xf7, 0x62, 0x0a, 0xb5, 0x82, 0x07, 0x65, 0x2f, 0x81, 0x61, 0x32, 0xf3,
	0x0d, 0x3c, 0x7a, 0xd7, 0x38, 0xc7, 0xb1, 0x91, 0xbd, 0xe6, 0x01, 0x84, 0x1e, 0xf1, 0x05, 0xa2,
	0xaf, 0xfd, 0xc8, 0x9c, 0x07, 0x84, 0x6c, 0x97, 0xac, 0xe3, 0x3e, 0x29, 0x9c, 0x9c, 0xe6, 0x5f,
	0xcf, 0x6c, 0xe6, 0x35, 0xbd, 0xfa, 0x91, 0xc9, 0x4a, 0xba, 0xed, 0x20, 0xfa, 0x03, 0x39, 0x91,
	0xe2, 0xf2, 0xf1, 0x9a, 0x2d, 0x29, 0x86, 0x66, 0xd0, 0xc2, 0x81, 0xb1, 0x28, 0x82, 0x56, 0x55,
	0xfa, 0x96, 0xed, 0xe9, 0x4e, 0x41, 0xb4, 0x78, 0xd8, 0x9c, 0x48, 0xb1, 0xad, 0x66, 0x91, 0xbf,
	0xd2, 0x63, 0x00, 0xd1, 0x9b, 0x80, 0xab, 0x6a, 0x21, 0x27, 0x91, 0x32, 0xff, 0x43, 0xe4, 0x92,
	0xde, 0xfc, 0x5e, 0x96, 0x58, 0xdf, 0x0b, 0xc8, 0x6a, 0xc1, 0xc2, 0xd8, 0xfc, 0x8d, 0xf1, 0x46,
	0x6f, 0x03, 0x88, 0x68, 0xcd, 0x47, 0xd4, 0xe7, 0x4f, 0x7e, 0x7e, 0x78, 0xbf, 0x6f, 0x0d, 0x76,
	0x51, 0x32, 0x05, 0x03, 0x57, 0x10, 0x54, 0xd7, 0x18, 0xcb, 0x15, 0x58, 0xd4, 0x16, 0x21, 0x94,
	0x2d, 0xf0, 0xe5, 0x68, 0xf2, 0x22, 0xb5, 0x3d, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff, 0x91, 0x0b,
	0x58, 0xb1, 0xc8, 0x01, 0x00, 0x00,
}
