// Code generated by protoc-gen-go. DO NOT EDIT.
// source: net/net.proto

package net

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

type IOCountersStat struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	BytesSent            uint64               `protobuf:"varint,3,opt,name=BytesSent,proto3" json:"BytesSent,omitempty"`
	BytesRecv            uint64               `protobuf:"varint,4,opt,name=BytesRecv,proto3" json:"BytesRecv,omitempty"`
	PacketsSent          uint64               `protobuf:"varint,5,opt,name=PacketsSent,proto3" json:"PacketsSent,omitempty"`
	PacketsRecv          uint64               `protobuf:"varint,6,opt,name=PacketsRecv,proto3" json:"PacketsRecv,omitempty"`
	Errin                uint64               `protobuf:"varint,7,opt,name=Errin,proto3" json:"Errin,omitempty"`
	Errout               uint64               `protobuf:"varint,8,opt,name=Errout,proto3" json:"Errout,omitempty"`
	Dropin               uint64               `protobuf:"varint,9,opt,name=Dropin,proto3" json:"Dropin,omitempty"`
	Dropout              uint64               `protobuf:"varint,10,opt,name=Dropout,proto3" json:"Dropout,omitempty"`
	Fifoin               uint64               `protobuf:"varint,11,opt,name=Fifoin,proto3" json:"Fifoin,omitempty"`
	Fifoout              uint64               `protobuf:"varint,12,opt,name=Fifoout,proto3" json:"Fifoout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *IOCountersStat) Reset()         { *m = IOCountersStat{} }
func (m *IOCountersStat) String() string { return proto.CompactTextString(m) }
func (*IOCountersStat) ProtoMessage()    {}
func (*IOCountersStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b79b5a127a76ba1, []int{0}
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

func (m *IOCountersStat) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IOCountersStat) GetBytesSent() uint64 {
	if m != nil {
		return m.BytesSent
	}
	return 0
}

func (m *IOCountersStat) GetBytesRecv() uint64 {
	if m != nil {
		return m.BytesRecv
	}
	return 0
}

func (m *IOCountersStat) GetPacketsSent() uint64 {
	if m != nil {
		return m.PacketsSent
	}
	return 0
}

func (m *IOCountersStat) GetPacketsRecv() uint64 {
	if m != nil {
		return m.PacketsRecv
	}
	return 0
}

func (m *IOCountersStat) GetErrin() uint64 {
	if m != nil {
		return m.Errin
	}
	return 0
}

func (m *IOCountersStat) GetErrout() uint64 {
	if m != nil {
		return m.Errout
	}
	return 0
}

func (m *IOCountersStat) GetDropin() uint64 {
	if m != nil {
		return m.Dropin
	}
	return 0
}

func (m *IOCountersStat) GetDropout() uint64 {
	if m != nil {
		return m.Dropout
	}
	return 0
}

func (m *IOCountersStat) GetFifoin() uint64 {
	if m != nil {
		return m.Fifoin
	}
	return 0
}

func (m *IOCountersStat) GetFifoout() uint64 {
	if m != nil {
		return m.Fifoout
	}
	return 0
}

// coming soon
type ProtoCountersStat struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ProtoCountersStat) Reset()         { *m = ProtoCountersStat{} }
func (m *ProtoCountersStat) String() string { return proto.CompactTextString(m) }
func (*ProtoCountersStat) ProtoMessage()    {}
func (*ProtoCountersStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b79b5a127a76ba1, []int{1}
}

func (m *ProtoCountersStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoCountersStat.Unmarshal(m, b)
}
func (m *ProtoCountersStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoCountersStat.Marshal(b, m, deterministic)
}
func (m *ProtoCountersStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoCountersStat.Merge(m, src)
}
func (m *ProtoCountersStat) XXX_Size() int {
	return xxx_messageInfo_ProtoCountersStat.Size(m)
}
func (m *ProtoCountersStat) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoCountersStat.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoCountersStat proto.InternalMessageInfo

func (m *ProtoCountersStat) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type Addr struct {
	IP                   string   `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"`
	Port                 uint32   `protobuf:"varint,2,opt,name=Port,proto3" json:"Port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Addr) Reset()         { *m = Addr{} }
func (m *Addr) String() string { return proto.CompactTextString(m) }
func (*Addr) ProtoMessage()    {}
func (*Addr) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b79b5a127a76ba1, []int{2}
}

func (m *Addr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Addr.Unmarshal(m, b)
}
func (m *Addr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Addr.Marshal(b, m, deterministic)
}
func (m *Addr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Addr.Merge(m, src)
}
func (m *Addr) XXX_Size() int {
	return xxx_messageInfo_Addr.Size(m)
}
func (m *Addr) XXX_DiscardUnknown() {
	xxx_messageInfo_Addr.DiscardUnknown(m)
}

var xxx_messageInfo_Addr proto.InternalMessageInfo

func (m *Addr) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *Addr) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type ConnectionStat struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Fd                   uint32               `protobuf:"varint,2,opt,name=Fd,proto3" json:"Fd,omitempty"`
	Family               uint32               `protobuf:"varint,3,opt,name=Family,proto3" json:"Family,omitempty"`
	Type                 uint32               `protobuf:"varint,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Laddr                *Addr                `protobuf:"bytes,5,opt,name=Laddr,proto3" json:"Laddr,omitempty"`
	Raddr                *Addr                `protobuf:"bytes,6,opt,name=Raddr,proto3" json:"Raddr,omitempty"`
	Status               string               `protobuf:"bytes,7,opt,name=Status,proto3" json:"Status,omitempty"`
	Uids                 []int32              `protobuf:"varint,8,rep,packed,name=Uids,proto3" json:"Uids,omitempty"`
	Pid                  int32                `protobuf:"varint,9,opt,name=Pid,proto3" json:"Pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ConnectionStat) Reset()         { *m = ConnectionStat{} }
func (m *ConnectionStat) String() string { return proto.CompactTextString(m) }
func (*ConnectionStat) ProtoMessage()    {}
func (*ConnectionStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b79b5a127a76ba1, []int{3}
}

func (m *ConnectionStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionStat.Unmarshal(m, b)
}
func (m *ConnectionStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionStat.Marshal(b, m, deterministic)
}
func (m *ConnectionStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionStat.Merge(m, src)
}
func (m *ConnectionStat) XXX_Size() int {
	return xxx_messageInfo_ConnectionStat.Size(m)
}
func (m *ConnectionStat) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionStat.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionStat proto.InternalMessageInfo

func (m *ConnectionStat) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ConnectionStat) GetFd() uint32 {
	if m != nil {
		return m.Fd
	}
	return 0
}

func (m *ConnectionStat) GetFamily() uint32 {
	if m != nil {
		return m.Family
	}
	return 0
}

func (m *ConnectionStat) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ConnectionStat) GetLaddr() *Addr {
	if m != nil {
		return m.Laddr
	}
	return nil
}

func (m *ConnectionStat) GetRaddr() *Addr {
	if m != nil {
		return m.Raddr
	}
	return nil
}

func (m *ConnectionStat) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ConnectionStat) GetUids() []int32 {
	if m != nil {
		return m.Uids
	}
	return nil
}

func (m *ConnectionStat) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

// coming soon
type FilterStat struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FilterStat) Reset()         { *m = FilterStat{} }
func (m *FilterStat) String() string { return proto.CompactTextString(m) }
func (*FilterStat) ProtoMessage()    {}
func (*FilterStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b79b5a127a76ba1, []int{4}
}

func (m *FilterStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterStat.Unmarshal(m, b)
}
func (m *FilterStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterStat.Marshal(b, m, deterministic)
}
func (m *FilterStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterStat.Merge(m, src)
}
func (m *FilterStat) XXX_Size() int {
	return xxx_messageInfo_FilterStat.Size(m)
}
func (m *FilterStat) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterStat.DiscardUnknown(m)
}

var xxx_messageInfo_FilterStat proto.InternalMessageInfo

func (m *FilterStat) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type NetRequest struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	IP                   string               `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	NodeName             string               `protobuf:"bytes,3,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	IOCountersStat       []*IOCountersStat    `protobuf:"bytes,4,rep,name=iOCountersStat,proto3" json:"iOCountersStat,omitempty"`
	ConnectionStat       []*ConnectionStat    `protobuf:"bytes,5,rep,name=connectionStat,proto3" json:"connectionStat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *NetRequest) Reset()         { *m = NetRequest{} }
