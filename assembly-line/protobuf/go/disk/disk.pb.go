// Code generated by protoc-gen-go. DO NOT EDIT.
// source: disk/disk.proto

package disk

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

type UsageStat struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Path                 string               `protobuf:"bytes,2,opt,name=Path,proto3" json:"Path,omitempty"`
	Fstype               string               `protobuf:"bytes,3,opt,name=Fstype,proto3" json:"Fstype,omitempty"`
	Total                uint64               `protobuf:"varint,4,opt,name=Total,proto3" json:"Total,omitempty"`
	Free                 uint64               `protobuf:"varint,5,opt,name=Free,proto3" json:"Free,omitempty"`
	Used                 uint64               `protobuf:"varint,6,opt,name=Used,proto3" json:"Used,omitempty"`
	UsedPercent          float64              `protobuf:"fixed64,7,opt,name=UsedPercent,proto3" json:"UsedPercent,omitempty"`
	InodesTotal          uint64               `protobuf:"varint,8,opt,name=InodesTotal,proto3" json:"InodesTotal,omitempty"`
	InodesUsed           uint64               `protobuf:"varint,9,opt,name=InodesUsed,proto3" json:"InodesUsed,omitempty"`
	InodesFree           uint64               `protobuf:"varint,10,opt,name=InodesFree,proto3" json:"InodesFree,omitempty"`
	InodesUsedPercent    float64              `protobuf:"fixed64,11,opt,name=InodesUsedPercent,proto3" json:"InodesUsedPercent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UsageStat) Reset()         { *m = UsageStat{} }
func (m *UsageStat) String() string { return proto.CompactTextString(m) }
func (*UsageStat) ProtoMessage()    {}
func (*UsageStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_87be47849fd92a15, []int{0}
}

func (m *UsageStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsageStat.Unmarshal(m, b)
}
func (m *UsageStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsageStat.Marshal(b, m, deterministic)
}
func (m *UsageStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsageStat.Merge(m, src)
}
func (m *UsageStat) XXX_Size() int {
	return xxx_messageInfo_UsageStat.Size(m)
}
func (m *UsageStat) XXX_DiscardUnknown() {
	xxx_messageInfo_UsageStat.DiscardUnknown(m)
}

var xxx_messageInfo_UsageStat proto.InternalMessageInfo

func (m *UsageStat) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *UsageStat) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *UsageStat) GetFstype() string {
	if m != nil {
		return m.Fstype
	}
	return ""
}

func (m *UsageStat) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *UsageStat) GetFree() uint64 {
	if m != nil {
		return m.Free
	}
	return 0
}

func (m *UsageStat) GetUsed() uint64 {
	if m != nil {
		return m.Used
	}
	return 0
}

func (m *UsageStat) GetUsedPercent() float64 {
	if m != nil {
		return m.UsedPercent
	}
	return 0
}

func (m *UsageStat) GetInodesTotal() uint64 {
	if m != nil {
		return m.InodesTotal
	}
	return 0
}

func (m *UsageStat) GetInodesUsed() uint64 {
	if m != nil {
		return m.InodesUsed
	}
	return 0
}

func (m *UsageStat) GetInodesFree() uint64 {
	if m != nil {
		return m.InodesFree
	}
	return 0
}

func (m *UsageStat) GetInodesUsedPercent() float64 {
	if m != nil {
		return m.InodesUsedPercent
	}
	return 0
}

type PartitionStat struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Device               string               `protobuf:"bytes,2,opt,name=Device,proto3" json:"Device,omitempty"`
	Mountpoint           string               `protobuf:"bytes,3,opt,name=Mountpoint,proto3" json:"Mountpoint,omitempty"`
	Fstype               string               `protobuf:"bytes,4,opt,name=Fstype,proto3" json:"Fstype,omitempty"`
	Opts                 string               `protobuf:"bytes,5,opt,name=Opts,proto3" json:"Opts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PartitionStat) Reset()         { *m = PartitionStat{} }
func (m *PartitionStat) String() string { return proto.CompactTextString(m) }
func (*PartitionStat) ProtoMessage()    {}
func (*PartitionStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_87be47849fd92a15, []int{1}
}

func (m *PartitionStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionStat.Unmarshal(m, b)
}
func (m *PartitionStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionStat.Marshal(b, m, deterministic)
}
func (m *PartitionStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionStat.Merge(m, src)
}
func (m *PartitionStat) XXX_Size() int {
	return xxx_messageInfo_PartitionStat.Size(m)
}
func (m *PartitionStat) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionStat.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionStat proto.InternalMessageInfo

