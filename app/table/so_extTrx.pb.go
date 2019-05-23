// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_extTrx.proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SoExtTrx struct {
	TrxId                *prototype.Sha256             `protobuf:"bytes,1,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	BlockHeight          uint64                        `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	TrxWrap              *prototype.TransactionWrapper `protobuf:"bytes,3,opt,name=trx_wrap,json=trxWrap,proto3" json:"trx_wrap,omitempty"`
	BlockTime            *prototype.TimePointSec       `protobuf:"bytes,4,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`
	TrxCreateOrder       *prototype.UserTrxCreateOrder `protobuf:"bytes,5,opt,name=trx_create_order,json=trxCreateOrder,proto3" json:"trx_create_order,omitempty"`
	BlockId              *prototype.Sha256             `protobuf:"bytes,6,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SoExtTrx) Reset()         { *m = SoExtTrx{} }
func (m *SoExtTrx) String() string { return proto.CompactTextString(m) }
func (*SoExtTrx) ProtoMessage()    {}
func (*SoExtTrx) Descriptor() ([]byte, []int) {
	return fileDescriptor_76957eaae1f8a1bc, []int{0}
}

func (m *SoExtTrx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoExtTrx.Unmarshal(m, b)
}
func (m *SoExtTrx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoExtTrx.Marshal(b, m, deterministic)
}
func (m *SoExtTrx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoExtTrx.Merge(m, src)
}
func (m *SoExtTrx) XXX_Size() int {
	return xxx_messageInfo_SoExtTrx.Size(m)
}
func (m *SoExtTrx) XXX_DiscardUnknown() {
	xxx_messageInfo_SoExtTrx.DiscardUnknown(m)
}

var xxx_messageInfo_SoExtTrx proto.InternalMessageInfo

func (m *SoExtTrx) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

func (m *SoExtTrx) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *SoExtTrx) GetTrxWrap() *prototype.TransactionWrapper {
	if m != nil {
		return m.TrxWrap
	}
	return nil
}

func (m *SoExtTrx) GetBlockTime() *prototype.TimePointSec {
	if m != nil {
		return m.BlockTime
	}
	return nil
}

func (m *SoExtTrx) GetTrxCreateOrder() *prototype.UserTrxCreateOrder {
	if m != nil {
		return m.TrxCreateOrder
	}
	return nil
}

func (m *SoExtTrx) GetBlockId() *prototype.Sha256 {
	if m != nil {
		return m.BlockId
	}
	return nil
}

type SoListExtTrxByTrxId struct {
	TrxId                *prototype.Sha256 `protobuf:"bytes,1,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SoListExtTrxByTrxId) Reset()         { *m = SoListExtTrxByTrxId{} }
func (m *SoListExtTrxByTrxId) String() string { return proto.CompactTextString(m) }
func (*SoListExtTrxByTrxId) ProtoMessage()    {}
func (*SoListExtTrxByTrxId) Descriptor() ([]byte, []int) {
	return fileDescriptor_76957eaae1f8a1bc, []int{1}
}

