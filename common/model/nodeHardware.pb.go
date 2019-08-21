// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/nodeHardware.proto

package model

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

type GetNodeHardwareResponse struct {
	NodeHardware         *NodeHardware `protobuf:"bytes,1,opt,name=NodeHardware,proto3" json:"NodeHardware,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetNodeHardwareResponse) Reset()         { *m = GetNodeHardwareResponse{} }
func (m *GetNodeHardwareResponse) String() string { return proto.CompactTextString(m) }
func (*GetNodeHardwareResponse) ProtoMessage()    {}
func (*GetNodeHardwareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce967b26533b171c, []int{0}
}

func (m *GetNodeHardwareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNodeHardwareResponse.Unmarshal(m, b)
}
func (m *GetNodeHardwareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNodeHardwareResponse.Marshal(b, m, deterministic)
}
func (m *GetNodeHardwareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNodeHardwareResponse.Merge(m, src)
}
func (m *GetNodeHardwareResponse) XXX_Size() int {
	return xxx_messageInfo_GetNodeHardwareResponse.Size(m)
}
func (m *GetNodeHardwareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNodeHardwareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNodeHardwareResponse proto.InternalMessageInfo

func (m *GetNodeHardwareResponse) GetNodeHardware() *NodeHardware {
	if m != nil {
		return m.NodeHardware
	}
	return nil
}

type GetNodeHardwareRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNodeHardwareRequest) Reset()         { *m = GetNodeHardwareRequest{} }
func (m *GetNodeHardwareRequest) String() string { return proto.CompactTextString(m) }
func (*GetNodeHardwareRequest) ProtoMessage()    {}
func (*GetNodeHardwareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce967b26533b171c, []int{1}
}

func (m *GetNodeHardwareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNodeHardwareRequest.Unmarshal(m, b)
}
func (m *GetNodeHardwareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNodeHardwareRequest.Marshal(b, m, deterministic)
}
func (m *GetNodeHardwareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNodeHardwareRequest.Merge(m, src)
}
func (m *GetNodeHardwareRequest) XXX_Size() int {
	return xxx_messageInfo_GetNodeHardwareRequest.Size(m)
}
func (m *GetNodeHardwareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNodeHardwareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNodeHardwareRequest proto.InternalMessageInfo

type NodeHardware struct {
	CPUInformation       []*CPUInformation   `protobuf:"bytes,1,rep,name=CPUInformation,proto3" json:"CPUInformation,omitempty"`
	MemoryInformation    *MemoryInformation  `protobuf:"bytes,2,opt,name=MemoryInformation,proto3" json:"MemoryInformation,omitempty"`
	StorageInformation   *StorageInformation `protobuf:"bytes,3,opt,name=StorageInformation,proto3" json:"StorageInformation,omitempty"`
	HostInformation      *HostInformation    `protobuf:"bytes,4,opt,name=HostInformation,proto3" json:"HostInformation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NodeHardware) Reset()         { *m = NodeHardware{} }
func (m *NodeHardware) String() string { return proto.CompactTextString(m) }
func (*NodeHardware) ProtoMessage()    {}
func (*NodeHardware) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce967b26533b171c, []int{2}
}

func (m *NodeHardware) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeHardware.Unmarshal(m, b)
}
func (m *NodeHardware) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeHardware.Marshal(b, m, deterministic)
}
func (m *NodeHardware) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeHardware.Merge(m, src)
}
func (m *NodeHardware) XXX_Size() int {
	return xxx_messageInfo_NodeHardware.Size(m)
}
func (m *NodeHardware) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeHardware.DiscardUnknown(m)
}

var xxx_messageInfo_NodeHardware proto.InternalMessageInfo

func (m *NodeHardware) GetCPUInformation() []*CPUInformation {
	if m != nil {
		return m.CPUInformation
	}
	return nil
}

func (m *NodeHardware) GetMemoryInformation() *MemoryInformation {
	if m != nil {
		return m.MemoryInformation
	}
	return nil
}

func (m *NodeHardware) GetStorageInformation() *StorageInformation {
	if m != nil {
		return m.StorageInformation
	}
	return nil
}

func (m *NodeHardware) GetHostInformation() *HostInformation {
	if m != nil {
		return m.HostInformation
	}
	return nil
}

