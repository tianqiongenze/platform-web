// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mem/mem.proto

package mem

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	basic "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/basic"
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

type VirtualMemoryStat struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Total                uint64               `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"`
	Available            uint64               `protobuf:"varint,3,opt,name=Available,proto3" json:"Available,omitempty"`
	Used                 uint64               `protobuf:"varint,4,opt,name=Used,proto3" json:"Used,omitempty"`
	UsedPercent          float64              `protobuf:"fixed64,5,opt,name=UsedPercent,proto3" json:"UsedPercent,omitempty"`
	Free                 uint64               `protobuf:"varint,6,opt,name=Free,proto3" json:"Free,omitempty"`
	Active               uint64               `protobuf:"varint,7,opt,name=Active,proto3" json:"Active,omitempty"`
	Inactive             uint64               `protobuf:"varint,8,opt,name=Inactive,proto3" json:"Inactive,omitempty"`
	Wired                uint64               `protobuf:"varint,9,opt,name=Wired,proto3" json:"Wired,omitempty"`
	Laundry              uint64               `protobuf:"varint,10,opt,name=Laundry,proto3" json:"Laundry,omitempty"`
	Buffers              uint64               `protobuf:"varint,11,opt,name=Buffers,proto3" json:"Buffers,omitempty"`
	Cached               uint64               `protobuf:"varint,12,opt,name=Cached,proto3" json:"Cached,omitempty"`
	Writeback            uint64               `protobuf:"varint,13,opt,name=Writeback,proto3" json:"Writeback,omitempty"`
	Dirty                uint64               `protobuf:"varint,14,opt,name=Dirty,proto3" json:"Dirty,omitempty"`
	WritebackTmp         uint64               `protobuf:"varint,15,opt,name=WritebackTmp,proto3" json:"WritebackTmp,omitempty"`
	Shared               uint64               `protobuf:"varint,16,opt,name=Shared,proto3" json:"Shared,omitempty"`
	Slab                 uint64               `protobuf:"varint,17,opt,name=Slab,proto3" json:"Slab,omitempty"`
	SReclaimable         uint64               `protobuf:"varint,18,opt,name=SReclaimable,proto3" json:"SReclaimable,omitempty"`
	PageTables           uint64               `protobuf:"varint,19,opt,name=PageTables,proto3" json:"PageTables,omitempty"`
	SwapCached           uint64               `protobuf:"varint,20,opt,name=SwapCached,proto3" json:"SwapCached,omitempty"`
	CommitLimit          uint64               `protobuf:"varint,21,opt,name=CommitLimit,proto3" json:"CommitLimit,omitempty"`
	CommittedAS          uint64               `protobuf:"varint,22,opt,name=CommittedAS,proto3" json:"CommittedAS,omitempty"`
	HighTotal            uint64               `protobuf:"varint,23,opt,name=HighTotal,proto3" json:"HighTotal,omitempty"`
	HighFree             uint64               `protobuf:"varint,24,opt,name=HighFree,proto3" json:"HighFree,omitempty"`
	LowTotal             uint64               `protobuf:"varint,25,opt,name=LowTotal,proto3" json:"LowTotal,omitempty"`
	LowFree              uint64               `protobuf:"varint,26,opt,name=LowFree,proto3" json:"LowFree,omitempty"`
	SwapTotal            uint64               `protobuf:"varint,27,opt,name=SwapTotal,proto3" json:"SwapTotal,omitempty"`
	SwapFree             uint64               `protobuf:"varint,28,opt,name=SwapFree,proto3" json:"SwapFree,omitempty"`
	Mapped               uint64               `protobuf:"varint,29,opt,name=Mapped,proto3" json:"Mapped,omitempty"`
	VMallocTotal         uint64               `protobuf:"varint,30,opt,name=VMallocTotal,proto3" json:"VMallocTotal,omitempty"`
	VMallocUsed          uint64               `protobuf:"varint,31,opt,name=VMallocUsed,proto3" json:"VMallocUsed,omitempty"`
	VMallocChunk         uint64               `protobuf:"varint,32,opt,name=VMallocChunk,proto3" json:"VMallocChunk,omitempty"`
	HugePagesTotal       uint64               `protobuf:"varint,33,opt,name=HugePagesTotal,proto3" json:"HugePagesTotal,omitempty"`
	HugePagesFree        uint64               `protobuf:"varint,34,opt,name=HugePagesFree,proto3" json:"HugePagesFree,omitempty"`
	HugePageSize         uint64               `protobuf:"varint,35,opt,name=HugePageSize,proto3" json:"HugePageSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *VirtualMemoryStat) Reset()         { *m = VirtualMemoryStat{} }
func (m *VirtualMemoryStat) String() string { return proto.CompactTextString(m) }
func (*VirtualMemoryStat) ProtoMessage()    {}
func (*VirtualMemoryStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d2907961eca7e66, []int{0}
}

func (m *VirtualMemoryStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMemoryStat.Unmarshal(m, b)
}
func (m *VirtualMemoryStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMemoryStat.Marshal(b, m, deterministic)
}
func (m *VirtualMemoryStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMemoryStat.Merge(m, src)
}
func (m *VirtualMemoryStat) XXX_Size() int {
	return xxx_messageInfo_VirtualMemoryStat.Size(m)
}
func (m *VirtualMemoryStat) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMemoryStat.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMemoryStat proto.InternalMessageInfo

func (m *VirtualMemoryStat) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *VirtualMemoryStat) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *VirtualMemoryStat) GetAvailable() uint64 {
	if m != nil {
		return m.Available
	}
	return 0
}

func (m *VirtualMemoryStat) GetUsed() uint64 {
	if m != nil {
		return m.Used
	}
	return 0
}

func (m *VirtualMemoryStat) GetUsedPercent() float64 {
	if m != nil {
		return m.UsedPercent
	}
	return 0
}

func (m *VirtualMemoryStat) GetFree() uint64 {
	if m != nil {
		return m.Free
	}
	return 0
}

func (m *VirtualMemoryStat) GetActive() uint64 {
	if m != nil {
		return m.Active
	}
	return 0
}

func (m *VirtualMemoryStat) GetInactive() uint64 {
	if m != nil {
		return m.Inactive
	}
	return 0
}

func (m *VirtualMemoryStat) GetWired() uint64 {
	if m != nil {
		return m.Wired
	}
	return 0
}

func (m *VirtualMemoryStat) GetLaundry() uint64 {
	if m != nil {
		return m.Laundry
	}
	return 0
}

func (m *VirtualMemoryStat) GetBuffers() uint64 {
	if m != nil {
		return m.Buffers
	}
	return 0
}

func (m *VirtualMemoryStat) GetCached() uint64 {
	if m != nil {
		return m.Cached
	}
	return 0
}

func (m *VirtualMemoryStat) GetWriteback() uint64 {
	if m != nil {
		return m.Writeback
	}
	return 0
}

func (m *VirtualMemoryStat) GetDirty() uint64 {
	if m != nil {
		return m.Dirty
	}
	return 0
}

func (m *VirtualMemoryStat) GetWritebackTmp() uint64 {
	if m != nil {
		return m.WritebackTmp
	}
	return 0
}

func (m *VirtualMemoryStat) GetShared() uint64 {
	if m != nil {
		return m.Shared
	}
	return 0
}

func (m *VirtualMemoryStat) GetSlab() uint64 {
	if m != nil {
		return m.Slab
	}
	return 0
}

func (m *VirtualMemoryStat) GetSReclaimable() uint64 {
	if m != nil {
		return m.SReclaimable
	}
	return 0
}

func (m *VirtualMemoryStat) GetPageTables() uint64 {
	if m != nil {
		return m.PageTables
	}
	return 0
}

func (m *VirtualMemoryStat) GetSwapCached() uint64 {
	if m != nil {
		return m.SwapCached
	}
	return 0
}

func (m *VirtualMemoryStat) GetCommitLimit() uint64 {
	if m != nil {
		return m.CommitLimit
	}
	return 0
}

func (m *VirtualMemoryStat) GetCommittedAS() uint64 {
	if m != nil {
		return m.CommittedAS
	}
	return 0
}

func (m *VirtualMemoryStat) GetHighTotal() uint64 {
	if m != nil {
		return m.HighTotal
	}
	return 0
}

func (m *VirtualMemoryStat) GetHighFree() uint64 {
	if m != nil {
		return m.HighFree
	}
	return 0
}

func (m *VirtualMemoryStat) GetLowTotal() uint64 {
	if m != nil {
		return m.LowTotal
	}
	return 0
}

func (m *VirtualMemoryStat) GetLowFree() uint64 {
	if m != nil {
		return m.LowFree
	}
	return 0
}

func (m *VirtualMemoryStat) GetSwapTotal() uint64 {
	if m != nil {
		return m.SwapTotal
	}
	return 0
}

func (m *VirtualMemoryStat) GetSwapFree() uint64 {
	if m != nil {
		return m.SwapFree
	}
	return 0
}

func (m *VirtualMemoryStat) GetMapped() uint64 {
	if m != nil {
		return m.Mapped
	}
	return 0
}

func (m *VirtualMemoryStat) GetVMallocTotal() uint64 {
	if m != nil {
		return m.VMallocTotal
	}
	return 0
}

func (m *VirtualMemoryStat) GetVMallocUsed() uint64 {
	if m != nil {
		return m.VMallocUsed
	}
	return 0
}

func (m *VirtualMemoryStat) GetVMallocChunk() uint64 {
	if m != nil {
		return m.VMallocChunk
	}
	return 0
}

func (m *VirtualMemoryStat) GetHugePagesTotal() uint64 {
	if m != nil {
		return m.HugePagesTotal
	}
	return 0
}

func (m *VirtualMemoryStat) GetHugePagesFree() uint64 {
	if m != nil {
		return m.HugePagesFree
	}
	return 0
}

func (m *VirtualMemoryStat) GetHugePageSize() uint64 {
	if m != nil {
		return m.HugePageSize
	}
	return 0
}

type SwapMemoryStat struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Total                uint64               `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"`
	Used                 uint64               `protobuf:"varint,3,opt,name=Used,proto3" json:"Used,omitempty"`
	Free                 uint64               `protobuf:"varint,4,opt,name=Free,proto3" json:"Free,omitempty"`
	UsedPercent          float64              `protobuf:"fixed64,5,opt,name=UsedPercent,proto3" json:"UsedPercent,omitempty"`
	Sin                  uint64               `protobuf:"varint,6,opt,name=Sin,proto3" json:"Sin,omitempty"`
	Sout                 uint64               `protobuf:"varint,7,opt,name=Sout,proto3" json:"Sout,omitempty"`
	PgIn                 uint64               `protobuf:"varint,8,opt,name=PgIn,proto3" json:"PgIn,omitempty"`
	PgOut                uint64               `protobuf:"varint,9,opt,name=PgOut,proto3" json:"PgOut,omitempty"`
	PgFault              uint64               `protobuf:"varint,10,opt,name=PgFault,proto3" json:"PgFault,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SwapMemoryStat) Reset()         { *m = SwapMemoryStat{} }
func (m *SwapMemoryStat) String() string { return proto.CompactTextString(m) }
func (*SwapMemoryStat) ProtoMessage()    {}
func (*SwapMemoryStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d2907961eca7e66, []int{1}
}

func (m *SwapMemoryStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwapMemoryStat.Unmarshal(m, b)
}
func (m *SwapMemoryStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwapMemoryStat.Marshal(b, m, deterministic)
}
func (m *SwapMemoryStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwapMemoryStat.Merge(m, src)
}
func (m *SwapMemoryStat) XXX_Size() int {
	return xxx_messageInfo_SwapMemoryStat.Size(m)
}
func (m *SwapMemoryStat) XXX_DiscardUnknown() {
	xxx_messageInfo_SwapMemoryStat.DiscardUnknown(m)
}

var xxx_messageInfo_SwapMemoryStat proto.InternalMessageInfo

func (m *SwapMemoryStat) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *SwapMemoryStat) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *SwapMemoryStat) GetUsed() uint64 {
	if m != nil {
		return m.Used
	}
	return 0
}

func (m *SwapMemoryStat) GetFree() uint64 {
	if m != nil {
		return m.Free
	}
	return 0
}

func (m *SwapMemoryStat) GetUsedPercent() float64 {
	if m != nil {
		return m.UsedPercent
	}
	return 0
}

func (m *SwapMemoryStat) GetSin() uint64 {
	if m != nil {
		return m.Sin
	}
	return 0
}

func (m *SwapMemoryStat) GetSout() uint64 {
	if m != nil {
		return m.Sout
	}
	return 0
}

func (m *SwapMemoryStat) GetPgIn() uint64 {
	if m != nil {
		return m.PgIn
	}
	return 0
}

func (m *SwapMemoryStat) GetPgOut() uint64 {
	if m != nil {
		return m.PgOut
	}
	return 0
}

func (m *SwapMemoryStat) GetPgFault() uint64 {
	if m != nil {
		return m.PgFault
	}
	return 0
}

type MemRequest struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	IP                   string               `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	NodeName             string               `protobuf:"bytes,3,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	VirtualMemoryStat    []*VirtualMemoryStat `protobuf:"bytes,4,rep,name=virtualMemoryStat,proto3" json:"virtualMemoryStat,omitempty"`
	SwapMemoryStat       []*SwapMemoryStat    `protobuf:"bytes,5,rep,name=swapMemoryStat,proto3" json:"swapMemoryStat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MemRequest) Reset()         { *m = MemRequest{} }
