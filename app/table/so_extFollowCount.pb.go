// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_extFollowCount.proto

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

type SoExtFollowCount struct {
	Account              *prototype.AccountName  `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	FollowerCnt          uint32                  `protobuf:"varint,2,opt,name=follower_cnt,json=followerCnt,proto3" json:"follower_cnt,omitempty"`
	FollowingCnt         uint32                  `protobuf:"varint,3,opt,name=following_cnt,json=followingCnt,proto3" json:"following_cnt,omitempty"`
	UpdateTime           *prototype.TimePointSec `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoExtFollowCount) Reset()         { *m = SoExtFollowCount{} }
func (m *SoExtFollowCount) String() string { return proto.CompactTextString(m) }
func (*SoExtFollowCount) ProtoMessage()    {}
func (*SoExtFollowCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c3525aab04885f, []int{0}
}

func (m *SoExtFollowCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoExtFollowCount.Unmarshal(m, b)
}
func (m *SoExtFollowCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoExtFollowCount.Marshal(b, m, deterministic)
}
func (m *SoExtFollowCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoExtFollowCount.Merge(m, src)
}
func (m *SoExtFollowCount) XXX_Size() int {
	return xxx_messageInfo_SoExtFollowCount.Size(m)
}
func (m *SoExtFollowCount) XXX_DiscardUnknown() {
	xxx_messageInfo_SoExtFollowCount.DiscardUnknown(m)
}

var xxx_messageInfo_SoExtFollowCount proto.InternalMessageInfo

func (m *SoExtFollowCount) GetAccount() *prototype.AccountName {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *SoExtFollowCount) GetFollowerCnt() uint32 {
	if m != nil {
		return m.FollowerCnt
	}
	return 0
}

func (m *SoExtFollowCount) GetFollowingCnt() uint32 {
	if m != nil {
		return m.FollowingCnt
	}
	return 0
}