func (m *PartitionStat) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *PartitionStat) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *PartitionStat) GetMountpoint() string {
	if m != nil {
		return m.Mountpoint
	}
	return ""
}

func (m *PartitionStat) GetFstype() string {
	if m != nil {
		return m.Fstype
	}
	return ""
}

func (m *PartitionStat) GetOpts() string {
	if m != nil {
		return m.Opts
	}
	return ""
}

type IOCountersStat struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ReadCount            uint64               `protobuf:"varint,2,opt,name=ReadCount,proto3" json:"ReadCount,omitempty"`
	MergedReadCount      uint64               `protobuf:"varint,3,opt,name=MergedReadCount,proto3" json:"MergedReadCount,omitempty"`
	WriteCount           uint64               `protobuf:"varint,4,opt,name=WriteCount,proto3" json:"WriteCount,omitempty"`
	MergedWriteCount     uint64               `protobuf:"varint,5,opt,name=MergedWriteCount,proto3" json:"MergedWriteCount,omitempty"`
	ReadBytes            uint64               `protobuf:"varint,6,opt,name=ReadBytes,proto3" json:"ReadBytes,omitempty"`
	WriteBytes           uint64               `protobuf:"varint,7,opt,name=WriteBytes,proto3" json:"WriteBytes,omitempty"`
	ReadTime             uint64               `protobuf:"varint,8,opt,name=ReadTime,proto3" json:"ReadTime,omitempty"`
	WriteTime            uint64               `protobuf:"varint,9,opt,name=WriteTime,proto3" json:"WriteTime,omitempty"`
	IopsInProgress       uint64               `protobuf:"varint,10,opt,name=IopsInProgress,proto3" json:"IopsInProgress,omitempty"`
	IoTime               uint64               `protobuf:"varint,11,opt,name=IoTime,proto3" json:"IoTime,omitempty"`
	WeightedIO           uint64               `protobuf:"varint,12,opt,name=WeightedIO,proto3" json:"WeightedIO,omitempty"`
	Name                 string               `protobuf:"bytes,13,opt,name=Name,proto3" json:"Name,omitempty"`
	SerialNumber         string               `protobuf:"bytes,14,opt,name=SerialNumber,proto3" json:"SerialNumber,omitempty"`
	Label                string               `protobuf:"bytes,15,opt,name=Label,proto3" json:"Label,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *IOCountersStat) Reset()         { *m = IOCountersStat{} }
func (m *IOCountersStat) String() string { return proto.CompactTextString(m) }
func (*IOCountersStat) ProtoMessage()    {}
func (*IOCountersStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_87be47849fd92a15, []int{2}
}

func (m *IOCountersStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IOCountersStat.Unmarshal(m, b)
}
func (m *IOCountersStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IOCountersStat.Marshal(b, m, deterministic)
}
func (m *IOCountersStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IOCountersStat.Merge(m, src)
}
func (m *IOCountersStat) XXX_Size() int {
	return xxx_messageInfo_IOCountersStat.Size(m)
}
func (m *IOCountersStat) XXX_DiscardUnknown() {
	xxx_messageInfo_IOCountersStat.DiscardUnknown(m)
}

var xxx_messageInfo_IOCountersStat proto.InternalMessageInfo

func (m *IOCountersStat) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *IOCountersStat) GetReadCount() uint64 {
	if m != nil {
		return m.ReadCount
	}
	return 0
}

func (m *IOCountersStat) GetMergedReadCount() uint64 {
	if m != nil {
		return m.MergedReadCount
	}
	return 0
}

func (m *IOCountersStat) GetWriteCount() uint64 {
	if m != nil {
		return m.WriteCount
	}
	return 0
}

func (m *IOCountersStat) GetMergedWriteCount() uint64 {
	if m != nil {
		return m.MergedWriteCount
	}
	return 0
}

func (m *IOCountersStat) GetReadBytes() uint64 {
	if m != nil {
		return m.ReadBytes
	}
	return 0
}

func (m *IOCountersStat) GetWriteBytes() uint64 {
	if m != nil {
		return m.WriteBytes
	}
	return 0
}

func (m *IOCountersStat) GetReadTime() uint64 {
	if m != nil {
		return m.ReadTime
	}
	return 0
}

func (m *IOCountersStat) GetWriteTime() uint64 {
	if m != nil {
		return m.WriteTime
	}
	return 0
}

func (m *IOCountersStat) GetIopsInProgress() uint64 {
	if m != nil {
		return m.IopsInProgress
	}
	return 0
}

func (m *IOCountersStat) GetIoTime() uint64 {
	if m != nil {
		return m.IoTime
	}
	return 0
}

func (m *IOCountersStat) GetWeightedIO() uint64 {
	if m != nil {
		return m.WeightedIO
	}
	return 0
}

func (m *IOCountersStat) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IOCountersStat) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *IOCountersStat) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type DiskRequest struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	IP                   string               `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	NodeName             string               `protobuf:"bytes,3,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	UsageStat            []*UsageStat         `protobuf:"bytes,4,rep,name=usageStat,proto3" json:"usageStat,omitempty"`
	PartitionStat        []*PartitionStat     `protobuf:"bytes,5,rep,name=partitionStat,proto3" json:"partitionStat,omitempty"`
	IoCountersStat       []*IOCountersStat    `protobuf:"bytes,6,rep,name=ioCountersStat,proto3" json:"ioCountersStat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DiskRequest) Reset()         { *m = DiskRequest{} }
