// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_witness.proto

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

type SoWitness struct {
	Owner                 *prototype.AccountName   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	CreatedTime           *prototype.TimePointSec  `protobuf:"bytes,2,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	Url                   string                   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	VoteCount             uint64                   `protobuf:"varint,4,opt,name=vote_count,json=voteCount,proto3" json:"vote_count,omitempty"`
	LastConfirmedBlockNum uint32                   `protobuf:"varint,5,opt,name=last_confirmed_block_num,json=lastConfirmedBlockNum,proto3" json:"last_confirmed_block_num,omitempty"`
	TotalMissed           uint32                   `protobuf:"varint,6,opt,name=total_missed,json=totalMissed,proto3" json:"total_missed,omitempty"`
	PowWorker             uint32                   `protobuf:"varint,7,opt,name=pow_worker,json=powWorker,proto3" json:"pow_worker,omitempty"`
	SigningKey            *prototype.PublicKeyType `protobuf:"bytes,8,opt,name=signing_key,json=signingKey,proto3" json:"signing_key,omitempty"`
	LastWork              *prototype.Sha256        `protobuf:"bytes,9,opt,name=last_work,json=lastWork,proto3" json:"last_work,omitempty"`
	RunningVersion        uint32                   `protobuf:"varint,10,opt,name=running_version,json=runningVersion,proto3" json:"running_version,omitempty"`
	LastAslot             uint32                   `protobuf:"varint,11,opt,name=last_aslot,json=lastAslot,proto3" json:"last_aslot,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                 `json:"-"`
	XXX_unrecognized      []byte                   `json:"-"`
	XXX_sizecache         int32                    `json:"-"`
}

func (m *SoWitness) Reset()         { *m = SoWitness{} }
func (m *SoWitness) String() string { return proto.CompactTextString(m) }
func (*SoWitness) ProtoMessage()    {}
func (*SoWitness) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_witness_3e7e301b243138ab, []int{0}
}
func (m *SoWitness) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoWitness.Unmarshal(m, b)
}
func (m *SoWitness) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoWitness.Marshal(b, m, deterministic)
}
func (dst *SoWitness) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoWitness.Merge(dst, src)
}
func (m *SoWitness) XXX_Size() int {
	return xxx_messageInfo_SoWitness.Size(m)
}
func (m *SoWitness) XXX_DiscardUnknown() {
	xxx_messageInfo_SoWitness.DiscardUnknown(m)
}

var xxx_messageInfo_SoWitness proto.InternalMessageInfo

func (m *SoWitness) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *SoWitness) GetCreatedTime() *prototype.TimePointSec {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

func (m *SoWitness) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SoWitness) GetVoteCount() uint64 {
	if m != nil {
		return m.VoteCount
	}
	return 0
}

func (m *SoWitness) GetLastConfirmedBlockNum() uint32 {
	if m != nil {
		return m.LastConfirmedBlockNum
	}
	return 0
}

func (m *SoWitness) GetTotalMissed() uint32 {
	if m != nil {
		return m.TotalMissed
	}
	return 0
}

func (m *SoWitness) GetPowWorker() uint32 {
	if m != nil {
		return m.PowWorker
	}
	return 0
}

func (m *SoWitness) GetSigningKey() *prototype.PublicKeyType {
	if m != nil {
		return m.SigningKey
	}
	return nil
}

func (m *SoWitness) GetLastWork() *prototype.Sha256 {
	if m != nil {
		return m.LastWork
	}
	return nil
}

func (m *SoWitness) GetRunningVersion() uint32 {
	if m != nil {
		return m.RunningVersion
	}
	return 0
}

func (m *SoWitness) GetLastAslot() uint32 {
	if m != nil {
		return m.LastAslot
	}
	return 0
}

