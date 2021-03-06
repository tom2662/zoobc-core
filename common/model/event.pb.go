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
// source: model/event.proto

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

type EventType int32

const (
	EventType_EventAny                               EventType = 0
	EventType_EventSendMoneyTransaction              EventType = 1
	EventType_EventNodeRegistrationTransaction       EventType = 2
	EventType_EventUpdateNodeRegistrationTransaction EventType = 3
	EventType_EventRemoveNodeRegistrationTransaction EventType = 4
	EventType_EventClaimNodeRegistrationTransaction  EventType = 5
	EventType_EventSetupAccountDatasetTransaction    EventType = 6
	EventType_EventRemoveAccountDatasetTransaction   EventType = 7
	EventType_EventReward                            EventType = 8
	EventType_EventApprovalEscrowTransaction         EventType = 9
	EventType_EventMultiSignatureTransaction         EventType = 10
	EventType_EventFeeVoteCommitTransaction          EventType = 11
	EventType_EventFeeVoteRevealTransaction          EventType = 12
	EventType_EventLiquidPaymentTransaction          EventType = 13
	EventType_EventLiquidPaymentPaidTransaction      EventType = 14
	EventType_EventLiquidPaymentStopTransaction      EventType = 15
	EventType_EventEscrowedTransaction               EventType = 16
)

var EventType_name = map[int32]string{
	0:  "EventAny",
	1:  "EventSendMoneyTransaction",
	2:  "EventNodeRegistrationTransaction",
	3:  "EventUpdateNodeRegistrationTransaction",
	4:  "EventRemoveNodeRegistrationTransaction",
	5:  "EventClaimNodeRegistrationTransaction",
	6:  "EventSetupAccountDatasetTransaction",
	7:  "EventRemoveAccountDatasetTransaction",
	8:  "EventReward",
	9:  "EventApprovalEscrowTransaction",
	10: "EventMultiSignatureTransaction",
	11: "EventFeeVoteCommitTransaction",
	12: "EventFeeVoteRevealTransaction",
	13: "EventLiquidPaymentTransaction",
	14: "EventLiquidPaymentPaidTransaction",
	15: "EventLiquidPaymentStopTransaction",
	16: "EventEscrowedTransaction",
}

var EventType_value = map[string]int32{
	"EventAny":                               0,
	"EventSendMoneyTransaction":              1,
	"EventNodeRegistrationTransaction":       2,
	"EventUpdateNodeRegistrationTransaction": 3,
	"EventRemoveNodeRegistrationTransaction": 4,
	"EventClaimNodeRegistrationTransaction":  5,
	"EventSetupAccountDatasetTransaction":    6,
	"EventRemoveAccountDatasetTransaction":   7,
	"EventReward":                            8,
	"EventApprovalEscrowTransaction":         9,
	"EventMultiSignatureTransaction":         10,
	"EventFeeVoteCommitTransaction":          11,
	"EventFeeVoteRevealTransaction":          12,
	"EventLiquidPaymentTransaction":          13,
	"EventLiquidPaymentPaidTransaction":      14,
	"EventLiquidPaymentStopTransaction":      15,
	"EventEscrowedTransaction":               16,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_24dabb9f57ff37c9, []int{0}
}

func init() {
	proto.RegisterEnum("model.EventType", EventType_name, EventType_value)
}

func init() {
	proto.RegisterFile("model/event.proto", fileDescriptor_24dabb9f57ff37c9)
}