func (m *DiskRequest) String() string { return proto.CompactTextString(m) }
func (*DiskRequest) ProtoMessage()    {}
func (*DiskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_87be47849fd92a15, []int{3}
}

func (m *DiskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiskRequest.Unmarshal(m, b)
}
func (m *DiskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiskRequest.Marshal(b, m, deterministic)
}
func (m *DiskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiskRequest.Merge(m, src)
}
func (m *DiskRequest) XXX_Size() int {
	return xxx_messageInfo_DiskRequest.Size(m)
}
func (m *DiskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiskRequest proto.InternalMessageInfo

func (m *DiskRequest) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *DiskRequest) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *DiskRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *DiskRequest) GetUsageStat() []*UsageStat {
	if m != nil {
		return m.UsageStat
	}
	return nil
}

func (m *DiskRequest) GetPartitionStat() []*PartitionStat {
	if m != nil {
		return m.PartitionStat
	}
	return nil
}

func (m *DiskRequest) GetIoCountersStat() []*IOCountersStat {
	if m != nil {
		return m.IoCountersStat
	}
	return nil
}

type DiskResponse struct {
	Error                *basic.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DiskResponse) Reset()         { *m = DiskResponse{} }
func (m *DiskResponse) String() string { return proto.CompactTextString(m) }
func (*DiskResponse) ProtoMessage()    {}
func (*DiskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_87be47849fd92a15, []int{4}
}

func (m *DiskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiskResponse.Unmarshal(m, b)
}
func (m *DiskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiskResponse.Marshal(b, m, deterministic)
}
func (m *DiskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiskResponse.Merge(m, src)
}
func (m *DiskResponse) XXX_Size() int {
	return xxx_messageInfo_DiskResponse.Size(m)
}
func (m *DiskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DiskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DiskResponse proto.InternalMessageInfo

func (m *DiskResponse) GetError() *basic.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*UsageStat)(nil), "protobuf.pb.disk.UsageStat")
	proto.RegisterType((*PartitionStat)(nil), "protobuf.pb.disk.PartitionStat")
	proto.RegisterType((*IOCountersStat)(nil), "protobuf.pb.disk.IOCountersStat")
	proto.RegisterType((*DiskRequest)(nil), "protobuf.pb.disk.DiskRequest")
	proto.RegisterType((*DiskResponse)(nil), "protobuf.pb.disk.DiskResponse")
}

func init() { proto.RegisterFile("disk/disk.proto", fileDescriptor_87be47849fd92a15) }