type SoMemWitnessByOwner struct {
	Owner                *prototype.AccountName `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoMemWitnessByOwner) Reset()         { *m = SoMemWitnessByOwner{} }
func (m *SoMemWitnessByOwner) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByOwner) ProtoMessage()    {}
func (*SoMemWitnessByOwner) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_witness_3e7e301b243138ab, []int{1}
}
func (m *SoMemWitnessByOwner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByOwner.Unmarshal(m, b)
}
func (m *SoMemWitnessByOwner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByOwner.Marshal(b, m, deterministic)
}
func (dst *SoMemWitnessByOwner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByOwner.Merge(dst, src)
}
func (m *SoMemWitnessByOwner) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByOwner.Size(m)
}
func (m *SoMemWitnessByOwner) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByOwner.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByOwner proto.InternalMessageInfo

func (m *SoMemWitnessByOwner) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

type SoMemWitnessByCreatedTime struct {
	CreatedTime          *prototype.TimePointSec `protobuf:"bytes,1,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoMemWitnessByCreatedTime) Reset()         { *m = SoMemWitnessByCreatedTime{} }
func (m *SoMemWitnessByCreatedTime) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByCreatedTime) ProtoMessage()    {}
func (*SoMemWitnessByCreatedTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_witness_3e7e301b243138ab, []int{2}
}
func (m *SoMemWitnessByCreatedTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByCreatedTime.Unmarshal(m, b)
}
func (m *SoMemWitnessByCreatedTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByCreatedTime.Marshal(b, m, deterministic)
}
func (dst *SoMemWitnessByCreatedTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByCreatedTime.Merge(dst, src)
}
func (m *SoMemWitnessByCreatedTime) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByCreatedTime.Size(m)
}
func (m *SoMemWitnessByCreatedTime) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByCreatedTime.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByCreatedTime proto.InternalMessageInfo

func (m *SoMemWitnessByCreatedTime) GetCreatedTime() *prototype.TimePointSec {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

type SoMemWitnessByUrl struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemWitnessByUrl) Reset()         { *m = SoMemWitnessByUrl{} }
func (m *SoMemWitnessByUrl) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByUrl) ProtoMessage()    {}
func (*SoMemWitnessByUrl) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_witness_3e7e301b243138ab, []int{3}
}
func (m *SoMemWitnessByUrl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByUrl.Unmarshal(m, b)
}
func (m *SoMemWitnessByUrl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByUrl.Marshal(b, m, deterministic)
}
func (dst *SoMemWitnessByUrl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByUrl.Merge(dst, src)
}
func (m *SoMemWitnessByUrl) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByUrl.Size(m)
}
func (m *SoMemWitnessByUrl) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByUrl.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByUrl proto.InternalMessageInfo

func (m *SoMemWitnessByUrl) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type SoMemWitnessByVoteCount struct {
	VoteCount            uint64   `protobuf:"varint,1,opt,name=vote_count,json=voteCount,proto3" json:"vote_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemWitnessByVoteCount) Reset()         { *m = SoMemWitnessByVoteCount{} }
func (m *SoMemWitnessByVoteCount) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByVoteCount) ProtoMessage()    {}
func (*SoMemWitnessByVoteCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_witness_3e7e301b243138ab, []int{4}
}
func (m *SoMemWitnessByVoteCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByVoteCount.Unmarshal(m, b)
}
func (m *SoMemWitnessByVoteCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByVoteCount.Marshal(b, m, deterministic)
}
func (dst *SoMemWitnessByVoteCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByVoteCount.Merge(dst, src)
}
func (m *SoMemWitnessByVoteCount) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByVoteCount.Size(m)
}
func (m *SoMemWitnessByVoteCount) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByVoteCount.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByVoteCount proto.InternalMessageInfo

func (m *SoMemWitnessByVoteCount) GetVoteCount() uint64 {
	if m != nil {
		return m.VoteCount
	}
	return 0
}

type SoMemWitnessByLastConfirmedBlockNum struct {
	LastConfirmedBlockNum uint32   `protobuf:"varint,1,opt,name=last_confirmed_block_num,json=lastConfirmedBlockNum,proto3" json:"last_confirmed_block_num,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *SoMemWitnessByLastConfirmedBlockNum) Reset()         { *m = SoMemWitnessByLastConfirmedBlockNum{} }
func (m *SoMemWitnessByLastConfirmedBlockNum) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByLastConfirmedBlockNum) ProtoMessage()    {}
func (*SoMemWitnessByLastConfirmedBlockNum) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_witness_3e7e301b243138ab, []int{5}
}
func (m *SoMemWitnessByLastConfirmedBlockNum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByLastConfirmedBlockNum.Unmarshal(m, b)
}
func (m *SoMemWitnessByLastConfirmedBlockNum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByLastConfirmedBlockNum.Marshal(b, m, deterministic)
}
func (dst *SoMemWitnessByLastConfirmedBlockNum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByLastConfirmedBlockNum.Merge(dst, src)
}
func (m *SoMemWitnessByLastConfirmedBlockNum) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByLastConfirmedBlockNum.Size(m)
}
func (m *SoMemWitnessByLastConfirmedBlockNum) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByLastConfirmedBlockNum.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByLastConfirmedBlockNum proto.InternalMessageInfo