func (m *MemRequest) String() string { return proto.CompactTextString(m) }
func (*MemRequest) ProtoMessage()    {}
func (*MemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d2907961eca7e66, []int{2}
}

func (m *MemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemRequest.Unmarshal(m, b)
}
func (m *MemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemRequest.Marshal(b, m, deterministic)
}
func (m *MemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemRequest.Merge(m, src)
}
func (m *MemRequest) XXX_Size() int {
	return xxx_messageInfo_MemRequest.Size(m)
}
func (m *MemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MemRequest proto.InternalMessageInfo

func (m *MemRequest) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *MemRequest) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *MemRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *MemRequest) GetVirtualMemoryStat() []*VirtualMemoryStat {
	if m != nil {
		return m.VirtualMemoryStat
	}
	return nil
}

func (m *MemRequest) GetSwapMemoryStat() []*SwapMemoryStat {
	if m != nil {
		return m.SwapMemoryStat
	}
	return nil
}

type MemResponse struct {
	Error                *basic.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MemResponse) Reset()         { *m = MemResponse{} }
func (m *MemResponse) String() string { return proto.CompactTextString(m) }
func (*MemResponse) ProtoMessage()    {}
func (*MemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d2907961eca7e66, []int{3}
}

func (m *MemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemResponse.Unmarshal(m, b)
}
func (m *MemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemResponse.Marshal(b, m, deterministic)
}
func (m *MemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemResponse.Merge(m, src)
}
func (m *MemResponse) XXX_Size() int {
	return xxx_messageInfo_MemResponse.Size(m)
}
func (m *MemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MemResponse proto.InternalMessageInfo

func (m *MemResponse) GetError() *basic.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*VirtualMemoryStat)(nil), "protobuf.pb.mem.VirtualMemoryStat")
	proto.RegisterType((*SwapMemoryStat)(nil), "protobuf.pb.mem.SwapMemoryStat")
	proto.RegisterType((*MemRequest)(nil), "protobuf.pb.mem.MemRequest")
	proto.RegisterType((*MemResponse)(nil), "protobuf.pb.mem.MemResponse")
}