func (m *SoExtFollowCount) GetUpdateTime() *prototype.TimePointSec {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

type SoMemExtFollowCountByAccount struct {
	Account              *prototype.AccountName `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoMemExtFollowCountByAccount) Reset()         { *m = SoMemExtFollowCountByAccount{} }
func (m *SoMemExtFollowCountByAccount) String() string { return proto.CompactTextString(m) }
func (*SoMemExtFollowCountByAccount) ProtoMessage()    {}
func (*SoMemExtFollowCountByAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c3525aab04885f, []int{1}
}

func (m *SoMemExtFollowCountByAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtFollowCountByAccount.Unmarshal(m, b)
}
func (m *SoMemExtFollowCountByAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtFollowCountByAccount.Marshal(b, m, deterministic)
}
func (m *SoMemExtFollowCountByAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtFollowCountByAccount.Merge(m, src)
}
func (m *SoMemExtFollowCountByAccount) XXX_Size() int {
	return xxx_messageInfo_SoMemExtFollowCountByAccount.Size(m)
}
func (m *SoMemExtFollowCountByAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtFollowCountByAccount.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtFollowCountByAccount proto.InternalMessageInfo

func (m *SoMemExtFollowCountByAccount) GetAccount() *prototype.AccountName {
	if m != nil {
		return m.Account
	}
	return nil
}

type SoMemExtFollowCountByFollowerCnt struct {
	FollowerCnt          uint32   `protobuf:"varint,1,opt,name=follower_cnt,json=followerCnt,proto3" json:"follower_cnt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemExtFollowCountByFollowerCnt) Reset()         { *m = SoMemExtFollowCountByFollowerCnt{} }
func (m *SoMemExtFollowCountByFollowerCnt) String() string { return proto.CompactTextString(m) }
func (*SoMemExtFollowCountByFollowerCnt) ProtoMessage()    {}
func (*SoMemExtFollowCountByFollowerCnt) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c3525aab04885f, []int{2}
}

func (m *SoMemExtFollowCountByFollowerCnt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtFollowCountByFollowerCnt.Unmarshal(m, b)
}
func (m *SoMemExtFollowCountByFollowerCnt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtFollowCountByFollowerCnt.Marshal(b, m, deterministic)
}
func (m *SoMemExtFollowCountByFollowerCnt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtFollowCountByFollowerCnt.Merge(m, src)
}
func (m *SoMemExtFollowCountByFollowerCnt) XXX_Size() int {
	return xxx_messageInfo_SoMemExtFollowCountByFollowerCnt.Size(m)
}
func (m *SoMemExtFollowCountByFollowerCnt) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtFollowCountByFollowerCnt.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtFollowCountByFollowerCnt proto.InternalMessageInfo

func (m *SoMemExtFollowCountByFollowerCnt) GetFollowerCnt() uint32 {
	if m != nil {
		return m.FollowerCnt
	}
	return 0
}

type SoMemExtFollowCountByFollowingCnt struct {
	FollowingCnt         uint32   `protobuf:"varint,1,opt,name=following_cnt,json=followingCnt,proto3" json:"following_cnt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemExtFollowCountByFollowingCnt) Reset()         { *m = SoMemExtFollowCountByFollowingCnt{} }
func (m *SoMemExtFollowCountByFollowingCnt) String() string { return proto.CompactTextString(m) }
func (*SoMemExtFollowCountByFollowingCnt) ProtoMessage()    {}
func (*SoMemExtFollowCountByFollowingCnt) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c3525aab04885f, []int{3}
}

func (m *SoMemExtFollowCountByFollowingCnt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtFollowCountByFollowingCnt.Unmarshal(m, b)
}
func (m *SoMemExtFollowCountByFollowingCnt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtFollowCountByFollowingCnt.Marshal(b, m, deterministic)
}
func (m *SoMemExtFollowCountByFollowingCnt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtFollowCountByFollowingCnt.Merge(m, src)
}
func (m *SoMemExtFollowCountByFollowingCnt) XXX_Size() int {
	return xxx_messageInfo_SoMemExtFollowCountByFollowingCnt.Size(m)
}
func (m *SoMemExtFollowCountByFollowingCnt) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtFollowCountByFollowingCnt.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtFollowCountByFollowingCnt proto.InternalMessageInfo

func (m *SoMemExtFollowCountByFollowingCnt) GetFollowingCnt() uint32 {
	if m != nil {
		return m.FollowingCnt
	}
	return 0
}

type SoMemExtFollowCountByUpdateTime struct {
	UpdateTime           *prototype.TimePointSec `protobuf:"bytes,1,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoMemExtFollowCountByUpdateTime) Reset()         { *m = SoMemExtFollowCountByUpdateTime{} }
func (m *SoMemExtFollowCountByUpdateTime) String() string { return proto.CompactTextString(m) }
func (*SoMemExtFollowCountByUpdateTime) ProtoMessage()    {}
func (*SoMemExtFollowCountByUpdateTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c3525aab04885f, []int{4}
}

func (m *SoMemExtFollowCountByUpdateTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtFollowCountByUpdateTime.Unmarshal(m, b)
}
func (m *SoMemExtFollowCountByUpdateTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtFollowCountByUpdateTime.Marshal(b, m, deterministic)
}
func (m *SoMemExtFollowCountByUpdateTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtFollowCountByUpdateTime.Merge(m, src)
}
func (m *SoMemExtFollowCountByUpdateTime) XXX_Size() int {
	return xxx_messageInfo_SoMemExtFollowCountByUpdateTime.Size(m)
}
func (m *SoMemExtFollowCountByUpdateTime) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtFollowCountByUpdateTime.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtFollowCountByUpdateTime proto.InternalMessageInfo

func (m *SoMemExtFollowCountByUpdateTime) GetUpdateTime() *prototype.TimePointSec {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

type SoUniqueExtFollowCountByAccount struct {
	Account              *prototype.AccountName `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoUniqueExtFollowCountByAccount) Reset()         { *m = SoUniqueExtFollowCountByAccount{} }
func (m *SoUniqueExtFollowCountByAccount) String() string { return proto.CompactTextString(m) }
func (*SoUniqueExtFollowCountByAccount) ProtoMessage()    {}
func (*SoUniqueExtFollowCountByAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c3525aab04885f, []int{5}
}

func (m *SoUniqueExtFollowCountByAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueExtFollowCountByAccount.Unmarshal(m, b)
}
func (m *SoUniqueExtFollowCountByAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueExtFollowCountByAccount.Marshal(b, m, deterministic)
}
func (m *SoUniqueExtFollowCountByAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueExtFollowCountByAccount.Merge(m, src)
}
func (m *SoUniqueExtFollowCountByAccount) XXX_Size() int {
	return xxx_messageInfo_SoUniqueExtFollowCountByAccount.Size(m)
}
func (m *SoUniqueExtFollowCountByAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueExtFollowCountByAccount.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueExtFollowCountByAccount proto.InternalMessageInfo

func (m *SoUniqueExtFollowCountByAccount) GetAccount() *prototype.AccountName {
	if m != nil {
		return m.Account
	}
	return nil
}

func init() {
	proto.RegisterType((*SoExtFollowCount)(nil), "table.so_extFollowCount")
	proto.RegisterType((*SoMemExtFollowCountByAccount)(nil), "table.so_mem_extFollowCount_by_account")
	proto.RegisterType((*SoMemExtFollowCountByFollowerCnt)(nil), "table.so_mem_extFollowCount_by_follower_cnt")
	proto.RegisterType((*SoMemExtFollowCountByFollowingCnt)(nil), "table.so_mem_extFollowCount_by_following_cnt")
	proto.RegisterType((*SoMemExtFollowCountByUpdateTime)(nil), "table.so_mem_extFollowCount_by_update_time")
	proto.RegisterType((*SoUniqueExtFollowCountByAccount)(nil), "table.so_unique_extFollowCount_by_account")
}

func init() { proto.RegisterFile("app/table/so_extFollowCount.proto", fileDescriptor_30c3525aab04885f) }

var fileDescriptor_30c3525aab04885f = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x4b, 0x33, 0x31,
	0x10, 0xc6, 0xc9, 0xfb, 0xfa, 0x07, 0xa6, 0xed, 0xc1, 0x45, 0xb0, 0x7a, 0x6a, 0xb7, 0x5a, 0x7a,
	0x71, 0x17, 0xf5, 0xe6, 0xd1, 0x82, 0x07, 0xc1, 0x4b, 0x51, 0x10, 0x2f, 0x61, 0x37, 0x8e, 0x6d,
	0xa0, 0xc9, 0xc4, 0x66, 0x16, 0xed, 0x57, 0xf4, 0x53, 0xc9, 0x66, 0xdb, 0x52, 0xba, 0x14, 0x29,
	0x78, 0x09, 0x99, 0x99, 0x67, 0x9e, 0xc9, 0xfc, 0x08, 0x74, 0x33, 0xe7, 0x52, 0xce, 0xf2, 0x29,
	0xa6, 0x9e, 0x24, 0x7e, 0xf1, 0x3d, 0x4d, 0xa7, 0xf4, 0x39, 0xa4, 0xc2, 0x72, 0xe2, 0x66, 0xc4,
	0x14, 0xed, 0x87, 0xf2, 0xd9, 0x71, 0x88, 0x78, 0xee, 0x30, 0x2d, 0x8f, 0xaa, 0x18, 0x7f, 0x0b,
	0x38, 0xaa, 0x35, 0x46, 0x57, 0x70, 0x98, 0x29, 0x55, 0x5e, 0xdb, 0xa2, 0x23, 0x06, 0x8d, 0xeb,
	0x93, 0x64, 0xd5, 0x9d, 0x2c, 0x2a, 0xd2, 0x66, 0x06, 0x47, 0x4b, 0x5d, 0xd4, 0x85, 0xe6, 0x7b,
	0x70, 0xc0, 0x99, 0x54, 0x96, 0xdb, 0xff, 0x3a, 0x62, 0xd0, 0x1a, 0x35, 0x96, 0xb9, 0xa1, 0xe5,
	0xa8, 0x07, 0xad, 0x2a, 0xd4, 0x76, 0x1c, 0x34, 0xff, 0x83, 0xa6, 0xb9, 0x4a, 0x96, 0xa2, 0x5b,
	0x68, 0x14, 0xee, 0x2d, 0x63, 0x94, 0xac, 0x0d, 0xb6, 0xf7, 0xc2, 0xf8, 0xd3, 0xb5, 0xf1, 0x65,
	0x5a, 0x3a, 0xd2, 0x96, 0xa5, 0x47, 0x35, 0x82, 0x4a, 0xfd, 0xa4, 0x0d, 0xc6, 0xcf, 0xd0, 0xf1,
	0x24, 0x0d, 0x9a, 0x8d, 0x7d, 0x64, 0x3e, 0x97, 0xcb, 0x77, 0xee, 0xbe, 0x5a, 0xfc, 0x00, 0x17,
	0x5b, 0x6d, 0xd7, 0x77, 0xae, 0x31, 0x10, 0x35, 0x06, 0xf1, 0x23, 0xf4, 0x7f, 0xf1, 0x5a, 0xc0,
	0xa9, 0xd3, 0x12, 0x75, 0x5a, 0x71, 0x0e, 0xe7, 0x5b, 0xed, 0xd6, 0x30, 0x6e, 0x52, 0x15, 0xbb,
	0x50, 0x7d, 0x81, 0x9e, 0x27, 0x59, 0x58, 0xfd, 0x51, 0xe0, 0x9f, 0x82, 0xbd, 0x1b, 0xbc, 0xf6,
	0xc7, 0x9a, 0x27, 0x45, 0x9e, 0x28, 0x32, 0xa9, 0x22, 0xaf, 0x26, 0x99, 0xb6, 0xa9, 0x22, 0xcb,
	0x68, 0x99, 0xfc, 0xe5, 0x98, 0xaa, 0xbf, 0x9d, 0x1f, 0x04, 0xab, 0x9b, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x98, 0x31, 0xfa, 0x97, 0xef, 0x02, 0x00, 0x00,
}
