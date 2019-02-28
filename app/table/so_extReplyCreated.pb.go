// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_extReplyCreated.proto

package table // import "github.com/coschain/contentos-go/table"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import prototype "github.com/coschain/contentos-go/prototype"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SoExtReplyCreated struct {
	PostId               uint64                       `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	CreatedOrder         *prototype.ReplyCreatedOrder `protobuf:"bytes,2,opt,name=created_order,json=createdOrder,proto3" json:"created_order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SoExtReplyCreated) Reset()         { *m = SoExtReplyCreated{} }
func (m *SoExtReplyCreated) String() string { return proto.CompactTextString(m) }
func (*SoExtReplyCreated) ProtoMessage()    {}
func (*SoExtReplyCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_extReplyCreated_dfe38825a295c52c, []int{0}
}
func (m *SoExtReplyCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoExtReplyCreated.Unmarshal(m, b)
}
func (m *SoExtReplyCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoExtReplyCreated.Marshal(b, m, deterministic)
}
func (dst *SoExtReplyCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoExtReplyCreated.Merge(dst, src)
}
func (m *SoExtReplyCreated) XXX_Size() int {
	return xxx_messageInfo_SoExtReplyCreated.Size(m)
}
func (m *SoExtReplyCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_SoExtReplyCreated.DiscardUnknown(m)
}

var xxx_messageInfo_SoExtReplyCreated proto.InternalMessageInfo

func (m *SoExtReplyCreated) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func (m *SoExtReplyCreated) GetCreatedOrder() *prototype.ReplyCreatedOrder {
	if m != nil {
		return m.CreatedOrder
	}
	return nil
}