type CPUInformation struct {
	Family               string   `protobuf:"bytes,1,opt,name=Family,proto3" json:"Family,omitempty"`
	CPUIndex             int32    `protobuf:"varint,2,opt,name=CPUIndex,proto3" json:"CPUIndex,omitempty"`
	Model                string   `protobuf:"bytes,3,opt,name=Model,proto3" json:"Model,omitempty"`
	ModelName            string   `protobuf:"bytes,4,opt,name=ModelName,proto3" json:"ModelName,omitempty"`
	VendorId             string   `protobuf:"bytes,5,opt,name=VendorId,proto3" json:"VendorId,omitempty"`
	Mhz                  float64  `protobuf:"fixed64,6,opt,name=Mhz,proto3" json:"Mhz,omitempty"`
	CacheSize            int32    `protobuf:"varint,7,opt,name=CacheSize,proto3" json:"CacheSize,omitempty"`
	UsedPercent          float64  `protobuf:"fixed64,8,opt,name=UsedPercent,proto3" json:"UsedPercent,omitempty"`
	CoreID               string   `protobuf:"bytes,9,opt,name=CoreID,proto3" json:"CoreID,omitempty"`
	Cores                int32    `protobuf:"varint,10,opt,name=Cores,proto3" json:"Cores,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPUInformation) Reset()         { *m = CPUInformation{} }
func (m *CPUInformation) String() string { return proto.CompactTextString(m) }
func (*CPUInformation) ProtoMessage()    {}
func (*CPUInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce967b26533b171c, []int{3}
}

func (m *CPUInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPUInformation.Unmarshal(m, b)
}
func (m *CPUInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPUInformation.Marshal(b, m, deterministic)
}
func (m *CPUInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPUInformation.Merge(m, src)
}
func (m *CPUInformation) XXX_Size() int {
	return xxx_messageInfo_CPUInformation.Size(m)
}
func (m *CPUInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_CPUInformation.DiscardUnknown(m)
}

var xxx_messageInfo_CPUInformation proto.InternalMessageInfo

func (m *CPUInformation) GetFamily() string {
	if m != nil {
		return m.Family
	}
	return ""
}

func (m *CPUInformation) GetCPUIndex() int32 {
	if m != nil {
		return m.CPUIndex
	}
	return 0
}

func (m *CPUInformation) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *CPUInformation) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *CPUInformation) GetVendorId() string {
	if m != nil {
		return m.VendorId
	}
	return ""
}

func (m *CPUInformation) GetMhz() float64 {
	if m != nil {
		return m.Mhz
	}
	return 0
}

func (m *CPUInformation) GetCacheSize() int32 {
	if m != nil {
		return m.CacheSize
	}
	return 0
}

func (m *CPUInformation) GetUsedPercent() float64 {
	if m != nil {
		return m.UsedPercent
	}
	return 0
}

func (m *CPUInformation) GetCoreID() string {
	if m != nil {
		return m.CoreID
	}
	return ""
}

func (m *CPUInformation) GetCores() int32 {
	if m != nil {
		return m.Cores
	}
	return 0
}

type HostInformation struct {
	Uptime                 uint64   `protobuf:"varint,1,opt,name=Uptime,proto3" json:"Uptime,omitempty"`
	OS                     string   `protobuf:"bytes,2,opt,name=OS,proto3" json:"OS,omitempty"`
	Platform               string   `protobuf:"bytes,3,opt,name=Platform,proto3" json:"Platform,omitempty"`
	PlatformFamily         string   `protobuf:"bytes,4,opt,name=PlatformFamily,proto3" json:"PlatformFamily,omitempty"`
	PlatformVersion        string   `protobuf:"bytes,5,opt,name=PlatformVersion,proto3" json:"PlatformVersion,omitempty"`
	NumberOfRunningProcess uint64   `protobuf:"varint,6,opt,name=NumberOfRunningProcess,proto3" json:"NumberOfRunningProcess,omitempty"`
	HostID                 string   `protobuf:"bytes,7,opt,name=HostID,proto3" json:"HostID,omitempty"`
	HostName               string   `protobuf:"bytes,8,opt,name=HostName,proto3" json:"HostName,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *HostInformation) Reset()         { *m = HostInformation{} }
func (m *HostInformation) String() string { return proto.CompactTextString(m) }
func (*HostInformation) ProtoMessage()    {}
func (*HostInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce967b26533b171c, []int{4}
}

func (m *HostInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostInformation.Unmarshal(m, b)
}
func (m *HostInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostInformation.Marshal(b, m, deterministic)
}
func (m *HostInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostInformation.Merge(m, src)
}
func (m *HostInformation) XXX_Size() int {
	return xxx_messageInfo_HostInformation.Size(m)
}
func (m *HostInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_HostInformation.DiscardUnknown(m)
}

var xxx_messageInfo_HostInformation proto.InternalMessageInfo

func (m *HostInformation) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *HostInformation) GetOS() string {
	if m != nil {
		return m.OS
	}
	return ""
}

