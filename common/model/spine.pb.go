// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/spine.proto

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

// RequestType used to sign a node administration request
type SpinePublicKeyAction int32

const (
	SpinePublicKeyAction_AddKey    SpinePublicKeyAction = 0
	SpinePublicKeyAction_RemoveKey SpinePublicKeyAction = 1
)

var SpinePublicKeyAction_name = map[int32]string{
	0: "AddKey",
	1: "RemoveKey",
}

var SpinePublicKeyAction_value = map[string]int32{
	"AddKey":    0,
	"RemoveKey": 1,
}

func (x SpinePublicKeyAction) String() string {
	return proto.EnumName(SpinePublicKeyAction_name, int32(x))
}

func (SpinePublicKeyAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_59a5e7c6458ac090, []int{0}
}

// Block represent the block data structure stored in the database
type SpinePublicKey struct {
	NodePublicKey        []byte               `protobuf:"bytes,1,opt,name=NodePublicKey,proto3" json:"NodePublicKey,omitempty"`
	NodeID               int64                `protobuf:"varint,2,opt,name=NodeID,proto3" json:"NodeID,omitempty"`
	MainBlockHeight      uint32               `protobuf:"varint,3,opt,name=MainBlockHeight,proto3" json:"MainBlockHeight,omitempty"`
	PublicKeyAction      SpinePublicKeyAction `protobuf:"varint,4,opt,name=PublicKeyAction,proto3,enum=model.SpinePublicKeyAction" json:"PublicKeyAction,omitempty"`
	Latest               bool                 `protobuf:"varint,5,opt,name=Latest,proto3" json:"Latest,omitempty"`
	Height               uint32               `protobuf:"varint,6,opt,name=Height,proto3" json:"Height,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SpinePublicKey) Reset()         { *m = SpinePublicKey{} }
func (m *SpinePublicKey) String() string { return proto.CompactTextString(m) }
func (*SpinePublicKey) ProtoMessage()    {}
func (*SpinePublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_59a5e7c6458ac090, []int{0}
}

func (m *SpinePublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpinePublicKey.Unmarshal(m, b)
}
func (m *SpinePublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpinePublicKey.Marshal(b, m, deterministic)
}
func (m *SpinePublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpinePublicKey.Merge(m, src)
}
func (m *SpinePublicKey) XXX_Size() int {
	return xxx_messageInfo_SpinePublicKey.Size(m)
}
func (m *SpinePublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_SpinePublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_SpinePublicKey proto.InternalMessageInfo

func (m *SpinePublicKey) GetNodePublicKey() []byte {
	if m != nil {
		return m.NodePublicKey
	}
	return nil
}

func (m *SpinePublicKey) GetNodeID() int64 {
	if m != nil {
		return m.NodeID
	}
	return 0
}

func (m *SpinePublicKey) GetMainBlockHeight() uint32 {
	if m != nil {
		return m.MainBlockHeight
	}
	return 0
}

func (m *SpinePublicKey) GetPublicKeyAction() SpinePublicKeyAction {
	if m != nil {
		return m.PublicKeyAction
	}
	return SpinePublicKeyAction_AddKey
}

func (m *SpinePublicKey) GetLatest() bool {
	if m != nil {
		return m.Latest
	}
	return false
}

func (m *SpinePublicKey) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterEnum("model.SpinePublicKeyAction", SpinePublicKeyAction_name, SpinePublicKeyAction_value)
	proto.RegisterType((*SpinePublicKey)(nil), "model.SpinePublicKey")
}

func init() { proto.RegisterFile("model/spine.proto", fileDescriptor_59a5e7c6458ac090) }

var fileDescriptor_59a5e7c6458ac090 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd0, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0x06, 0x70, 0xd3, 0xda, 0xa0, 0x83, 0x6d, 0x35, 0x88, 0x04, 0xbd, 0x04, 0xf1, 0x10, 0x0a,
	0xee, 0xfa, 0xe7, 0x09, 0xba, 0x28, 0x28, 0x55, 0x91, 0x78, 0xf3, 0xd6, 0xcd, 0x0e, 0x6d, 0x70,
	0xb3, 0x53, 0xda, 0x54, 0xa8, 0xaf, 0xed, 0x0b, 0xc8, 0xa6, 0x8b, 0xd2, 0xe2, 0x25, 0xf0, 0xfd,
	0x48, 0x66, 0x3e, 0x02, 0x47, 0x9e, 0x0a, 0x2c, 0xd3, 0xc5, 0xcc, 0x55, 0x98, 0xcc, 0xe6, 0x14,
	0x48, 0x74, 0x22, 0x9d, 0x7f, 0x33, 0xe8, 0xbd, 0xd5, 0xfc, 0xba, 0xcc, 0x4b, 0x67, 0x47, 0xb8,
	0x12, 0x17, 0xd0, 0x7d, 0xa1, 0xe2, 0x0f, 0x24, 0x53, 0x4c, 0x1f, 0x98, 0x4d, 0x14, 0xa7, 0xc0,
	0x6b, 0x78, 0xbc, 0x93, 0x2d, 0xc5, 0x74, 0x3b, 0x6b, 0x5d, 0x31, 0xd3, 0x88, 0xd0, 0xd0, 0x7f,
	0x1e, 0xbb, 0x2a, 0x2b, 0xc9, 0x7e, 0x3c, 0xa0, 0x9b, 0x4c, 0x83, 0x6c, 0x2b, 0xa6, 0xbb, 0x66,
	0x9b, 0xc5, 0x3d, 0xf4, 0x7f, 0x47, 0x0e, 0x6d, 0x70, 0x54, 0xc9, 0x5d, 0xc5, 0x74, 0xef, 0xe6,
	0x2c, 0x89, 0xfd, 0x92, 0xcd, 0x6e, 0xeb, 0x2b, 0x66, 0xfb, 0x8d, 0x38, 0x01, 0xfe, 0x34, 0x0e,
	0xb8, 0x08, 0xb2, 0xa3, 0x98, 0xde, 0x33, 0x4d, 0xaa, 0xbd, 0xd9, 0xcf, 0xe3, 0xfe, 0x26, 0x0d,
	0xae, 0xe1, 0xf8, 0xbf, 0xc1, 0x02, 0x80, 0x0f, 0x8b, 0x62, 0x84, 0xab, 0xc3, 0x1d, 0xd1, 0x85,
	0x7d, 0x83, 0x9e, 0x3e, 0xb1, 0x8e, 0x2c, 0x1b, 0xbc, 0xeb, 0x89, 0x0b, 0xd3, 0x65, 0x9e, 0x58,
	0xf2, 0xe9, 0x17, 0x51, 0x6e, 0xd7, 0xe7, 0xa5, 0xa5, 0x39, 0xa6, 0x96, 0xbc, 0xa7, 0x2a, 0x8d,
	0xa5, 0x73, 0x1e, 0xbf, 0xf8, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x60, 0x79, 0xa1, 0xbc, 0x77,
	0x01, 0x00, 0x00,
}