type SoMemExtReplyCreatedByPostId struct {
	PostId               uint64   `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemExtReplyCreatedByPostId) Reset()         { *m = SoMemExtReplyCreatedByPostId{} }
func (m *SoMemExtReplyCreatedByPostId) String() string { return proto.CompactTextString(m) }
func (*SoMemExtReplyCreatedByPostId) ProtoMessage()    {}
func (*SoMemExtReplyCreatedByPostId) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_extReplyCreated_dfe38825a295c52c, []int{1}
}
func (m *SoMemExtReplyCreatedByPostId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtReplyCreatedByPostId.Unmarshal(m, b)
}
func (m *SoMemExtReplyCreatedByPostId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtReplyCreatedByPostId.Marshal(b, m, deterministic)
}
func (dst *SoMemExtReplyCreatedByPostId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtReplyCreatedByPostId.Merge(dst, src)
}
func (m *SoMemExtReplyCreatedByPostId) XXX_Size() int {
	return xxx_messageInfo_SoMemExtReplyCreatedByPostId.Size(m)
}
func (m *SoMemExtReplyCreatedByPostId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtReplyCreatedByPostId.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtReplyCreatedByPostId proto.InternalMessageInfo

func (m *SoMemExtReplyCreatedByPostId) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

type SoMemExtReplyCreatedByCreatedOrder struct {
	CreatedOrder         *prototype.ReplyCreatedOrder `protobuf:"bytes,1,opt,name=created_order,json=createdOrder,proto3" json:"created_order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SoMemExtReplyCreatedByCreatedOrder) Reset()         { *m = SoMemExtReplyCreatedByCreatedOrder{} }
func (m *SoMemExtReplyCreatedByCreatedOrder) String() string { return proto.CompactTextString(m) }
func (*SoMemExtReplyCreatedByCreatedOrder) ProtoMessage()    {}
func (*SoMemExtReplyCreatedByCreatedOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_extReplyCreated_dfe38825a295c52c, []int{2}
}
func (m *SoMemExtReplyCreatedByCreatedOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtReplyCreatedByCreatedOrder.Unmarshal(m, b)
}
func (m *SoMemExtReplyCreatedByCreatedOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtReplyCreatedByCreatedOrder.Marshal(b, m, deterministic)
}
func (dst *SoMemExtReplyCreatedByCreatedOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtReplyCreatedByCreatedOrder.Merge(dst, src)
}
func (m *SoMemExtReplyCreatedByCreatedOrder) XXX_Size() int {
	return xxx_messageInfo_SoMemExtReplyCreatedByCreatedOrder.Size(m)
}
func (m *SoMemExtReplyCreatedByCreatedOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtReplyCreatedByCreatedOrder.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtReplyCreatedByCreatedOrder proto.InternalMessageInfo

func (m *SoMemExtReplyCreatedByCreatedOrder) GetCreatedOrder() *prototype.ReplyCreatedOrder {
	if m != nil {
		return m.CreatedOrder
	}
	return nil
}

type SoListExtReplyCreatedByCreatedOrder struct {
	CreatedOrder         *prototype.ReplyCreatedOrder `protobuf:"bytes,1,opt,name=created_order,json=createdOrder,proto3" json:"created_order,omitempty"`
	PostId               uint64                       `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SoListExtReplyCreatedByCreatedOrder) Reset()         { *m = SoListExtReplyCreatedByCreatedOrder{} }
func (m *SoListExtReplyCreatedByCreatedOrder) String() string { return proto.CompactTextString(m) }
func (*SoListExtReplyCreatedByCreatedOrder) ProtoMessage()    {}
func (*SoListExtReplyCreatedByCreatedOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_extReplyCreated_dfe38825a295c52c, []int{3}
}
func (m *SoListExtReplyCreatedByCreatedOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtReplyCreatedByCreatedOrder.Unmarshal(m, b)
}
func (m *SoListExtReplyCreatedByCreatedOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtReplyCreatedByCreatedOrder.Marshal(b, m, deterministic)
}
func (dst *SoListExtReplyCreatedByCreatedOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtReplyCreatedByCreatedOrder.Merge(dst, src)
}
func (m *SoListExtReplyCreatedByCreatedOrder) XXX_Size() int {
	return xxx_messageInfo_SoListExtReplyCreatedByCreatedOrder.Size(m)
}
func (m *SoListExtReplyCreatedByCreatedOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtReplyCreatedByCreatedOrder.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtReplyCreatedByCreatedOrder proto.InternalMessageInfo

func (m *SoListExtReplyCreatedByCreatedOrder) GetCreatedOrder() *prototype.ReplyCreatedOrder {
	if m != nil {
		return m.CreatedOrder
	}
	return nil
}

func (m *SoListExtReplyCreatedByCreatedOrder) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

type SoUniqueExtReplyCreatedByPostId struct {
	PostId               uint64   `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoUniqueExtReplyCreatedByPostId) Reset()         { *m = SoUniqueExtReplyCreatedByPostId{} }
func (m *SoUniqueExtReplyCreatedByPostId) String() string { return proto.CompactTextString(m) }
func (*SoUniqueExtReplyCreatedByPostId) ProtoMessage()    {}
func (*SoUniqueExtReplyCreatedByPostId) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_extReplyCreated_dfe38825a295c52c, []int{4}
}
func (m *SoUniqueExtReplyCreatedByPostId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueExtReplyCreatedByPostId.Unmarshal(m, b)
}
func (m *SoUniqueExtReplyCreatedByPostId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueExtReplyCreatedByPostId.Marshal(b, m, deterministic)
}
func (dst *SoUniqueExtReplyCreatedByPostId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueExtReplyCreatedByPostId.Merge(dst, src)
}
func (m *SoUniqueExtReplyCreatedByPostId) XXX_Size() int {
	return xxx_messageInfo_SoUniqueExtReplyCreatedByPostId.Size(m)
}
func (m *SoUniqueExtReplyCreatedByPostId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueExtReplyCreatedByPostId.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueExtReplyCreatedByPostId proto.InternalMessageInfo

func (m *SoUniqueExtReplyCreatedByPostId) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func init() {
	proto.RegisterType((*SoExtReplyCreated)(nil), "table.so_extReplyCreated")
	proto.RegisterType((*SoMemExtReplyCreatedByPostId)(nil), "table.so_mem_extReplyCreated_by_post_id")
	proto.RegisterType((*SoMemExtReplyCreatedByCreatedOrder)(nil), "table.so_mem_extReplyCreated_by_created_order")
	proto.RegisterType((*SoListExtReplyCreatedByCreatedOrder)(nil), "table.so_list_extReplyCreated_by_created_order")
	proto.RegisterType((*SoUniqueExtReplyCreatedByPostId)(nil), "table.so_unique_extReplyCreated_by_post_id")
}

func init() {
	proto.RegisterFile("app/table/so_extReplyCreated.proto", fileDescriptor_so_extReplyCreated_dfe38825a295c52c)
}

var fileDescriptor_so_extReplyCreated_dfe38825a295c52c = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0xc9, 0xa1, 0x27, 0x44, 0x5d, 0xba, 0x58, 0x1c, 0xe4, 0x2c, 0xa2, 0x5d, 0x6c, 0x40,
	0x57, 0x41, 0xf0, 0x26, 0x27, 0xa1, 0xa3, 0xcb, 0xa3, 0x4d, 0x1f, 0x77, 0x81, 0x36, 0x2f, 0x26,
	0xaf, 0x60, 0xbf, 0x81, 0x1f, 0x5b, 0xda, 0x2b, 0x62, 0x3d, 0x6e, 0xe9, 0xe0, 0xf8, 0x27, 0xff,
	0x97, 0xdf, 0xef, 0x85, 0xc8, 0xa4, 0x70, 0x4e, 0x71, 0x51, 0xd6, 0xa8, 0x02, 0x01, 0x7e, 0x72,
	0x8e, 0xae, 0xee, 0xd6, 0x1e, 0x0b, 0xc6, 0x2a, 0x73, 0x9e, 0x98, 0xa2, 0xe3, 0xe1, 0xfc, 0x32,
	0x1e, 0x12, 0x77, 0x0e, 0x55, 0xd3, 0xd6, 0x6c, 0xc0, 0x8c, 0x85, 0xc4, 0xcb, 0x68, 0x7f, 0x38,
	0xba, 0x90, 0x27, 0x8e, 0x02, 0x83, 0xa9, 0x62, 0xb1, 0x12, 0xe9, 0x51, 0xbe, 0xec, 0xe3, 0x6b,
	0x15, 0xad, 0xe5, 0xb9, 0xde, 0x75, 0x80, 0x7c, 0x85, 0x3e, 0x5e, 0xac, 0x44, 0x7a, 0xfa, 0x70,
	0x95, 0xfd, 0x00, 0x32, 0xdf, 0x5f, 0x04, 0x93, 0x56, 0x7e, 0x36, 0xc6, 0xb7, 0x3e, 0x25, 0x4f,
	0xf2, 0x3a, 0x10, 0x34, 0xd8, 0xfc, 0xe5, 0x42, 0xd9, 0xc1, 0xc8, 0x3d, 0xa8, 0x90, 0x58, 0x79,
	0x77, 0x78, 0x7a, 0x82, 0xdd, 0xb7, 0x15, 0x33, 0x6c, 0xbf, 0x84, 0x4c, 0x03, 0x41, 0x6d, 0x02,
	0xff, 0x0f, 0xf1, 0xf7, 0xea, 0x8b, 0xc9, 0xea, 0xcf, 0xf2, 0x26, 0x10, 0xb4, 0xd6, 0x7c, 0xb4,
	0x38, 0xe7, 0xed, 0x5e, 0xd2, 0xf7, 0xdb, 0x8d, 0xe1, 0x6d, 0x5b, 0x66, 0x9a, 0x1a, 0xa5, 0x29,
	0xe8, 0x6d, 0x61, 0xac, 0xd2, 0x64, 0x19, 0x2d, 0x53, 0xb8, 0xdf, 0xd0, 0xee, 0x47, 0x95, 0xcb,
	0x41, 0xf8, 0xf1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xde, 0x9f, 0xd1, 0x4d, 0x65, 0x02, 0x00, 0x00,
}