func (m *SoListExtTrxByTrxId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtTrxByTrxId.Unmarshal(m, b)
}
func (m *SoListExtTrxByTrxId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtTrxByTrxId.Marshal(b, m, deterministic)
}
func (m *SoListExtTrxByTrxId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtTrxByTrxId.Merge(m, src)
}
func (m *SoListExtTrxByTrxId) XXX_Size() int {
	return xxx_messageInfo_SoListExtTrxByTrxId.Size(m)
}
func (m *SoListExtTrxByTrxId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtTrxByTrxId.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtTrxByTrxId proto.InternalMessageInfo

func (m *SoListExtTrxByTrxId) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

type SoListExtTrxByBlockHeight struct {
	BlockHeight          uint64            `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	TrxId                *prototype.Sha256 `protobuf:"bytes,2,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SoListExtTrxByBlockHeight) Reset()         { *m = SoListExtTrxByBlockHeight{} }
func (m *SoListExtTrxByBlockHeight) String() string { return proto.CompactTextString(m) }
func (*SoListExtTrxByBlockHeight) ProtoMessage()    {}
func (*SoListExtTrxByBlockHeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_76957eaae1f8a1bc, []int{2}
}

func (m *SoListExtTrxByBlockHeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtTrxByBlockHeight.Unmarshal(m, b)
}
func (m *SoListExtTrxByBlockHeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtTrxByBlockHeight.Marshal(b, m, deterministic)
}
func (m *SoListExtTrxByBlockHeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtTrxByBlockHeight.Merge(m, src)
}
func (m *SoListExtTrxByBlockHeight) XXX_Size() int {
	return xxx_messageInfo_SoListExtTrxByBlockHeight.Size(m)
}
func (m *SoListExtTrxByBlockHeight) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtTrxByBlockHeight.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtTrxByBlockHeight proto.InternalMessageInfo

func (m *SoListExtTrxByBlockHeight) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *SoListExtTrxByBlockHeight) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

type SoListExtTrxByBlockTime struct {
	BlockTime            *prototype.TimePointSec `protobuf:"bytes,1,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`
	TrxId                *prototype.Sha256       `protobuf:"bytes,2,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoListExtTrxByBlockTime) Reset()         { *m = SoListExtTrxByBlockTime{} }
func (m *SoListExtTrxByBlockTime) String() string { return proto.CompactTextString(m) }
func (*SoListExtTrxByBlockTime) ProtoMessage()    {}
func (*SoListExtTrxByBlockTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_76957eaae1f8a1bc, []int{3}
}

func (m *SoListExtTrxByBlockTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtTrxByBlockTime.Unmarshal(m, b)
}
func (m *SoListExtTrxByBlockTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtTrxByBlockTime.Marshal(b, m, deterministic)
}
func (m *SoListExtTrxByBlockTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtTrxByBlockTime.Merge(m, src)
}
func (m *SoListExtTrxByBlockTime) XXX_Size() int {
	return xxx_messageInfo_SoListExtTrxByBlockTime.Size(m)
}
func (m *SoListExtTrxByBlockTime) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtTrxByBlockTime.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtTrxByBlockTime proto.InternalMessageInfo

func (m *SoListExtTrxByBlockTime) GetBlockTime() *prototype.TimePointSec {
	if m != nil {
		return m.BlockTime
	}
	return nil
}

func (m *SoListExtTrxByBlockTime) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

type SoListExtTrxByTrxCreateOrder struct {
	TrxCreateOrder       *prototype.UserTrxCreateOrder `protobuf:"bytes,1,opt,name=trx_create_order,json=trxCreateOrder,proto3" json:"trx_create_order,omitempty"`
	TrxId                *prototype.Sha256             `protobuf:"bytes,2,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SoListExtTrxByTrxCreateOrder) Reset()         { *m = SoListExtTrxByTrxCreateOrder{} }
func (m *SoListExtTrxByTrxCreateOrder) String() string { return proto.CompactTextString(m) }
func (*SoListExtTrxByTrxCreateOrder) ProtoMessage()    {}
func (*SoListExtTrxByTrxCreateOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_76957eaae1f8a1bc, []int{4}
}

func (m *SoListExtTrxByTrxCreateOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtTrxByTrxCreateOrder.Unmarshal(m, b)
}
func (m *SoListExtTrxByTrxCreateOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtTrxByTrxCreateOrder.Marshal(b, m, deterministic)
}
func (m *SoListExtTrxByTrxCreateOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtTrxByTrxCreateOrder.Merge(m, src)
}
func (m *SoListExtTrxByTrxCreateOrder) XXX_Size() int {
	return xxx_messageInfo_SoListExtTrxByTrxCreateOrder.Size(m)
}
func (m *SoListExtTrxByTrxCreateOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtTrxByTrxCreateOrder.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtTrxByTrxCreateOrder proto.InternalMessageInfo

func (m *SoListExtTrxByTrxCreateOrder) GetTrxCreateOrder() *prototype.UserTrxCreateOrder {
	if m != nil {
		return m.TrxCreateOrder
	}
	return nil
}

func (m *SoListExtTrxByTrxCreateOrder) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

type SoUniqueExtTrxByTrxId struct {
	TrxId                *prototype.Sha256 `protobuf:"bytes,1,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SoUniqueExtTrxByTrxId) Reset()         { *m = SoUniqueExtTrxByTrxId{} }
func (m *SoUniqueExtTrxByTrxId) String() string { return proto.CompactTextString(m) }
func (*SoUniqueExtTrxByTrxId) ProtoMessage()    {}
func (*SoUniqueExtTrxByTrxId) Descriptor() ([]byte, []int) {
	return fileDescriptor_76957eaae1f8a1bc, []int{5}
}

func (m *SoUniqueExtTrxByTrxId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueExtTrxByTrxId.Unmarshal(m, b)
}
func (m *SoUniqueExtTrxByTrxId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueExtTrxByTrxId.Marshal(b, m, deterministic)
}
func (m *SoUniqueExtTrxByTrxId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueExtTrxByTrxId.Merge(m, src)
}
func (m *SoUniqueExtTrxByTrxId) XXX_Size() int {
	return xxx_messageInfo_SoUniqueExtTrxByTrxId.Size(m)
}
func (m *SoUniqueExtTrxByTrxId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueExtTrxByTrxId.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueExtTrxByTrxId proto.InternalMessageInfo

func (m *SoUniqueExtTrxByTrxId) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

func init() {
	proto.RegisterType((*SoExtTrx)(nil), "table.so_extTrx")
	proto.RegisterType((*SoListExtTrxByTrxId)(nil), "table.so_list_extTrx_by_trx_id")
	proto.RegisterType((*SoListExtTrxByBlockHeight)(nil), "table.so_list_extTrx_by_block_height")
	proto.RegisterType((*SoListExtTrxByBlockTime)(nil), "table.so_list_extTrx_by_block_time")
	proto.RegisterType((*SoListExtTrxByTrxCreateOrder)(nil), "table.so_list_extTrx_by_trx_create_order")
	proto.RegisterType((*SoUniqueExtTrxByTrxId)(nil), "table.so_unique_extTrx_by_trx_id")
}

func init() { proto.RegisterFile("app/table/so_extTrx.proto", fileDescriptor_76957eaae1f8a1bc) }

var fileDescriptor_76957eaae1f8a1bc = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x5d, 0x0b, 0xd3, 0x30,
	0x14, 0x25, 0x73, 0x9f, 0x99, 0x88, 0x16, 0x1f, 0xb2, 0x29, 0x63, 0xf6, 0x69, 0xc8, 0x6c, 0x61,
	0xa2, 0xe8, 0xab, 0x8a, 0x38, 0x5f, 0x84, 0x32, 0x10, 0x7c, 0x09, 0x69, 0x1a, 0xd6, 0x60, 0xdb,
	0xc4, 0xe4, 0x16, 0xbb, 0x57, 0x7f, 0x86, 0x7f, 0xc8, 0xbf, 0x25, 0x4d, 0xc7, 0x56, 0xf6, 0x81,
	0x53, 0x5f, 0x0a, 0x39, 0xf7, 0x9e, 0x73, 0x2e, 0xe7, 0x50, 0x3c, 0x61, 0x5a, 0x87, 0xc0, 0xe2,
	0x4c, 0x84, 0x56, 0x51, 0x51, 0xc1, 0xc6, 0x54, 0x81, 0x36, 0x0a, 0x94, 0xd7, 0x73, 0xf0, 0x94,
	0xb8, 0x17, 0xec, 0xb4, 0x08, 0xf3, 0x32, 0x03, 0x49, 0x65, 0xd2, 0x2c, 0x4c, 0x1f, 0x1e, 0x27,
	0xf5, 0x67, 0x8f, 0x3e, 0x6a, 0xa1, 0x86, 0x15, 0x96, 0x71, 0x90, 0xaa, 0x68, 0x86, 0xfe, 0xaf,
	0x0e, 0x1e, 0x1d, 0x7c, 0xbc, 0x05, 0xee, 0x83, 0xa9, 0xa8, 0x4c, 0x08, 0x9a, 0xa3, 0xc5, 0x78,
	0xf5, 0x20, 0x38, 0x70, 0x03, 0x9b, 0xb2, 0xd5, 0x8b, 0x97, 0x51, 0x0f, 0x4c, 0xb5, 0x4e, 0xbc,
	0x27, 0xf8, 0x6e, 0x9c, 0x29, 0xfe, 0x95, 0xa6, 0x42, 0x6e, 0x53, 0x20, 0x9d, 0x39, 0x5a, 0x74,
	0xa3, 0xb1, 0xc3, 0x3e, 0x38, 0xc8, 0x7b, 0x8d, 0x87, 0xb5, 0xd8, 0x77, 0xc3, 0x34, 0xb9, 0xe3,
	0xe4, 0x66, 0x2d, 0xb9, 0xd6, 0x29, 0x6e, 0x45, 0x0b, 0x13, 0x0d, 0xc0, 0x54, 0x9f, 0x0d, 0xd3,
	0xde, 0x2b, 0x8c, 0x1b, 0x75, 0x90, 0xb9, 0x20, 0x5d, 0x47, 0x9e, 0xb4, 0xc9, 0x32, 0x17, 0x54,
	0x2b, 0x59, 0x00, 0xb5, 0x82, 0x47, 0x23, 0xb7, 0xbc, 0x91, 0xb9, 0xf0, 0x3e, 0xe2, 0xfb, 0xb5,
	0x29, 0x37, 0x82, 0x81, 0xa0, 0xca, 0x24, 0xc2, 0x90, 0x9e, 0xe3, 0xcf, 0x5b, 0xfc, 0xd2, 0x0a,
	0x43, 0x4f, 0xf7, 0xa2, 0x7b, 0x60, 0xaa, 0xb7, 0x0e, 0xf8, 0x54, 0xbf, 0xbd, 0x25, 0x1e, 0x36,
	0x57, 0xc8, 0x84, 0xf4, 0xaf, 0xe5, 0x31, 0x70, 0x2b, 0xeb, 0xc4, 0x7f, 0x87, 0x89, 0x55, 0x34,
	0x93, 0x16, 0xf6, 0x69, 0xd2, 0x78, 0x47, 0x9b, 0x34, 0x6f, 0xcf, 0xd5, 0xcf, 0xf1, 0xec, 0x5c,
	0xa5, 0x9d, 0xf4, 0x59, 0xf2, 0xe8, 0x3c, 0xf9, 0xa3, 0x5d, 0xe7, 0x0f, 0x76, 0x3f, 0x10, 0x7e,
	0x7c, 0xcd, 0xaf, 0x0e, 0xf9, 0xa4, 0x09, 0xf4, 0x17, 0x4d, 0xdc, 0x7e, 0xc4, 0x4f, 0x84, 0xfd,
	0xcb, 0xd1, 0xb5, 0xeb, 0xb9, 0x58, 0x2d, 0xfa, 0xc7, 0x6a, 0x6f, 0x3f, 0xee, 0x3d, 0x9e, 0x5a,
	0x45, 0xcb, 0x42, 0x7e, 0x2b, 0xc5, 0x7f, 0x14, 0xfb, 0x66, 0xf9, 0xe5, 0xe9, 0x56, 0x42, 0x5a,
	0xc6, 0x01, 0x57, 0x79, 0xc8, 0x95, 0xe5, 0x29, 0x93, 0x45, 0xc8, 0x55, 0x01, 0xa2, 0x00, 0x65,
	0x9f, 0x6d, 0x55, 0x78, 0xf8, 0xf5, 0xe3, 0xbe, 0x93, 0x79, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0x5c, 0x86, 0x6c, 0x4d, 0x0e, 0x04, 0x00, 0x00,
}