func (m *NetRequest) String() string { return proto.CompactTextString(m) }
func (*NetRequest) ProtoMessage()    {}
func (*NetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b79b5a127a76ba1, []int{5}
}

func (m *NetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetRequest.Unmarshal(m, b)
}
func (m *NetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetRequest.Marshal(b, m, deterministic)
}
func (m *NetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetRequest.Merge(m, src)
}
func (m *NetRequest) XXX_Size() int {
	return xxx_messageInfo_NetRequest.Size(m)
}
func (m *NetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetRequest proto.InternalMessageInfo

func (m *NetRequest) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *NetRequest) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *NetRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *NetRequest) GetIOCountersStat() []*IOCountersStat {
	if m != nil {
		return m.IOCountersStat
	}
	return nil
}

func (m *NetRequest) GetConnectionStat() []*ConnectionStat {
	if m != nil {
		return m.ConnectionStat
	}
	return nil
}

type NetResponse struct {
	Error                *basic.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *NetResponse) Reset()         { *m = NetResponse{} }
func (m *NetResponse) String() string { return proto.CompactTextString(m) }
func (*NetResponse) ProtoMessage()    {}
func (*NetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b79b5a127a76ba1, []int{6}
}

func (m *NetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetResponse.Unmarshal(m, b)
}
func (m *NetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetResponse.Marshal(b, m, deterministic)
}
func (m *NetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetResponse.Merge(m, src)
}
func (m *NetResponse) XXX_Size() int {
	return xxx_messageInfo_NetResponse.Size(m)
}
func (m *NetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NetResponse proto.InternalMessageInfo

func (m *NetResponse) GetError() *basic.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*IOCountersStat)(nil), "protobuf.pb.net.IOCountersStat")
	proto.RegisterType((*ProtoCountersStat)(nil), "protobuf.pb.net.ProtoCountersStat")
	proto.RegisterType((*Addr)(nil), "protobuf.pb.net.Addr")
	proto.RegisterType((*ConnectionStat)(nil), "protobuf.pb.net.ConnectionStat")
	proto.RegisterType((*FilterStat)(nil), "protobuf.pb.net.FilterStat")
	proto.RegisterType((*NetRequest)(nil), "protobuf.pb.net.NetRequest")
	proto.RegisterType((*NetResponse)(nil), "protobuf.pb.net.NetResponse")
}