func (m *SoMemWitnessByLastConfirmedBlockNum) GetLastConfirmedBlockNum() uint32 {
	if m != nil {
		return m.LastConfirmedBlockNum
	}
	return 0
}

type SoMemWitnessByTotalMissed struct {
	TotalMissed          uint32   `protobuf:"varint,1,opt,name=total_missed,json=totalMissed,proto3" json:"total_missed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemWitnessByTotalMissed) Reset()         { *m = SoMemWitnessByTotalMissed{} }
func (m *SoMemWitnessByTotalMissed) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByTotalMissed) ProtoMessage()    {}
func (*SoMemWitnessByTotalMissed) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_witness_3e7e301b243138ab, []int{6}
}
func (m *SoMemWitnessByTotalMissed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByTotalMissed.Unmarshal(m, b)
}
func (m *SoMemWitnessByTotalMissed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByTotalMissed.Marshal(b, m, deterministic)
}
func (dst *SoMemWitnessByTotalMissed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByTotalMissed.Merge(dst, src)
}
func (m *SoMemWitnessByTotalMissed) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByTotalMissed.Size(m)
}
func (m *SoMemWitnessByTotalMissed) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByTotalMissed.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByTotalMissed proto.InternalMessageInfo

func (m *SoMemWitnessByTotalMissed) GetTotalMissed() uint32 {
	if m != nil {
		return m.TotalMissed
	}
	return 0
}

type SoMemWitnessByPowWorker struct {
	PowWorker            uint32   `protobuf:"varint,1,opt,name=pow_worker,json=powWorker,proto3" json:"pow_worker,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemWitnessByPowWorker) Reset()         { *m = SoMemWitnessByPowWorker{} }
func (m *SoMemWitnessByPowWorker) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByPowWorker) ProtoMessage()    {}
func (*SoMemWitnessByPowWorker) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_witness_3e7e301b243138ab, []int{7}
}
func (m *SoMemWitnessByPowWorker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByPowWorker.Unmarshal(m, b)
}
func (m *SoMemWitnessByPowWorker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByPowWorker.Marshal(b, m, deterministic)
}
func (dst *SoMemWitnessByPowWorker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByPowWorker.Merge(dst, src)
}
func (m *SoMemWitnessByPowWorker) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByPowWorker.Size(m)
}
func (m *SoMemWitnessByPowWorker) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByPowWorker.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByPowWorker proto.InternalMessageInfo

func (m *SoMemWitnessByPowWorker) GetPowWorker() uint32 {
	if m != nil {
		return m.PowWorker
	}
	return 0
}

type SoMemWitnessBySigningKey struct {
	SigningKey           *prototype.PublicKeyType `protobuf:"bytes,1,opt,name=signing_key,json=signingKey,proto3" json:"signing_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SoMemWitnessBySigningKey) Reset()         { *m = SoMemWitnessBySigningKey{} }
func (m *SoMemWitnessBySigningKey) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessBySigningKey) ProtoMessage()    {}
func (*SoMemWitnessBySigningKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_witness_3e7e301b243138ab, []int{8}
}
func (m *SoMemWitnessBySigningKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessBySigningKey.Unmarshal(m, b)
}
func (m *SoMemWitnessBySigningKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessBySigningKey.Marshal(b, m, deterministic)
}
func (dst *SoMemWitnessBySigningKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessBySigningKey.Merge(dst, src)
}
func (m *SoMemWitnessBySigningKey) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessBySigningKey.Size(m)
}
func (m *SoMemWitnessBySigningKey) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessBySigningKey.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessBySigningKey proto.InternalMessageInfo

func (m *SoMemWitnessBySigningKey) GetSigningKey() *prototype.PublicKeyType {
	if m != nil {
		return m.SigningKey
	}
	return nil
}

type SoMemWitnessByLastWork struct {
	LastWork             *prototype.Sha256 `protobuf:"bytes,1,opt,name=last_work,json=lastWork,proto3" json:"last_work,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SoMemWitnessByLastWork) Reset()         { *m = SoMemWitnessByLastWork{} }
func (m *SoMemWitnessByLastWork) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByLastWork) ProtoMessage()    {}
func (*SoMemWitnessByLastWork) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_witness_3e7e301b243138ab, []int{9}
}
func (m *SoMemWitnessByLastWork) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByLastWork.Unmarshal(m, b)
}
func (m *SoMemWitnessByLastWork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByLastWork.Marshal(b, m, deterministic)
}
func (dst *SoMemWitnessByLastWork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByLastWork.Merge(dst, src)
}
func (m *SoMemWitnessByLastWork) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByLastWork.Size(m)
}
func (m *SoMemWitnessByLastWork) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByLastWork.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByLastWork proto.InternalMessageInfo

