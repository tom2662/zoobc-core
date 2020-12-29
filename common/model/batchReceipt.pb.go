// ZooBC Copyright (C) 2020 Quasisoft Limited - Hong Kong
// This file is part of ZooBC <https://github.com/zoobc/zoobc-core>
//
// ZooBC is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ZooBC is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ZooBC.  If not, see <http://www.gnu.org/licenses/>.
//
// Additional Permission Under GNU GPL Version 3 section 7.
// As the special exception permitted under Section 7b, c and e,
// in respect with the Author’s copyright, please refer to this section:
//
// 1. You are free to convey this Program according to GNU GPL Version 3,
//     as long as you respect and comply with the Author’s copyright by
//     showing in its user interface an Appropriate Notice that the derivate
//     program and its source code are “powered by ZooBC”.
//     This is an acknowledgement for the copyright holder, ZooBC,
//     as the implementation of appreciation of the exclusive right of the
//     creator and to avoid any circumvention on the rights under trademark
//     law for use of some trade names, trademarks, or service marks.
//
// 2. Complying to the GNU GPL Version 3, you may distribute
//     the program without any permission from the Author.
//     However a prior notification to the authors will be appreciated.
//
// ZooBC is architected by Roberto Capodieci & Barton Johnston
//             contact us at roberto.capodieci[at]blockchainzoo.com
//             and barton.johnston[at]blockchainzoo.com
//
// Core developers that contributed to the current implementation of the
// software are:
//             Ahmad Ali Abdilah ahmad.abdilah[at]blockchainzoo.com
//             Allan Bintoro allan.bintoro[at]blockchainzoo.com
//             Andy Herman
//             Gede Sukra
//             Ketut Ariasa
//             Nawi Kartini nawi.kartini[at]blockchainzoo.com
//             Stefano Galassi stefano.galassi[at]blockchainzoo.com
//
// IMPORTANT: The above copyright notice and this permission notice
// shall be included in all copies or substantial portions of the Software.
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/batchReceipt.proto

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

// BatchReceipt is a receipt that already includes RMR ready to publish
type BatchReceipt struct {
	Receipt              *Receipt `protobuf:"bytes,1,opt,name=Receipt,proto3" json:"Receipt,omitempty"`
	RMR                  []byte   `protobuf:"bytes,2,opt,name=RMR,proto3" json:"RMR,omitempty"`
	RMRIndex             uint32   `protobuf:"varint,3,opt,name=RMRIndex,proto3" json:"RMRIndex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchReceipt) Reset()         { *m = BatchReceipt{} }
func (m *BatchReceipt) String() string { return proto.CompactTextString(m) }
func (*BatchReceipt) ProtoMessage()    {}
func (*BatchReceipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3d73f8c2557fbad, []int{0}
}

func (m *BatchReceipt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchReceipt.Unmarshal(m, b)
}
func (m *BatchReceipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchReceipt.Marshal(b, m, deterministic)
}
func (m *BatchReceipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchReceipt.Merge(m, src)
}
func (m *BatchReceipt) XXX_Size() int {
	return xxx_messageInfo_BatchReceipt.Size(m)
}
func (m *BatchReceipt) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchReceipt.DiscardUnknown(m)
}

var xxx_messageInfo_BatchReceipt proto.InternalMessageInfo

func (m *BatchReceipt) GetReceipt() *Receipt {
	if m != nil {
		return m.Receipt
	}
	return nil
}

func (m *BatchReceipt) GetRMR() []byte {
	if m != nil {
		return m.RMR
	}
	return nil
}

func (m *BatchReceipt) GetRMRIndex() uint32 {
	if m != nil {
		return m.RMRIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*BatchReceipt)(nil), "model.BatchReceipt")
}

func init() {
	proto.RegisterFile("model/batchReceipt.proto", fileDescriptor_f3d73f8c2557fbad)
}

var fileDescriptor_f3d73f8c2557fbad = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0x4f, 0x4a, 0x2c, 0x49, 0xce, 0x08, 0x4a, 0x4d, 0x4e, 0xcd, 0x2c, 0x28, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0xcb, 0x48, 0x09, 0x43, 0x14, 0x14, 0x21, 0xcb, 0x29,
	0xa5, 0x71, 0xf1, 0x38, 0x21, 0xe9, 0x10, 0xd2, 0xe0, 0x62, 0x87, 0x32, 0x25, 0x18, 0x15, 0x18,
	0x35, 0xb8, 0x8d, 0xf8, 0xf4, 0xc0, 0xda, 0xf4, 0xa0, 0xa2, 0x41, 0x30, 0x69, 0x21, 0x01, 0x2e,
	0xe6, 0x20, 0xdf, 0x20, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x10, 0x53, 0x48, 0x8a, 0x8b,
	0x23, 0xc8, 0x37, 0xc8, 0x33, 0x2f, 0x25, 0xb5, 0x42, 0x82, 0x59, 0x81, 0x51, 0x83, 0x37, 0x08,
	0xce, 0x77, 0xd2, 0x8a, 0xd2, 0x48, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5,
	0xaf, 0xca, 0xcf, 0x4f, 0x4a, 0x86, 0x90, 0xba, 0xc9, 0xf9, 0x45, 0xa9, 0xfa, 0xc9, 0xf9, 0xb9,
	0xb9, 0xf9, 0x79, 0xfa, 0x60, 0xab, 0x92, 0xd8, 0xc0, 0x4e, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xc4, 0x96, 0x18, 0x6e, 0xd2, 0x00, 0x00, 0x00,
}