func (m *HostInformation) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *HostInformation) GetPlatformFamily() string {
	if m != nil {
		return m.PlatformFamily
	}
	return ""
}

func (m *HostInformation) GetPlatformVersion() string {
	if m != nil {
		return m.PlatformVersion
	}
	return ""
}

func (m *HostInformation) GetNumberOfRunningProcess() uint64 {
	if m != nil {
		return m.NumberOfRunningProcess
	}
	return 0
}

func (m *HostInformation) GetHostID() string {
	if m != nil {
		return m.HostID
	}
	return ""
}

func (m *HostInformation) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

type MemoryInformation struct {
	Total uint64 `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	// This is the kernel's notion of free memory; RAM chips whose bits nobody
	// cares about the value of right now. For a human consumable number,
	// Available is what you really want.
	Free uint64 `protobuf:"varint,2,opt,name=Free,proto3" json:"Free,omitempty"`
	// RAM available for programs to allocate
	Available            uint64   `protobuf:"varint,3,opt,name=Available,proto3" json:"Available,omitempty"`
	Used                 uint64   `protobuf:"varint,4,opt,name=Used,proto3" json:"Used,omitempty"`
	UsedPercent          float64  `protobuf:"fixed64,5,opt,name=UsedPercent,proto3" json:"UsedPercent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemoryInformation) Reset()         { *m = MemoryInformation{} }
func (m *MemoryInformation) String() string { return proto.CompactTextString(m) }
func (*MemoryInformation) ProtoMessage()    {}
func (*MemoryInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce967b26533b171c, []int{5}
}

func (m *MemoryInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemoryInformation.Unmarshal(m, b)
}
func (m *MemoryInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemoryInformation.Marshal(b, m, deterministic)
}
func (m *MemoryInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemoryInformation.Merge(m, src)
}
func (m *MemoryInformation) XXX_Size() int {
	return xxx_messageInfo_MemoryInformation.Size(m)
}
func (m *MemoryInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_MemoryInformation.DiscardUnknown(m)
}

var xxx_messageInfo_MemoryInformation proto.InternalMessageInfo

func (m *MemoryInformation) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *MemoryInformation) GetFree() uint64 {
	if m != nil {
		return m.Free
	}
	return 0
}

func (m *MemoryInformation) GetAvailable() uint64 {
	if m != nil {
		return m.Available
	}
	return 0
}

func (m *MemoryInformation) GetUsed() uint64 {
	if m != nil {
		return m.Used
	}
	return 0
}

func (m *MemoryInformation) GetUsedPercent() float64 {
	if m != nil {
		return m.UsedPercent
	}
	return 0
}