func (m *SoMemWitnessByLastWork) GetLastWork() *prototype.Sha256 {
	if m != nil {
		return m.LastWork
	}
	return nil
}

type SoMemWitnessByRunningVersion struct {
	RunningVersion       uint32   `protobuf:"varint,1,opt,name=running_version,json=runningVersion,proto3" json:"running_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemWitnessByRunningVersion) Reset()         { *m = SoMemWitnessByRunningVersion{} }
func (m *SoMemWitnessByRunningVersion) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByRunningVersion) ProtoMessage()    {}
func (*SoMemWitnessByRunningVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_witness_3e7e301b243138ab, []int{10}
}
func (m *SoMemWitnessByRunningVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByRunningVersion.Unmarshal(m, b)
}
func (m *SoMemWitnessByRunningVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByRunningVersion.Marshal(b, m, deterministic)
}
func (dst *SoMemWitnessByRunningVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByRunningVersion.Merge(dst, src)
}
func (m *SoMemWitnessByRunningVersion) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByRunningVersion.Size(m)
}
func (m *SoMemWitnessByRunningVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByRunningVersion.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByRunningVersion proto.InternalMessageInfo

func (m *SoMemWitnessByRunningVersion) GetRunningVersion() uint32 {
	if m != nil {
		return m.RunningVersion
	}
	return 0
}

type SoMemWitnessByLastAslot struct {
	LastAslot            uint32   `protobuf:"varint,1,opt,name=last_aslot,json=lastAslot,proto3" json:"last_aslot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemWitnessByLastAslot) Reset()         { *m = SoMemWitnessByLastAslot{} }
func (m *SoMemWitnessByLastAslot) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByLastAslot) ProtoMessage()    {}
func (*SoMemWitnessByLastAslot) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_witness_3e7e301b243138ab, []int{11}
}
func (m *SoMemWitnessByLastAslot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByLastAslot.Unmarshal(m, b)
}
func (m *SoMemWitnessByLastAslot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByLastAslot.Marshal(b, m, deterministic)
}
func (dst *SoMemWitnessByLastAslot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByLastAslot.Merge(dst, src)
}
func (m *SoMemWitnessByLastAslot) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByLastAslot.Size(m)
}
func (m *SoMemWitnessByLastAslot) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByLastAslot.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByLastAslot proto.InternalMessageInfo

func (m *SoMemWitnessByLastAslot) GetLastAslot() uint32 {
	if m != nil {
		return m.LastAslot
	}
	return 0
}

type SoListWitnessByOwner struct {
	Owner                *prototype.AccountName `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoListWitnessByOwner) Reset()         { *m = SoListWitnessByOwner{} }
func (m *SoListWitnessByOwner) String() string { return proto.CompactTextString(m) }
func (*SoListWitnessByOwner) ProtoMessage()    {}
func (*SoListWitnessByOwner) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_witness_3e7e301b243138ab, []int{12}
}
func (m *SoListWitnessByOwner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListWitnessByOwner.Unmarshal(m, b)
}
func (m *SoListWitnessByOwner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListWitnessByOwner.Marshal(b, m, deterministic)
}
func (dst *SoListWitnessByOwner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListWitnessByOwner.Merge(dst, src)
}
func (m *SoListWitnessByOwner) XXX_Size() int {
	return xxx_messageInfo_SoListWitnessByOwner.Size(m)
}
func (m *SoListWitnessByOwner) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListWitnessByOwner.DiscardUnknown(m)
}

var xxx_messageInfo_SoListWitnessByOwner proto.InternalMessageInfo

func (m *SoListWitnessByOwner) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

type SoListWitnessByVoteCount struct {
	VoteCount            uint64                 `protobuf:"varint,1,opt,name=vote_count,json=voteCount,proto3" json:"vote_count,omitempty"`
	Owner                *prototype.AccountName `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoListWitnessByVoteCount) Reset()         { *m = SoListWitnessByVoteCount{} }
