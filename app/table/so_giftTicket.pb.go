// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_giftTicket.proto

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

type SoGiftTicket struct {
	Ticket               *prototype.GiftTicketKeyType `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
	Denom                uint64                       `protobuf:"varint,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Count                uint64                       `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	ExpireBlock          uint64                       `protobuf:"varint,4,opt,name=expire_block,json=expireBlock,proto3" json:"expire_block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SoGiftTicket) Reset()         { *m = SoGiftTicket{} }
func (m *SoGiftTicket) String() string { return proto.CompactTextString(m) }
func (*SoGiftTicket) ProtoMessage()    {}
func (*SoGiftTicket) Descriptor() ([]byte, []int) {
	return fileDescriptor_d65be634f7a5df25, []int{0}
}

func (m *SoGiftTicket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoGiftTicket.Unmarshal(m, b)
}
func (m *SoGiftTicket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoGiftTicket.Marshal(b, m, deterministic)
}
func (m *SoGiftTicket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoGiftTicket.Merge(m, src)
}
func (m *SoGiftTicket) XXX_Size() int {
	return xxx_messageInfo_SoGiftTicket.Size(m)
}
func (m *SoGiftTicket) XXX_DiscardUnknown() {
	xxx_messageInfo_SoGiftTicket.DiscardUnknown(m)
}

var xxx_messageInfo_SoGiftTicket proto.InternalMessageInfo

func (m *SoGiftTicket) GetTicket() *prototype.GiftTicketKeyType {
	if m != nil {
		return m.Ticket
	}
	return nil
}

func (m *SoGiftTicket) GetDenom() uint64 {
	if m != nil {
		return m.Denom
	}
	return 0
}

func (m *SoGiftTicket) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *SoGiftTicket) GetExpireBlock() uint64 {
	if m != nil {
		return m.ExpireBlock
	}
	return 0
}

type SoListGiftTicketByTicket struct {
	Ticket               *prototype.GiftTicketKeyType `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SoListGiftTicketByTicket) Reset()         { *m = SoListGiftTicketByTicket{} }
func (m *SoListGiftTicketByTicket) String() string { return proto.CompactTextString(m) }
func (*SoListGiftTicketByTicket) ProtoMessage()    {}
func (*SoListGiftTicketByTicket) Descriptor() ([]byte, []int) {
	return fileDescriptor_d65be634f7a5df25, []int{1}
}

func (m *SoListGiftTicketByTicket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListGiftTicketByTicket.Unmarshal(m, b)
}
func (m *SoListGiftTicketByTicket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListGiftTicketByTicket.Marshal(b, m, deterministic)
}
func (m *SoListGiftTicketByTicket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListGiftTicketByTicket.Merge(m, src)
}
func (m *SoListGiftTicketByTicket) XXX_Size() int {
	return xxx_messageInfo_SoListGiftTicketByTicket.Size(m)
}
func (m *SoListGiftTicketByTicket) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListGiftTicketByTicket.DiscardUnknown(m)
}

var xxx_messageInfo_SoListGiftTicketByTicket proto.InternalMessageInfo

func (m *SoListGiftTicketByTicket) GetTicket() *prototype.GiftTicketKeyType {
	if m != nil {
		return m.Ticket
	}
	return nil
}

type SoListGiftTicketByCount struct {
	Count                uint64                       `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Ticket               *prototype.GiftTicketKeyType `protobuf:"bytes,2,opt,name=ticket,proto3" json:"ticket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SoListGiftTicketByCount) Reset()         { *m = SoListGiftTicketByCount{} }
func (m *SoListGiftTicketByCount) String() string { return proto.CompactTextString(m) }
func (*SoListGiftTicketByCount) ProtoMessage()    {}
func (*SoListGiftTicketByCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_d65be634f7a5df25, []int{2}
}

func (m *SoListGiftTicketByCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListGiftTicketByCount.Unmarshal(m, b)
}
func (m *SoListGiftTicketByCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListGiftTicketByCount.Marshal(b, m, deterministic)
}
func (m *SoListGiftTicketByCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListGiftTicketByCount.Merge(m, src)
}
func (m *SoListGiftTicketByCount) XXX_Size() int {
	return xxx_messageInfo_SoListGiftTicketByCount.Size(m)
}
func (m *SoListGiftTicketByCount) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListGiftTicketByCount.DiscardUnknown(m)
}

var xxx_messageInfo_SoListGiftTicketByCount proto.InternalMessageInfo

func (m *SoListGiftTicketByCount) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *SoListGiftTicketByCount) GetTicket() *prototype.GiftTicketKeyType {
	if m != nil {
		return m.Ticket
	}
	return nil
}

type SoUniqueGiftTicketByTicket struct {
	Ticket               *prototype.GiftTicketKeyType `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SoUniqueGiftTicketByTicket) Reset()         { *m = SoUniqueGiftTicketByTicket{} }
func (m *SoUniqueGiftTicketByTicket) String() string { return proto.CompactTextString(m) }
func (*SoUniqueGiftTicketByTicket) ProtoMessage()    {}
func (*SoUniqueGiftTicketByTicket) Descriptor() ([]byte, []int) {
	return fileDescriptor_d65be634f7a5df25, []int{3}
}