type StorageInformation struct {
	FsType               string   `protobuf:"bytes,1,opt,name=FsType,proto3" json:"FsType,omitempty"`
	Total                uint64   `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"`
	Free                 uint64   `protobuf:"varint,3,opt,name=Free,proto3" json:"Free,omitempty"`
	Used                 uint64   `protobuf:"varint,4,opt,name=Used,proto3" json:"Used,omitempty"`
	UsedPercent          float64  `protobuf:"fixed64,5,opt,name=UsedPercent,proto3" json:"UsedPercent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StorageInformation) Reset()         { *m = StorageInformation{} }
func (m *StorageInformation) String() string { return proto.CompactTextString(m) }
func (*StorageInformation) ProtoMessage()    {}
func (*StorageInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce967b26533b171c, []int{6}
}

func (m *StorageInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageInformation.Unmarshal(m, b)
}
func (m *StorageInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageInformation.Marshal(b, m, deterministic)
}
func (m *StorageInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageInformation.Merge(m, src)
}
func (m *StorageInformation) XXX_Size() int {
	return xxx_messageInfo_StorageInformation.Size(m)
}
func (m *StorageInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageInformation.DiscardUnknown(m)
}

var xxx_messageInfo_StorageInformation proto.InternalMessageInfo

func (m *StorageInformation) GetFsType() string {
	if m != nil {
		return m.FsType
	}
	return ""
}

func (m *StorageInformation) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *StorageInformation) GetFree() uint64 {
	if m != nil {
		return m.Free
	}
	return 0
}

func (m *StorageInformation) GetUsed() uint64 {
	if m != nil {
		return m.Used
	}
	return 0
}

func (m *StorageInformation) GetUsedPercent() float64 {
	if m != nil {
		return m.UsedPercent
	}
	return 0
}

func init() {
	proto.RegisterType((*GetNodeHardwareResponse)(nil), "model.GetNodeHardwareResponse")
	proto.RegisterType((*GetNodeHardwareRequest)(nil), "model.GetNodeHardwareRequest")
	proto.RegisterType((*NodeHardware)(nil), "model.NodeHardware")
	proto.RegisterType((*CPUInformation)(nil), "model.CPUInformation")
	proto.RegisterType((*HostInformation)(nil), "model.HostInformation")
	proto.RegisterType((*MemoryInformation)(nil), "model.MemoryInformation")
	proto.RegisterType((*StorageInformation)(nil), "model.StorageInformation")
}

func init() { proto.RegisterFile("model/nodeHardware.proto", fileDescriptor_ce967b26533b171c) }

var fileDescriptor_ce967b26533b171c = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0xaf, 0xd2, 0x4c,
	0x14, 0x4e, 0x4b, 0xe1, 0xa5, 0xc3, 0x1b, 0xae, 0x8e, 0x8a, 0xd5, 0xb8, 0x20, 0x5d, 0x18, 0x62,
	0x22, 0x24, 0xd7, 0x44, 0x57, 0x26, 0x2a, 0x37, 0x78, 0x59, 0xf0, 0x91, 0xe1, 0x72, 0x17, 0xee,
	0x86, 0xf6, 0x5c, 0x68, 0xd2, 0x76, 0x70, 0x3a, 0xa8, 0xf0, 0x1b, 0xdc, 0x18, 0xfd, 0x05, 0xfe,
	0x52, 0x33, 0xa7, 0x05, 0x4a, 0x8b, 0x1b, 0x37, 0xcd, 0x79, 0xce, 0xc7, 0x33, 0xe7, 0x3c, 0x67,
	0xa6, 0xc4, 0x89, 0x84, 0x0f, 0x61, 0x2f, 0x16, 0x3e, 0x5c, 0x73, 0xe9, 0x7f, 0xe5, 0x12, 0xba,
	0x6b, 0x29, 0x94, 0xa0, 0x55, 0x8c, 0xb8, 0x8c, 0x3c, 0xfe, 0x08, 0x6a, 0x9c, 0x8b, 0x33, 0x48,
	0xd6, 0x22, 0x4e, 0x80, 0xbe, 0x21, 0xff, 0xe7, 0xfd, 0x8e, 0xd1, 0x36, 0x3a, 0x8d, 0xcb, 0x07,
	0x5d, 0x2c, 0xec, 0x9e, 0x94, 0x9c, 0x24, 0xba, 0x0e, 0x69, 0x95, 0x38, 0x3f, 0x6f, 0x20, 0x51,
	0xee, 0x6f, 0xf3, 0x94, 0x93, 0xbe, 0x25, 0xcd, 0xfe, 0x74, 0x3e, 0x8c, 0xef, 0x84, 0x8c, 0xb8,
	0x0a, 0x44, 0xec, 0x18, 0xed, 0x4a, 0xa7, 0x71, 0xf9, 0x28, 0x3b, 0xe5, 0x34, 0xc8, 0x0a, 0xc9,
	0x74, 0x40, 0xee, 0x8f, 0x20, 0x12, 0x72, 0x9b, 0x67, 0x30, 0xb1, 0x4f, 0x27, 0x63, 0x28, 0xc5,
	0x59, 0xb9, 0x84, 0x0e, 0x09, 0x9d, 0x29, 0x21, 0xf9, 0x12, 0xf2, 0x44, 0x15, 0x24, 0x7a, 0x92,
	0x11, 0x95, 0x13, 0xd8, 0x99, 0x22, 0xfa, 0x8e, 0x5c, 0x5c, 0x8b, 0x44, 0xe5, 0x79, 0x2c, 0xe4,
	0x69, 0x65, 0x3c, 0x85, 0x28, 0x2b, 0xa6, 0xbb, 0xbf, 0xcc, 0xa2, 0x28, 0xb4, 0x45, 0x6a, 0x03,
	0x1e, 0x05, 0xe1, 0x16, 0x97, 0x60, 0xb3, 0x0c, 0xd1, 0xa7, 0xa4, 0x8e, 0x99, 0x3e, 0x7c, 0xc3,
	0xb1, 0xab, 0xec, 0x80, 0xe9, 0x43, 0x52, 0x1d, 0xe9, 0x03, 0x71, 0x0c, 0x9b, 0xa5, 0x80, 0x3e,
	0x23, 0x36, 0x1a, 0x63, 0x1e, 0x01, 0x36, 0x66, 0xb3, 0xa3, 0x43, 0xf3, 0xdd, 0x42, 0xec, 0x0b,
	0x39, 0xf4, 0x9d, 0x2a, 0x06, 0x0f, 0x98, 0xde, 0x23, 0x95, 0xd1, 0x6a, 0xe7, 0xd4, 0xda, 0x46,
	0xc7, 0x60, 0xda, 0xd4, 0x5c, 0x7d, 0xee, 0xad, 0x60, 0x16, 0xec, 0xc0, 0xf9, 0x0f, 0x8f, 0x3f,
	0x3a, 0x68, 0x9b, 0x34, 0xe6, 0x09, 0xf8, 0x53, 0x90, 0x1e, 0xc4, 0xca, 0xa9, 0x63, 0x5d, 0xde,
	0xa5, 0xa7, 0xea, 0x0b, 0x09, 0xc3, 0x2b, 0xc7, 0x4e, 0xa7, 0x4a, 0x91, 0xee, 0x5c, 0x5b, 0x89,
	0x43, 0x90, 0x33, 0x05, 0xee, 0x4f, 0xb3, 0xa4, 0xac, 0x66, 0x98, 0xaf, 0x55, 0x10, 0xa5, 0x97,
	0xd3, 0x62, 0x19, 0xa2, 0x4d, 0x62, 0x4e, 0x66, 0xa8, 0x88, 0xcd, 0xcc, 0xc9, 0x4c, 0xcf, 0x35,
	0x0d, 0xb9, 0xd2, 0x85, 0x99, 0x1c, 0x07, 0x4c, 0x9f, 0x93, 0xe6, 0xde, 0xce, 0x34, 0x4e, 0x65,
	0x29, 0x78, 0x69, 0x87, 0x5c, 0xec, 0x3d, 0xb7, 0x20, 0x13, 0xbd, 0xd8, 0x54, 0xa2, 0xa2, 0x9b,
	0xbe, 0x26, 0xad, 0xf1, 0x26, 0x5a, 0x80, 0x9c, 0xdc, 0xb1, 0x4d, 0x1c, 0x07, 0xf1, 0x72, 0x2a,
	0x85, 0x07, 0x49, 0x82, 0xe2, 0x59, 0xec, 0x2f, 0x51, 0x3d, 0x0d, 0x0e, 0x78, 0x85, 0x62, 0xda,
	0x2c, 0x43, 0xba, 0x7b, 0x6d, 0xe1, 0xca, 0xea, 0x69, 0xf7, 0x7b, 0xec, 0xfe, 0x30, 0xce, 0x3c,
	0x01, 0xad, 0xe0, 0x8d, 0x50, 0x3c, 0xcc, 0x64, 0x49, 0x01, 0xa5, 0xc4, 0x1a, 0x48, 0x00, 0xd4,
	0xc5, 0x62, 0x68, 0xeb, 0x1d, 0xbe, 0xff, 0xc2, 0x83, 0x90, 0x2f, 0x42, 0x40, 0x69, 0x2c, 0x76,
	0x74, 0xe8, 0x0a, 0xbd, 0x30, 0x54, 0xc4, 0x62, 0x68, 0x17, 0xf7, 0x5a, 0x2d, 0xed, 0xd5, 0xfd,
	0x6e, 0x9c, 0x7b, 0x4e, 0x78, 0x89, 0x93, 0x9b, 0xed, 0x1a, 0x0e, 0x97, 0x18, 0xd1, 0xb1, 0x59,
	0xf3, 0x5c, 0xb3, 0x95, 0x5c, 0xb3, 0xff, 0xd4, 0xce, 0x87, 0x17, 0x9f, 0x3a, 0xcb, 0x40, 0xad,
	0x36, 0x8b, 0xae, 0x27, 0xa2, 0xde, 0x4e, 0x88, 0x85, 0x97, 0x7e, 0x5f, 0x7a, 0x42, 0x42, 0xcf,
	0x13, 0x51, 0x24, 0xe2, 0x1e, 0x3e, 0xce, 0x45, 0x0d, 0x7f, 0x8e, 0xaf, 0xfe, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x18, 0x0c, 0x98, 0x0b, 0x38, 0x05, 0x00, 0x00,
}