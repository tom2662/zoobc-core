// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/signature.proto

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

type SignatureType int32

const (
	// in bytes: []byte{0,0,0,0}, using Ed25519 signature algorithm
	SignatureType_DefaultSignature SignatureType = 0
	// in bytes: []byte{1,0,0,0}, bitcoin uses a specific Koblitz curve secp256k1
	// Koblitz curves are a type of Elliptic Curve Digital Signature Algorithm
	SignatureType_BitcoinSignature SignatureType = 1
	// in bytes: []byte{2,0,0,0} for multisig validation purpose only
	SignatureType_MultisigSignature SignatureType = 2
)

var SignatureType_name = map[int32]string{
	0: "DefaultSignature",
	1: "BitcoinSignature",
	2: "MultisigSignature",
}

var SignatureType_value = map[string]int32{
	"DefaultSignature":  0,
	"BitcoinSignature":  1,
	"MultisigSignature": 2,
}

func (x SignatureType) String() string {
	return proto.EnumName(SignatureType_name, int32(x))
}

func (SignatureType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a69ee5fbbdd37ed5, []int{0}
}

func init() {
	proto.RegisterEnum("model.SignatureType", SignatureType_name, SignatureType_value)
}

func init() { proto.RegisterFile("model/signature.proto", fileDescriptor_a69ee5fbbdd37ed5) }

var fileDescriptor_a69ee5fbbdd37ed5 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0x2f, 0xce, 0x4c, 0xcf, 0x4b, 0x2c, 0x29, 0x2d, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x05, 0x0b, 0x6b, 0x05, 0x71, 0xf1, 0x06, 0xc3, 0x64, 0x42, 0x2a, 0x0b, 0x52,
	0x85, 0x44, 0xb8, 0x04, 0x5c, 0x52, 0xd3, 0x12, 0x4b, 0x73, 0x4a, 0xe0, 0xe2, 0x02, 0x0c, 0x20,
	0x51, 0xa7, 0xcc, 0x92, 0xe4, 0xfc, 0xcc, 0x3c, 0x84, 0x28, 0xa3, 0x90, 0x28, 0x97, 0xa0, 0x6f,
	0x69, 0x4e, 0x49, 0x66, 0x71, 0x66, 0x3a, 0x42, 0x98, 0xc9, 0x49, 0x2b, 0x4a, 0x23, 0x3d, 0xb3,
	0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xbf, 0x2a, 0x3f, 0x3f, 0x29, 0x19, 0x42, 0xea,
	0x26, 0xe7, 0x17, 0xa5, 0xea, 0x27, 0xe7, 0xe7, 0xe6, 0xe6, 0xe7, 0xe9, 0x83, 0xed, 0x4f, 0x62,
	0x03, 0xbb, 0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x53, 0x54, 0x61, 0xa6, 0x00, 0x00,
	0x00,
}