func (m *SoUniqueGiftTicketByTicket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueGiftTicketByTicket.Unmarshal(m, b)
}
func (m *SoUniqueGiftTicketByTicket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueGiftTicketByTicket.Marshal(b, m, deterministic)
}
func (m *SoUniqueGiftTicketByTicket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueGiftTicketByTicket.Merge(m, src)
}
func (m *SoUniqueGiftTicketByTicket) XXX_Size() int {
	return xxx_messageInfo_SoUniqueGiftTicketByTicket.Size(m)
}
func (m *SoUniqueGiftTicketByTicket) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueGiftTicketByTicket.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueGiftTicketByTicket proto.InternalMessageInfo

func (m *SoUniqueGiftTicketByTicket) GetTicket() *prototype.GiftTicketKeyType {
	if m != nil {
		return m.Ticket
	}
	return nil
}

func init() {
	proto.RegisterType((*SoGiftTicket)(nil), "table.so_giftTicket")
	proto.RegisterType((*SoListGiftTicketByTicket)(nil), "table.so_list_giftTicket_by_ticket")
	proto.RegisterType((*SoListGiftTicketByCount)(nil), "table.so_list_giftTicket_by_count")
	proto.RegisterType((*SoUniqueGiftTicketByTicket)(nil), "table.so_unique_giftTicket_by_ticket")
}

func init() { proto.RegisterFile("app/table/so_giftTicket.proto", fileDescriptor_d65be634f7a5df25) }

var fileDescriptor_d65be634f7a5df25 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x49, 0x6d, 0x7b, 0x48, 0xf5, 0xb2, 0xf4, 0xb0, 0xf8, 0xb7, 0xf6, 0x54, 0x44, 0x37,
	0xa0, 0x07, 0xef, 0x7d, 0x84, 0x22, 0x88, 0x5e, 0xc2, 0x26, 0x8e, 0xdb, 0xb0, 0xdb, 0x4c, 0x6c,
	0x66, 0xc1, 0x7d, 0x12, 0x5f, 0x57, 0x92, 0x94, 0xb5, 0x82, 0x17, 0xc5, 0x4b, 0xc8, 0x7c, 0xdf,
	0xcc, 0x7c, 0xbf, 0x40, 0xf8, 0x59, 0xe9, 0x9c, 0xa0, 0x52, 0x35, 0x20, 0x3c, 0xca, 0xca, 0xbc,
	0xd2, 0x83, 0xd1, 0x35, 0x50, 0xe1, 0xb6, 0x48, 0x98, 0x8d, 0xa2, 0x75, 0x3c, 0x8d, 0x15, 0x75,
	0x0e, 0x44, 0x38, 0x92, 0x39, 0xff, 0x60, 0xfc, 0xe8, 0xdb, 0x50, 0x76, 0xcf, 0xc7, 0x14, 0x6f,
	0x39, 0x9b, 0xb1, 0xc5, 0xe4, 0xf6, 0xa2, 0xe8, 0x07, 0x8b, 0xd0, 0x26, 0x93, 0x2b, 0x6b, 0xe8,
	0x64, 0x10, 0x57, 0xbb, 0xf6, 0x6c, 0xca, 0x47, 0x2f, 0x60, 0x71, 0x93, 0x0f, 0x66, 0x6c, 0x31,
	0x5c, 0xa5, 0x22, 0xa8, 0x1a, 0x5b, 0x4b, 0xf9, 0x41, 0x52, 0x63, 0x91, 0x5d, 0xf2, 0x43, 0x78,
	0x77, 0x66, 0x0b, 0x52, 0x35, 0xa8, 0xeb, 0x7c, 0x18, 0xcd, 0x49, 0xd2, 0x96, 0x41, 0x9a, 0x3f,
	0xf2, 0x53, 0x8f, 0xb2, 0x31, 0x9e, 0xf6, 0xe8, 0xa4, 0xea, 0x76, 0xf9, 0x7f, 0xe6, 0x9c, 0x37,
	0xfc, 0xe4, 0xe7, 0xc5, 0x09, 0xad, 0x07, 0x66, 0xfb, 0xc0, 0x5f, 0x69, 0x83, 0xdf, 0xa5, 0x3d,
	0xf1, 0x73, 0x8f, 0xb2, 0xb5, 0xe6, 0xad, 0x85, 0xff, 0x7d, 0xc8, 0xf2, 0xfa, 0xf9, 0xaa, 0x32,
	0xb4, 0x6e, 0x55, 0xa1, 0x71, 0x23, 0x34, 0x7a, 0xbd, 0x2e, 0x8d, 0x15, 0x1a, 0x2d, 0x81, 0x25,
	0xf4, 0x37, 0x15, 0x8a, 0xfe, 0x6b, 0xa8, 0x71, 0xdc, 0x7a, 0xf7, 0x19, 0x00, 0x00, 0xff, 0xff,
	0xc1, 0x87, 0x8b, 0x27, 0x2e, 0x02, 0x00, 0x00,
}