func init() { proto.RegisterFile("mem/mem.proto", fileDescriptor_6d2907961eca7e66) }

var fileDescriptor_6d2907961eca7e66 = []byte{
	// 856 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x26, 0x69, 0xda, 0x6d, 0x26, 0xdb, 0xec, 0x76, 0xd8, 0x2e, 0x43, 0xb6, 0x6c, 0x43, 0x40,
	0xa8, 0x12, 0xaa, 0x23, 0x95, 0x1b, 0xc4, 0x5d, 0xb7, 0xfb, 0x57, 0xa9, 0x05, 0x63, 0x87, 0x5d,
	0x89, 0xbb, 0xb1, 0x33, 0x75, 0x46, 0xf5, 0x78, 0xcc, 0x78, 0xdc, 0xaa, 0xbc, 0x03, 0x77, 0xbc,
	0x05, 0xf7, 0x3c, 0x1f, 0x3a, 0x73, 0x6c, 0xc7, 0x49, 0x90, 0xb8, 0xa8, 0xb8, 0xf2, 0x7c, 0xdf,
	0xf9, 0x9b, 0xf9, 0xce, 0xf8, 0x0c, 0xd9, 0x53, 0x42, 0x4d, 0x95, 0x50, 0x5e, 0x6e, 0xb4, 0xd5,
	0xf4, 0x89, 0xfb, 0x44, 0xe5, 0xb5, 0x97, 0x47, 0x9e, 0x12, 0x6a, 0xb4, 0x1f, 0xf1, 0x42, 0xc6,
	0x53, 0x61, 0x8c, 0x36, 0xe8, 0x33, 0x3a, 0x4a, 0xb4, 0x4e, 0x52, 0x31, 0xad, 0x5d, 0xa7, 0x56,
	0x2a, 0x51, 0x58, 0xae, 0x72, 0x74, 0x98, 0xfc, 0xb5, 0x4b, 0xf6, 0x3f, 0x48, 0x63, 0x4b, 0x9e,
	0x5e, 0x09, 0xa5, 0xcd, 0x7d, 0x68, 0xb9, 0xa5, 0xdf, 0x93, 0x7e, 0xe3, 0xc8, 0x3a, 0xe3, 0xce,
	0xf1, 0xe0, 0x74, 0xe4, 0x61, 0x2a, 0xaf, 0xa9, 0x3a, 0xab, 0x3d, 0x82, 0xa5, 0x33, 0x7d, 0x46,
	0xb6, 0x67, 0xda, 0xf2, 0x94, 0x75, 0xc7, 0x9d, 0xe3, 0x5e, 0x80, 0x80, 0x1e, 0x92, 0xfe, 0xd9,
	0x2d, 0x97, 0x29, 0x8f, 0x52, 0xc1, 0xb6, 0x9c, 0x65, 0x49, 0x50, 0x4a, 0x7a, 0xbf, 0x14, 0x62,
	0xce, 0x7a, 0xce, 0xe0, 0xd6, 0x74, 0x4c, 0x06, 0xf0, 0xf5, 0x85, 0x89, 0x45, 0x66, 0xd9, 0xf6,
	0xb8, 0x73, 0xdc, 0x09, 0xda, 0x14, 0x44, 0xbd, 0x35, 0x42, 0xb0, 0x1d, 0x8c, 0x82, 0x35, 0x7d,
	0x4e, 0x76, 0xce, 0x62, 0x2b, 0x6f, 0x05, 0x7b, 0xe4, 0xd8, 0x0a, 0xd1, 0x11, 0xd9, 0xbd, 0xc8,
	0x38, 0x5a, 0x76, 0x9d, 0xa5, 0xc1, 0xb0, 0xe3, 0x8f, 0xd2, 0x88, 0x39, 0xeb, 0xe3, 0x8e, 0x1d,
	0xa0, 0x8c, 0x3c, 0xba, 0xe4, 0x65, 0x36, 0x37, 0xf7, 0x8c, 0x38, 0xbe, 0x86, 0x60, 0x79, 0x55,
	0x5e, 0x5f, 0x0b, 0x53, 0xb0, 0x01, 0x5a, 0x2a, 0x08, 0xd5, 0xcf, 0x79, 0xbc, 0x10, 0x73, 0xf6,
	0x18, 0xab, 0x23, 0x82, 0xd3, 0x7f, 0x34, 0xd2, 0x8a, 0x88, 0xc7, 0x37, 0x6c, 0x0f, 0x4f, 0xdf,
	0x10, 0x50, 0xff, 0xb5, 0x34, 0xf6, 0x9e, 0x0d, 0xb1, 0xbe, 0x03, 0x74, 0x42, 0x1e, 0x37, 0x2e,
	0x33, 0x95, 0xb3, 0x27, 0xce, 0xb8, 0xc2, 0x41, 0xbd, 0x70, 0xc1, 0x61, 0xeb, 0x4f, 0xb1, 0x1e,
	0x22, 0x50, 0x26, 0x4c, 0x79, 0xc4, 0xf6, 0x51, 0x19, 0x58, 0x43, 0xbe, 0x30, 0x10, 0x71, 0xca,
	0xa5, 0x72, 0x4d, 0xa0, 0x98, 0xaf, 0xcd, 0xd1, 0x97, 0x84, 0xf8, 0x3c, 0x11, 0x33, 0x00, 0x05,
	0xfb, 0xd4, 0x79, 0xb4, 0x18, 0xb0, 0x87, 0x77, 0x3c, 0xaf, 0xce, 0xf8, 0x0c, 0xed, 0x4b, 0x06,
	0x7a, 0x76, 0xae, 0x95, 0x92, 0xf6, 0x52, 0x2a, 0x69, 0xd9, 0x81, 0x73, 0x68, 0x53, 0x4b, 0x0f,
	0x2b, 0xe6, 0x67, 0x21, 0x7b, 0xde, 0xf6, 0x70, 0x14, 0x68, 0xf5, 0x5e, 0x26, 0x0b, 0xbc, 0x43,
	0x9f, 0xa1, 0x56, 0x0d, 0x01, 0x7d, 0x04, 0xe0, 0xfa, 0xce, 0xb0, 0x8f, 0x35, 0x06, 0xdb, 0xa5,
	0xbe, 0xc3, 0xc0, 0xcf, 0xd1, 0x56, 0x63, 0xd7, 0x4d, 0x7d, 0xe7, 0xc2, 0x46, 0x55, 0x37, 0x11,
	0x42, 0x3d, 0x38, 0x01, 0x86, 0xbd, 0xc0, 0x7a, 0x0d, 0x01, 0x39, 0x01, 0xb8, 0xc0, 0x43, 0xcc,
	0x59, 0x63, 0x50, 0xff, 0x8a, 0xe7, 0xb9, 0x98, 0xb3, 0x2f, 0x50, 0x7d, 0x44, 0xa0, 0xf4, 0x87,
	0x2b, 0x9e, 0xa6, 0x3a, 0xc6, 0xa4, 0x2f, 0x51, 0xe9, 0x36, 0x07, 0x3a, 0x54, 0xd8, 0x5d, 0xfc,
	0x23, 0xd4, 0xa1, 0x45, 0xb5, 0xb2, 0x9c, 0x2f, 0xca, 0xec, 0x86, 0x8d, 0x57, 0xb2, 0x38, 0x8e,
	0x7e, 0x43, 0x86, 0xef, 0xcb, 0x44, 0x40, 0x87, 0x0a, 0xac, 0xf5, 0xa5, 0xf3, 0x5a, 0x63, 0xe9,
	0xd7, 0x64, 0xaf, 0x61, 0xdc, 0x51, 0x26, 0xce, 0x6d, 0x95, 0x84, 0x8a, 0x35, 0x11, 0xca, 0xdf,
	0x05, 0xfb, 0x0a, 0x2b, 0xb6, 0xb9, 0xc9, 0x9f, 0x5d, 0x32, 0x04, 0x01, 0xfe, 0xc7, 0x51, 0x51,
	0x0f, 0x83, 0xad, 0xd6, 0x30, 0xa8, 0x7f, 0xf5, 0x5e, 0xeb, 0x57, 0xff, 0xef, 0x01, 0xf1, 0x94,
	0x6c, 0x85, 0x32, 0xab, 0xe6, 0x03, 0x2c, 0xdd, 0x8f, 0xa1, 0x4b, 0x5b, 0x0d, 0x07, 0xb7, 0x06,
	0xce, 0x4f, 0x2e, 0xb2, 0x6a, 0x2c, 0xb8, 0x35, 0xec, 0xcc, 0x4f, 0x7e, 0x2a, 0x6d, 0x3d, 0x12,
	0x1c, 0x80, 0x4b, 0xe4, 0x27, 0x6f, 0x79, 0x99, 0xda, 0x7a, 0x24, 0x54, 0x70, 0xf2, 0x47, 0x97,
	0x90, 0x2b, 0xa1, 0x02, 0xf1, 0x5b, 0x29, 0x8a, 0x87, 0x48, 0x32, 0x24, 0xdd, 0x0b, 0xdf, 0xe9,
	0xd1, 0x0f, 0xba, 0x17, 0x3e, 0xdc, 0xbf, 0x4c, 0xcf, 0xc5, 0x8f, 0x5c, 0xe1, 0xd8, 0xec, 0x07,
	0x0d, 0xa6, 0x3e, 0xd9, 0xbf, 0x5d, 0x1f, 0xdc, 0xac, 0x37, 0xde, 0x3a, 0x1e, 0x9c, 0x4e, 0xbc,
	0xb5, 0xa7, 0xc1, 0xdb, 0x18, 0xf1, 0xc1, 0x66, 0x30, 0x7d, 0x47, 0x86, 0xc5, 0x4a, 0x73, 0xd9,
	0xb6, 0x4b, 0x77, 0xb4, 0x91, 0x6e, 0xf5, 0x0e, 0x04, 0x6b, 0x61, 0x93, 0x1f, 0xc8, 0xc0, 0xc9,
	0x51, 0xe4, 0x3a, 0x2b, 0x04, 0xfd, 0x96, 0x6c, 0xbb, 0x37, 0xa9, 0xd2, 0xe2, 0x60, 0x99, 0xce,
	0x3d, 0x58, 0xde, 0x1b, 0x30, 0x06, 0xe8, 0x73, 0xfa, 0x77, 0x87, 0x0c, 0x5e, 0xcb, 0xe2, 0x26,
	0x14, 0xe6, 0x56, 0xc6, 0x82, 0xce, 0xc8, 0x81, 0x5f, 0x16, 0x8b, 0xcd, 0x37, 0xea, 0xc5, 0xc6,
	0xae, 0x96, 0x2d, 0x18, 0x1d, 0xfe, 0xbb, 0x11, 0x37, 0x34, 0xf9, 0x84, 0xfe, 0x4c, 0x28, 0x64,
	0x5d, 0xbb, 0xcb, 0x0f, 0x49, 0xf9, 0xea, 0xdd, 0xaf, 0x6f, 0x12, 0x69, 0x17, 0x65, 0xe4, 0xc5,
	0x5a, 0x4d, 0x95, 0x8c, 0x8d, 0x3e, 0x91, 0xd9, 0x49, 0x9c, 0x4d, 0xf3, 0x94, 0xdb, 0x6b, 0x6d,
	0xd4, 0xc9, 0x9d, 0x88, 0xa6, 0xbc, 0x28, 0x84, 0x8a, 0xd2, 0xfb, 0x93, 0x54, 0x66, 0xad, 0x77,
	0x39, 0xd1, 0xd3, 0xea, 0x75, 0x8f, 0x76, 0x1c, 0xf9, 0xdd, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xcf, 0x44, 0x11, 0x30, 0xef, 0x07, 0x00, 0x00,
}