var fileDescriptor_87be47849fd92a15 = []byte{
	// 737 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xd1, 0x4e, 0xdb, 0x4a,
	0x10, 0xbd, 0x09, 0x4e, 0x20, 0x13, 0x08, 0xb0, 0xba, 0x17, 0x59, 0xb9, 0xf7, 0x42, 0x94, 0x87,
	0x2a, 0x6a, 0x1b, 0x47, 0xa2, 0x2f, 0xad, 0xfa, 0x46, 0x01, 0xd5, 0x52, 0x81, 0xc8, 0x80, 0x2a,
	0xf5, 0xcd, 0x4e, 0x06, 0x67, 0x45, 0xec, 0x75, 0x77, 0x37, 0xad, 0x78, 0xef, 0x07, 0xf4, 0x47,
	0xfa, 0x0d, 0xfd, 0x96, 0xfe, 0x49, 0xb5, 0xb3, 0x4e, 0xec, 0x24, 0xea, 0x53, 0x5e, 0xec, 0x9d,
	0x33, 0x67, 0x66, 0x8f, 0x67, 0x67, 0xc7, 0xb0, 0x3f, 0xe6, 0xea, 0x71, 0x60, 0x1e, 0x5e, 0x26,
	0x85, 0x16, 0xec, 0x80, 0x5e, 0xd1, 0xec, 0xc1, 0xcb, 0x22, 0xcf, 0xe0, 0xed, 0xc3, 0x28, 0x54,
	0x7c, 0x34, 0x40, 0x29, 0x85, 0xb4, 0xa4, 0xf6, 0x49, 0x2c, 0x44, 0x3c, 0xc5, 0xc1, 0x9c, 0x3b,
	0xd0, 0x3c, 0x41, 0xa5, 0xc3, 0x24, 0xb3, 0x84, 0xee, 0xaf, 0x2a, 0x34, 0xee, 0x55, 0x18, 0xe3,
	0xad, 0x0e, 0x35, 0x7b, 0x0d, 0x8d, 0x05, 0xc1, 0xad, 0x74, 0x2a, 0xbd, 0xe6, 0x69, 0xdb, 0xb3,
	0x29, 0xbc, 0xc5, 0x76, 0x77, 0x73, 0x46, 0x50, 0x90, 0x19, 0x03, 0x67, 0x18, 0xea, 0x89, 0x5b,
	0xed, 0x54, 0x7a, 0x8d, 0x80, 0xd6, 0xec, 0x08, 0xea, 0x97, 0x4a, 0x3f, 0x65, 0xe8, 0x6e, 0x11,
	0x9a, 0x5b, 0xec, 0x6f, 0xa8, 0xdd, 0x09, 0x1d, 0x4e, 0x5d, 0xa7, 0x53, 0xe9, 0x39, 0x81, 0x35,
	0x4c, 0x86, 0x4b, 0x89, 0xe8, 0xd6, 0x08, 0xa4, 0xb5, 0xc1, 0xee, 0x15, 0x8e, 0xdd, 0xba, 0xc5,
	0xcc, 0x9a, 0x75, 0xa0, 0x69, 0xde, 0x43, 0x94, 0x23, 0x4c, 0xb5, 0xbb, 0xdd, 0xa9, 0xf4, 0x2a,
	0x41, 0x19, 0x32, 0x0c, 0x3f, 0x15, 0x63, 0x54, 0x76, 0x97, 0x1d, 0x0a, 0x2e, 0x43, 0xec, 0x18,
	0xc0, 0x9a, 0x94, 0xbd, 0x41, 0x84, 0x12, 0x52, 0xf8, 0x49, 0x11, 0x94, 0xfd, 0xa4, 0xeb, 0x25,
	0x1c, 0x16, 0xec, 0xb9, 0x92, 0x26, 0x29, 0x59, 0x77, 0x74, 0x7f, 0x54, 0x60, 0x6f, 0x18, 0x4a,
	0xcd, 0x35, 0x17, 0xe9, 0x86, 0x75, 0x3e, 0x82, 0xfa, 0x39, 0x7e, 0xe1, 0x23, 0xcc, 0x2b, 0x9d,
	0x5b, 0x46, 0xf1, 0x95, 0x98, 0xa5, 0x3a, 0x13, 0x3c, 0xd5, 0x79, 0xbd, 0x4b, 0x48, 0xe9, 0x2c,
	0x9c, 0xa5, 0xb3, 0x60, 0xe0, 0xdc, 0x64, 0x5a, 0x51, 0xd5, 0x1b, 0x01, 0xad, 0xbb, 0xdf, 0x1c,
	0x68, 0xf9, 0x37, 0xef, 0x4c, 0x30, 0x4a, 0xb5, 0xa1, 0xe0, 0xff, 0xa0, 0x11, 0x60, 0x38, 0xa6,
	0x6c, 0xa4, 0xd9, 0x09, 0x0a, 0x80, 0xf5, 0x60, 0xff, 0x0a, 0x65, 0x8c, 0xe3, 0x82, 0xb3, 0x45,
	0x9c, 0x55, 0xd8, 0x7c, 0xe0, 0x47, 0xc9, 0x35, 0x5a, 0x92, 0xed, 0x9c, 0x12, 0xc2, 0x9e, 0xc3,
	0x81, 0x0d, 0x29, 0xb1, 0x6c, 0x2b, 0xad, 0xe1, 0x73, 0x4d, 0x67, 0x4f, 0x1a, 0x55, 0xde, 0x5b,
	0x05, 0xb0, 0xd8, 0xc9, 0xba, 0xb7, 0x4b, 0x3b, 0x59, 0x7f, 0x1b, 0x76, 0x0c, 0xd9, 0x7c, 0x6d,
	0xde, 0x5b, 0x0b, 0xdb, 0x64, 0x26, 0x26, 0x39, 0x6d, 0x5f, 0x15, 0x00, 0x7b, 0x06, 0x2d, 0x5f,
	0x64, 0xca, 0x4f, 0x87, 0x52, 0xc4, 0x12, 0x95, 0xca, 0x5b, 0x6b, 0x05, 0x35, 0x87, 0xe5, 0x0b,
	0x4a, 0xd1, 0x24, 0x7f, 0x6e, 0x91, 0x32, 0xe4, 0xf1, 0x44, 0xe3, 0xd8, 0xbf, 0x71, 0x77, 0x73,
	0x65, 0x0b, 0xc4, 0x1c, 0xe6, 0x75, 0x98, 0xa0, 0xbb, 0x67, 0x0f, 0xd3, 0xac, 0x59, 0x17, 0x76,
	0x6f, 0x51, 0xf2, 0x70, 0x7a, 0x3d, 0x4b, 0x22, 0x94, 0x6e, 0x8b, 0x7c, 0x4b, 0x98, 0xb9, 0x90,
	0x1f, 0xc2, 0x08, 0xa7, 0xee, 0x3e, 0x39, 0xad, 0xd1, 0xfd, 0x59, 0x85, 0xe6, 0x39, 0x57, 0x8f,
	0x01, 0x7e, 0x9e, 0xa1, 0xda, 0xa4, 0x07, 0x5a, 0x50, 0xf5, 0x87, 0x79, 0xc3, 0x56, 0xfd, 0xa1,
	0xa9, 0xa0, 0xb9, 0x24, 0xa4, 0xd5, 0xb6, 0xea, 0xc2, 0x66, 0x6f, 0xa0, 0x31, 0x9b, 0xcf, 0x23,
	0xd7, 0xe9, 0x6c, 0xf5, 0x9a, 0xa7, 0xff, 0x7a, 0xab, 0xa3, 0xce, 0x5b, 0x8c, 0xac, 0xa0, 0x60,
	0xb3, 0x0b, 0xd8, 0xcb, 0xca, 0xd7, 0xcc, 0xad, 0x51, 0xf8, 0xc9, 0x7a, 0xf8, 0xd2, 0x6d, 0x0c,
	0x96, 0xa3, 0xd8, 0x7b, 0x68, 0x71, 0x51, 0xee, 0x7e, 0xb7, 0x4e, 0x79, 0x3a, 0xeb, 0x79, 0x96,
	0x6f, 0x49, 0xb0, 0x12, 0xd7, 0x7d, 0x0b, 0xbb, 0xb6, 0x80, 0x2a, 0x13, 0xa9, 0x42, 0xf6, 0x02,
	0x6a, 0x34, 0x9c, 0xf3, 0xea, 0xfd, 0x53, 0x24, 0xa4, 0xc9, 0xed, 0x5d, 0x18, 0x67, 0x60, 0x39,
	0xa7, 0xdf, 0xf3, 0xf2, 0xdf, 0xa2, 0xa4, 0x1b, 0x7e, 0x07, 0x87, 0xc3, 0x99, 0x9a, 0x18, 0xa8,
	0x18, 0xd8, 0xff, 0xaf, 0x6b, 0x2a, 0x1d, 0x59, 0xfb, 0xf8, 0x4f, 0x6e, 0x2b, 0xa8, 0xfb, 0xd7,
	0x3c, 0xeb, 0xf2, 0x78, 0xda, 0x38, 0xeb, 0x3d, 0x30, 0x93, 0x75, 0x65, 0x88, 0x6c, 0x9a, 0xf6,
	0xec, 0xfc, 0xd3, 0x59, 0xcc, 0xf5, 0x64, 0x16, 0x79, 0x23, 0x91, 0x0c, 0x12, 0x3e, 0x92, 0xa2,
	0xcf, 0xd3, 0xfe, 0x28, 0x1d, 0x64, 0xd3, 0x50, 0x3f, 0x08, 0x99, 0xf4, 0xbf, 0x62, 0x34, 0x08,
	0x95, 0xc2, 0x24, 0x9a, 0x3e, 0xf5, 0xa7, 0x3c, 0x2d, 0xfd, 0xfa, 0x62, 0x41, 0xbf, 0xcf, 0xa8,
	0x4e, 0xc8, 0xab, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xba, 0x78, 0x57, 0x20, 0x52, 0x07, 0x00,
	0x00,
}
