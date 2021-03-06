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
package chaintype

import (
	"time"

	"github.com/zoobc/zoobc-core/common/constant"
)

// MainChain is struct should has methods in below
type MainChain struct{}

// GetTypeInt return the value of the chain type in int
func (*MainChain) GetTypeInt() int32 {
	return 0
}

// GetTablePrefix return the value of current chain table prefix in the database
func (*MainChain) GetTablePrefix() string {
	return "main"
}

func (*MainChain) GetSmithingPeriod() int64 {
	return constant.MainChainSmithingPeriod
}

func (*MainChain) GetBlocksmithTimeGap() int64 {
	return constant.MainSmithingBlocksmithTimeGap
}

func (*MainChain) GetBlocksmithBlockCreationTime() int64 {
	return constant.MainSmithingBlockCreationTime
}

func (*MainChain) GetBlocksmithNetworkTolerance() int64 {
	return constant.MainSmithingNetworkTolerance
}

// GetName return the name of the chain : used in parsing chaintype across node
func (*MainChain) GetName() string {
	return "Mainchain"
}

// GetGenesisBlockID return the block ID of genesis block in the chain
func (*MainChain) GetGenesisBlockID() int64 {
	return constant.MainchainGenesisBlockID
}

func (*MainChain) GetGenesisBlockSeed() []byte {
	return constant.MainchainGenesisBlockSeed
}

func (*MainChain) GetGenesisNodePublicKey() []byte {
	return constant.MainchainGenesisNodePublicKey
}

func (*MainChain) GetGenesisBlockTimestamp() int64 {
	return constant.MainchainGenesisBlockTimestamp
}

func (*MainChain) GetGenesisBlockSignature() []byte {
	return constant.MainchainGenesisBlockSignature
}

func (*MainChain) HasTransactions() bool {
	return true
}

func (*MainChain) HasSnapshots() bool {
	return true
}

func (*MainChain) GetSnapshotInterval() uint32 {
	return constant.MainchainSnapshotInterval
}

func (*MainChain) GetSnapshotGenerationTimeout() time.Duration {
	return constant.MainchainSnapshotGenerationTimeout
}