var fileDescriptor_24dabb9f57ff37c9 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd2, 0xcf, 0x4e, 0x32, 0x31,
	0x10, 0x00, 0xf0, 0xef, 0x53, 0x40, 0x28, 0x28, 0x6b, 0x4f, 0x9a, 0x88, 0x11, 0x05, 0x45, 0x12,
	0xd9, 0x83, 0x4f, 0x80, 0x88, 0x27, 0x31, 0x04, 0xd0, 0x83, 0xb7, 0xd2, 0x9d, 0x60, 0x93, 0x6d,
	0x67, 0xed, 0x76, 0x97, 0xac, 0x8f, 0xeb, 0x93, 0x18, 0x0a, 0x21, 0xc5, 0x3f, 0x5c, 0x36, 0x99,
	0x99, 0xdf, 0x76, 0x32, 0xed, 0x90, 0x43, 0x89, 0x01, 0x84, 0x3e, 0xa4, 0xa0, 0x4c, 0x27, 0xd2,
	0x68, 0x90, 0xe6, 0x6d, 0xaa, 0xfd, 0x99, 0x23, 0xa5, 0xfe, 0x22, 0x3d, 0xc9, 0x22, 0xa0, 0x15,
	0x52, 0xb4, 0x41, 0x57, 0x65, 0xde, 0x3f, 0x5a, 0x23, 0xc7, 0x36, 0x1a, 0x83, 0x0a, 0x06, 0xa8,
	0x20, 0x9b, 0x68, 0xa6, 0x62, 0xc6, 0x8d, 0x40, 0xe5, 0xfd, 0xa7, 0x0d, 0x72, 0x66, 0xcb, 0x4f,
	0x18, 0xc0, 0x08, 0x66, 0x22, 0x36, 0x9a, 0x2d, 0x4a, 0xae, 0xda, 0xa1, 0x6d, 0x72, 0x69, 0xd5,
	0x73, 0x14, 0x30, 0x03, 0xdb, 0xec, 0xee, 0xda, 0x8e, 0x40, 0x62, 0xba, 0xd5, 0xe6, 0xe8, 0x35,
	0x69, 0x5a, 0xdb, 0x0b, 0x99, 0x90, 0xdb, 0x68, 0x9e, 0x5e, 0x91, 0x8b, 0xd5, 0x1c, 0x26, 0x89,
	0xba, 0x9c, 0x63, 0xa2, 0xcc, 0x3d, 0x33, 0x2c, 0x06, 0xe3, 0xc2, 0x02, 0x6d, 0x91, 0x86, 0xd3,
	0xff, 0x6f, 0xb9, 0x47, 0xab, 0xa4, 0xbc, 0x92, 0x73, 0xa6, 0x03, 0xaf, 0x48, 0xcf, 0xc9, 0xe9,
	0xf2, 0xe6, 0xa2, 0x48, 0x63, 0xca, 0xc2, 0x7e, 0xcc, 0x35, 0xce, 0xdd, 0x9f, 0x4a, 0x6b, 0x33,
	0x48, 0x42, 0x23, 0xc6, 0x62, 0xa6, 0x98, 0x49, 0x34, 0xb8, 0x86, 0xd0, 0x3a, 0xa9, 0x59, 0xf3,
	0x00, 0xf0, 0x82, 0x06, 0x7a, 0x28, 0xa5, 0xd8, 0xe8, 0x5d, 0xfe, 0x4e, 0x46, 0x90, 0x02, 0x0b,
	0x5d, 0x52, 0x59, 0x93, 0x47, 0xf1, 0x9e, 0x88, 0x60, 0xc8, 0x32, 0xb9, 0x78, 0x61, 0x87, 0xec,
	0xd3, 0x26, 0xa9, 0xff, 0x24, 0x43, 0x26, 0x02, 0x97, 0x1d, 0xfc, 0xce, 0xc6, 0x06, 0x23, 0x97,
	0x55, 0xe9, 0x09, 0x39, 0xb2, 0x6c, 0x39, 0x36, 0x6c, 0x1c, 0xe2, 0xdd, 0xb5, 0x5f, 0x5b, 0x33,
	0x61, 0xde, 0x92, 0x69, 0x87, 0xa3, 0xf4, 0x3f, 0x10, 0xa7, 0x7c, 0xf9, 0xbd, 0xe1, 0xa8, 0xc1,
	0xe7, 0x28, 0x25, 0x2a, 0xdf, 0x2e, 0xe4, 0xb4, 0x60, 0xd7, 0xf3, 0xf6, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0x57, 0x47, 0x7f, 0x63, 0xb3, 0x02, 0x00, 0x00,
}