func (m *SoListWitnessByVoteCount) String() string { return proto.CompactTextString(m) }
func (*SoListWitnessByVoteCount) ProtoMessage()    {}
func (*SoListWitnessByVoteCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_witness_3e7e301b243138ab, []int{13}
}
func (m *SoListWitnessByVoteCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListWitnessByVoteCount.Unmarshal(m, b)
}
func (m *SoListWitnessByVoteCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListWitnessByVoteCount.Marshal(b, m, deterministic)
}
func (dst *SoListWitnessByVoteCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListWitnessByVoteCount.Merge(dst, src)
}
func (m *SoListWitnessByVoteCount) XXX_Size() int {
	return xxx_messageInfo_SoListWitnessByVoteCount.Size(m)
}
func (m *SoListWitnessByVoteCount) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListWitnessByVoteCount.DiscardUnknown(m)
}

var xxx_messageInfo_SoListWitnessByVoteCount proto.InternalMessageInfo

func (m *SoListWitnessByVoteCount) GetVoteCount() uint64 {
	if m != nil {
		return m.VoteCount
	}
	return 0
}

func (m *SoListWitnessByVoteCount) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

type SoUniqueWitnessByOwner struct {
	Owner                *prototype.AccountName `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoUniqueWitnessByOwner) Reset()         { *m = SoUniqueWitnessByOwner{} }
func (m *SoUniqueWitnessByOwner) String() string { return proto.CompactTextString(m) }
func (*SoUniqueWitnessByOwner) ProtoMessage()    {}
func (*SoUniqueWitnessByOwner) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_witness_3e7e301b243138ab, []int{14}
}
func (m *SoUniqueWitnessByOwner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueWitnessByOwner.Unmarshal(m, b)
}
func (m *SoUniqueWitnessByOwner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueWitnessByOwner.Marshal(b, m, deterministic)
}
func (dst *SoUniqueWitnessByOwner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueWitnessByOwner.Merge(dst, src)
}
func (m *SoUniqueWitnessByOwner) XXX_Size() int {
	return xxx_messageInfo_SoUniqueWitnessByOwner.Size(m)
}
func (m *SoUniqueWitnessByOwner) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueWitnessByOwner.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueWitnessByOwner proto.InternalMessageInfo

func (m *SoUniqueWitnessByOwner) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

func init() {
	proto.RegisterType((*SoWitness)(nil), "table.so_witness")
	proto.RegisterType((*SoMemWitnessByOwner)(nil), "table.so_mem_witness_by_owner")
	proto.RegisterType((*SoMemWitnessByCreatedTime)(nil), "table.so_mem_witness_by_created_time")
	proto.RegisterType((*SoMemWitnessByUrl)(nil), "table.so_mem_witness_by_url")
	proto.RegisterType((*SoMemWitnessByVoteCount)(nil), "table.so_mem_witness_by_vote_count")
	proto.RegisterType((*SoMemWitnessByLastConfirmedBlockNum)(nil), "table.so_mem_witness_by_last_confirmed_block_num")
	proto.RegisterType((*SoMemWitnessByTotalMissed)(nil), "table.so_mem_witness_by_total_missed")
	proto.RegisterType((*SoMemWitnessByPowWorker)(nil), "table.so_mem_witness_by_pow_worker")
	proto.RegisterType((*SoMemWitnessBySigningKey)(nil), "table.so_mem_witness_by_signing_key")
	proto.RegisterType((*SoMemWitnessByLastWork)(nil), "table.so_mem_witness_by_last_work")
	proto.RegisterType((*SoMemWitnessByRunningVersion)(nil), "table.so_mem_witness_by_running_version")
	proto.RegisterType((*SoMemWitnessByLastAslot)(nil), "table.so_mem_witness_by_last_aslot")
	proto.RegisterType((*SoListWitnessByOwner)(nil), "table.so_list_witness_by_owner")
	proto.RegisterType((*SoListWitnessByVoteCount)(nil), "table.so_list_witness_by_vote_count")
	proto.RegisterType((*SoUniqueWitnessByOwner)(nil), "table.so_unique_witness_by_owner")
}

func init() {
	proto.RegisterFile("app/table/so_witness.proto", fileDescriptor_so_witness_3e7e301b243138ab)
}

var fileDescriptor_so_witness_3e7e301b243138ab = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x5d, 0x6f, 0xd3, 0x3c,
	0x14, 0xc7, 0xe5, 0x75, 0xdd, 0xb3, 0x9e, 0xee, 0xe1, 0x25, 0x62, 0x9a, 0x29, 0x0c, 0x75, 0xb9,
	0x80, 0x82, 0xb4, 0x56, 0x1a, 0x02, 0x2e, 0x80, 0x0b, 0xd6, 0x1b, 0xd0, 0x18, 0x17, 0x11, 0x02,
	0x09, 0x21, 0xac, 0x24, 0x35, 0xad, 0xd5, 0xd8, 0x27, 0xc4, 0xce, 0xaa, 0x7e, 0x55, 0x3e, 0x0d,
	0xb2, 0xd3, 0xad, 0x21, 0xe9, 0xc6, 0xaa, 0xdd, 0x54, 0xf5, 0xff, 0xbc, 0xda, 0xe7, 0x77, 0x14,
	0xe8, 0x84, 0x69, 0x3a, 0x30, 0x61, 0x94, 0xf0, 0x81, 0x46, 0x36, 0x13, 0x46, 0x71, 0xad, 0xfb,
	0x69, 0x86, 0x06, 0xbd, 0xa6, 0xd3, 0x3b, 0xf7, 0xdc, 0xc9, 0xcc, 0x53, 0x3e, 0xb0, 0x3f, 0x85,
	0xd1, 0xff, 0xdd, 0x00, 0x58, 0x46, 0x78, 0x87, 0xd0, 0xc4, 0x99, 0xe2, 0x19, 0x25, 0x5d, 0xd2,
	0x6b, 0x1f, 0xed, 0xf5, 0x2f, 0x82, 0xfa, 0x61, 0x1c, 0x63, 0xae, 0x0c, 0x53, 0xa1, 0xe4, 0x41,
	0xe1, 0xe5, 0xbd, 0x81, 0x9d, 0x38, 0xe3, 0xa1, 0xe1, 0x23, 0x66, 0x84, 0xe4, 0x74, 0xc3, 0x45,
	0xdd, 0x2f, 0x45, 0x59, 0x99, 0xa5, 0x28, 0x94, 0x61, 0x9a, 0xc7, 0x41, 0x7b, 0xe1, 0xfe, 0x59,
	0x48, 0xee, 0xdd, 0x81, 0x46, 0x9e, 0x25, 0xb4, 0xd1, 0x25, 0xbd, 0x56, 0x60, 0xff, 0x7a, 0xfb,
	0x00, 0x67, 0x68, 0x38, 0x73, 0x95, 0xe8, 0x66, 0x97, 0xf4, 0x36, 0x83, 0x96, 0x55, 0x86, 0x56,
	0xf0, 0x5e, 0x01, 0x4d, 0x42, 0x6d, 0x58, 0x8c, 0xea, 0xa7, 0xc8, 0x24, 0x1f, 0xb1, 0x28, 0xc1,
	0x78, 0xca, 0x54, 0x2e, 0x69, 0xb3, 0x4b, 0x7a, 0xff, 0x07, 0xbb, 0xd6, 0x3e, 0x3c, 0x37, 0x1f,
	0x5b, 0xeb, 0xa7, 0x5c, 0x7a, 0x07, 0xb0, 0x63, 0xd0, 0x84, 0x09, 0x93, 0x42, 0x6b, 0x3e, 0xa2,
	0x5b, 0xce, 0xb9, 0xed, 0xb4, 0x53, 0x27, 0xd9, 0xd2, 0x29, 0xce, 0xd8, 0x0c, 0xb3, 0x29, 0xcf,
	0xe8, 0x7f, 0xce, 0xa1, 0x95, 0xe2, 0xec, 0xab, 0x13, 0xbc, 0xd7, 0xd0, 0xd6, 0x62, 0xac, 0x84,
	0x1a, 0xb3, 0x29, 0x9f, 0xd3, 0x6d, 0x77, 0xd1, 0x4e, 0xe9, 0xa2, 0x69, 0x1e, 0x25, 0x22, 0xb6,
	0x46, 0x66, 0xcf, 0x01, 0x2c, 0xdc, 0x4f, 0xf8, 0xdc, 0xeb, 0x43, 0xcb, 0xf5, 0x6d, 0x93, 0xd3,
	0x96, 0x0b, 0xbd, 0x5b, 0x0a, 0xd5, 0x93, 0xf0, 0xe8, 0xc5, 0xcb, 0x60, 0xdb, 0xfa, 0xd8, 0x72,
	0xde, 0x13, 0xb8, 0x9d, 0xe5, 0xca, 0x15, 0x3b, 0xe3, 0x99, 0x16, 0xa8, 0x28, 0xb8, 0x86, 0x6e,
	0x2d, 0xe4, 0x2f, 0x85, 0x6a, 0x9b, 0x76, 0x89, 0x43, 0x9d, 0xa0, 0xa1, 0xed, 0xa2, 0x69, 0xab,
	0xbc, 0xb3, 0x82, 0xff, 0x1e, 0xf6, 0x34, 0x32, 0xc9, 0xe5, 0xf9, 0x7c, 0x59, 0x34, 0x67, 0xc5,
	0xe4, 0xd6, 0x1b, 0xb4, 0xff, 0x03, 0x1e, 0xd5, 0x33, 0x95, 0x47, 0x5f, 0x43, 0x81, 0xac, 0x83,
	0x82, 0xff, 0x14, 0x76, 0xeb, 0xf9, 0x2d, 0x11, 0x0b, 0x46, 0xc8, 0x05, 0x23, 0xfe, 0x5b, 0x78,
	0x58, 0x77, 0x5d, 0x52, 0x53, 0x61, 0x88, 0x54, 0x18, 0xf2, 0x39, 0x3c, 0xab, 0x87, 0x5f, 0x46,
	0xd5, 0x95, 0xc4, 0x91, 0x2b, 0x88, 0xf3, 0x87, 0xab, 0x1e, 0xac, 0xcc, 0x60, 0x8d, 0x49, 0x52,
	0x63, 0x72, 0xf5, 0x55, 0x97, 0x94, 0x56, 0x98, 0x25, 0x15, 0x66, 0xfd, 0xef, 0xb0, 0x5f, 0x0f,
	0x2f, 0x51, 0x5c, 0x85, 0x9a, 0xac, 0x03, 0xb5, 0x7f, 0x0a, 0x0f, 0x2e, 0x79, 0x48, 0xdb, 0xcf,
	0xdf, 0xcc, 0x93, 0x7f, 0x32, 0xef, 0x7f, 0x84, 0x83, 0x7a, 0xba, 0xca, 0x16, 0xac, 0x5a, 0x0c,
	0xb2, 0x6a, 0x31, 0x56, 0xbf, 0xdc, 0x72, 0x55, 0x2a, 0x8b, 0x43, 0xaa, 0x8b, 0xf3, 0x01, 0xa8,
	0x46, 0x96, 0x08, 0xdb, 0xff, 0x0d, 0x37, 0x47, 0xba, 0x21, 0x54, 0x53, 0x5d, 0x9b, 0xd7, 0x65,
	0xb9, 0x8d, 0x6b, 0x95, 0x3b, 0x81, 0x8e, 0x46, 0x96, 0x2b, 0xf1, 0x2b, 0xe7, 0x37, 0xed, 0xfd,
	0xb8, 0xf7, 0xed, 0xf1, 0x58, 0x98, 0x49, 0x1e, 0xf5, 0x63, 0x94, 0x83, 0x18, 0x75, 0x3c, 0x09,
	0x85, 0x1a, 0xc4, 0xa8, 0x0c, 0x57, 0x06, 0xf5, 0xe1, 0x18, 0x8b, 0x8f, 0x4e, 0xb4, 0xe5, 0x12,
	0x3d, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xa5, 0x1e, 0x8a, 0x88, 0x06, 0x00, 0x00,
}