func init() { proto.RegisterFile("net/net.proto", fileDescriptor_8b79b5a127a76ba1) }

var fileDescriptor_8b79b5a127a76ba1 = []byte{
	// 626 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x7e, 0xe3, 0x7c, 0xb4, 0x99, 0xbc, 0x09, 0x74, 0x05, 0xc8, 0x0a, 0x95, 0x1a, 0xf9, 0x14,
	0x81, 0x62, 0x4b, 0xe1, 0x82, 0xb8, 0xb5, 0xa5, 0x41, 0x91, 0xa0, 0x98, 0x6d, 0xb9, 0x70, 0xf3,
	0xc7, 0x36, 0x5d, 0x11, 0xef, 0x9a, 0xf5, 0xba, 0x28, 0x77, 0x7e, 0x07, 0x3f, 0x82, 0x5f, 0xc6,
	0x4f, 0x40, 0x3b, 0xeb, 0x7c, 0x16, 0xb8, 0x84, 0x53, 0xe6, 0xe3, 0x79, 0x26, 0xe3, 0x79, 0x76,
	0x06, 0xba, 0x82, 0xe9, 0x40, 0x30, 0xed, 0xe7, 0x4a, 0x6a, 0x49, 0x1e, 0xe0, 0x4f, 0x5c, 0xde,
	0xf8, 0x79, 0xec, 0x0b, 0xa6, 0xfb, 0x47, 0x71, 0x54, 0xf0, 0x24, 0x60, 0x4a, 0x49, 0x65, 0x31,
	0xfd, 0x93, 0x99, 0x94, 0xb3, 0x39, 0x0b, 0x96, 0xd0, 0x40, 0xf3, 0x8c, 0x15, 0x3a, 0xca, 0x72,
	0x0b, 0xf0, 0x7e, 0x3a, 0xd0, 0x9b, 0xbe, 0x3f, 0x97, 0xa5, 0xd0, 0x4c, 0x15, 0x57, 0x3a, 0xd2,
	0xe4, 0x25, 0xb4, 0x57, 0x28, 0xb7, 0x36, 0xa8, 0x0d, 0x3b, 0xe3, 0xbe, 0x6f, 0xeb, 0xf8, 0xab,
	0xbf, 0xbc, 0x5e, 0x22, 0xe8, 0x1a, 0x4c, 0x08, 0x34, 0x2e, 0xa3, 0x8c, 0xb9, 0xce, 0xa0, 0x36,
	0x6c, 0x53, 0xb4, 0xc9, 0x31, 0xb4, 0xcf, 0x16, 0x9a, 0x15, 0x57, 0x4c, 0x68, 0xb7, 0x3e, 0xa8,
	0x0d, 0x1b, 0x74, 0x1d, 0x58, 0x65, 0x29, 0x4b, 0xee, 0xdc, 0xc6, 0x46, 0xd6, 0x04, 0xc8, 0x00,
	0x3a, 0x61, 0x94, 0x7c, 0x66, 0xda, 0xb2, 0x9b, 0x98, 0xdf, 0x0c, 0x6d, 0x20, 0xb0, 0x42, 0x6b,
	0x0b, 0x81, 0x35, 0x1e, 0x41, 0xf3, 0x42, 0x29, 0x2e, 0xdc, 0x03, 0xcc, 0x59, 0x87, 0x3c, 0x81,
	0xd6, 0x85, 0x52, 0xb2, 0xd4, 0xee, 0x21, 0x86, 0x2b, 0xcf, 0xc4, 0x5f, 0x2b, 0x99, 0x73, 0xe1,
	0xb6, 0x6d, 0xdc, 0x7a, 0xc4, 0x85, 0x03, 0x63, 0x19, 0x02, 0x60, 0x62, 0xe9, 0x1a, 0xc6, 0x84,
	0xdf, 0x48, 0x2e, 0xdc, 0x8e, 0x65, 0x58, 0xcf, 0x30, 0x8c, 0x65, 0x18, 0xff, 0x5b, 0x46, 0xe5,
	0x7a, 0xef, 0xe0, 0x28, 0x34, 0x63, 0xfc, 0x37, 0x43, 0xf7, 0x9e, 0x41, 0xe3, 0x34, 0x4d, 0x15,
	0xe9, 0x81, 0x33, 0x0d, 0x91, 0xda, 0xa6, 0xce, 0x34, 0x34, 0x62, 0x84, 0x52, 0x69, 0x14, 0xa3,
	0x4b, 0xd1, 0xf6, 0xbe, 0x3b, 0xd0, 0x3b, 0x97, 0x42, 0xb0, 0x44, 0x73, 0x29, 0xf6, 0x54, 0xbb,
	0x07, 0xce, 0x24, 0xad, 0xca, 0x3b, 0x93, 0x14, 0x27, 0x11, 0x65, 0x7c, 0xbe, 0x40, 0x99, 0xbb,
	0xb4, 0xf2, 0x4c, 0x23, 0xd7, 0x8b, 0x9c, 0xa1, 0xbc, 0x5d, 0x8a, 0x36, 0x79, 0x0e, 0xcd, 0xb7,
	0x51, 0x9a, 0x2a, 0xd4, 0xb4, 0x33, 0x7e, 0xec, 0xef, 0xbc, 0x65, 0xdf, 0x7c, 0x12, 0xb5, 0x18,
	0x03, 0xa6, 0x08, 0x6e, 0xfd, 0x15, 0x8c, 0x18, 0xd3, 0x85, 0xf9, 0xae, 0xb2, 0x40, 0xc1, 0xdb,
	0xb4, 0xf2, 0x4c, 0x17, 0x1f, 0x79, 0x5a, 0xb8, 0x87, 0x83, 0xfa, 0xb0, 0x49, 0xd1, 0x26, 0x0f,
	0xa1, 0x1e, 0xf2, 0x14, 0xa5, 0x6e, 0x52, 0x63, 0x7a, 0x13, 0x80, 0x09, 0x9f, 0x6b, 0xa6, 0xf6,
	0x14, 0xe5, 0x9b, 0x03, 0x70, 0xc9, 0x34, 0x65, 0x5f, 0x4a, 0x56, 0xec, 0x39, 0xe4, 0x69, 0x58,
	0x2d, 0x94, 0x51, 0xb5, 0x0f, 0x87, 0x42, 0xa6, 0x0c, 0xd7, 0xac, 0x8e, 0xd1, 0x95, 0x4f, 0xde,
	0x40, 0x8f, 0x6f, 0xad, 0xb2, 0xdb, 0x18, 0xd4, 0x87, 0x9d, 0xf1, 0xc9, 0xbd, 0x81, 0x6d, 0x6f,
	0x3c, 0xdd, 0xa1, 0x99, 0x42, 0xc9, 0xd6, 0x2b, 0x71, 0x9b, 0x7f, 0x28, 0xb4, 0xfd, 0x98, 0xe8,
	0x0e, 0xcd, 0x7b, 0x05, 0x1d, 0x9c, 0x42, 0x91, 0x4b, 0x51, 0xa0, 0xea, 0x78, 0x9c, 0xaa, 0x11,
	0x6c, 0x08, 0x89, 0x97, 0xcb, 0x37, 0x4b, 0xa8, 0xa8, 0xc5, 0x8c, 0x7f, 0xd4, 0x70, 0x84, 0x57,
	0x4c, 0xdd, 0xf1, 0x84, 0x91, 0x0f, 0x40, 0xc2, 0xb2, 0xb8, 0xdd, 0xb9, 0x55, 0x4f, 0xef, 0x75,
	0xb4, 0x9e, 0x7a, 0xff, 0xf8, 0xf7, 0x49, 0xdb, 0x8c, 0xf7, 0xdf, 0xb2, 0xe4, 0xce, 0x42, 0xec,
	0x53, 0xf2, 0xec, 0xfc, 0xd3, 0xe9, 0x8c, 0xeb, 0xdb, 0x32, 0xf6, 0x13, 0x99, 0x05, 0x19, 0x4f,
	0x94, 0x1c, 0x71, 0x31, 0x4a, 0x44, 0x90, 0xcf, 0x23, 0x7d, 0x23, 0x55, 0x36, 0xfa, 0xca, 0xe2,
	0x20, 0x2a, 0x0a, 0x96, 0xc5, 0xf3, 0xc5, 0x68, 0xce, 0xc5, 0xc6, 0x71, 0x9e, 0x49, 0x73, 0xde,
	0xe3, 0x16, 0x06, 0x5e, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x77, 0xa2, 0x4d, 0x04, 0xf0, 0x05,
	0x00, 0x00,
}
