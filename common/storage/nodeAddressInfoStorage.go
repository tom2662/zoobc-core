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
package storage

import (
	"bytes"
	"encoding/gob"
	"strconv"
	"sync"

	"github.com/zoobc/zoobc-core/common/monitoring"

	"github.com/zoobc/zoobc-core/common/blocker"
	"github.com/zoobc/zoobc-core/common/model"
)

type (
	// NodeAddressInfoStorage represent list of node address info
	NodeAddressInfoStorage struct {
		sync.RWMutex
		transactionalLock                          sync.RWMutex
		isInTransaction                            bool
		nodeAddressInfoMapByID                     map[int64]map[string]model.NodeAddressInfo
		nodeAddressInfoMapByAddressPort            map[string]map[int64]bool
		nodeAddressInfoMapByStatus                 map[model.NodeAddressStatus]map[int64]map[string]bool
		transactionalRemovedNodeAddressInfoMapByID map[int64]map[string]bool
	}
	// NodeAddressInfoStorageKey represent a key for NodeAddressInfoStorage
	NodeAddressInfoStorageKey struct {
		NodeID      int64
		AddressPort string
		Statuses    []model.NodeAddressStatus
	}
)

func NewNodeAddressInfoStorage() *NodeAddressInfoStorage {
	return &NodeAddressInfoStorage{
		nodeAddressInfoMapByID:                     make(map[int64]map[string]model.NodeAddressInfo),
		nodeAddressInfoMapByAddressPort:            make(map[string]map[int64]bool),
		nodeAddressInfoMapByStatus:                 make(map[model.NodeAddressStatus]map[int64]map[string]bool),
		transactionalRemovedNodeAddressInfoMapByID: make(map[int64]map[string]bool),
	}
}

func (nas *NodeAddressInfoStorage) SetItem(key, item interface{}) error {
	var nodeAddressInfo, ok = item.(model.NodeAddressInfo)
	if !ok {
		return blocker.NewBlocker(blocker.ValidationErr, "WrongTypeKey")
	}
	fullAddress := nodeAddressInfo.Address + ":" + strconv.Itoa(int(nodeAddressInfo.Port))
	nas.Lock()
	defer nas.Unlock()
	if nas.nodeAddressInfoMapByID[nodeAddressInfo.NodeID] == nil {
		nas.nodeAddressInfoMapByID[nodeAddressInfo.NodeID] = make(map[string]model.NodeAddressInfo)
	}
	nas.nodeAddressInfoMapByID[nodeAddressInfo.NodeID][fullAddress] = nodeAddressInfo

	if nas.nodeAddressInfoMapByAddressPort[fullAddress] == nil {
		nas.nodeAddressInfoMapByAddressPort[fullAddress] = make(map[int64]bool)
	}
	nas.nodeAddressInfoMapByAddressPort[fullAddress][nodeAddressInfo.NodeID] = true

	if nas.nodeAddressInfoMapByStatus[nodeAddressInfo.Status] == nil {
		nas.nodeAddressInfoMapByStatus[nodeAddressInfo.Status] = make(map[int64]map[string]bool)
	}
	if nas.nodeAddressInfoMapByStatus[nodeAddressInfo.Status][nodeAddressInfo.NodeID] == nil {
		nas.nodeAddressInfoMapByStatus[nodeAddressInfo.Status][nodeAddressInfo.NodeID] = make(map[string]bool)
	}
	nas.nodeAddressInfoMapByStatus[nodeAddressInfo.Status][nodeAddressInfo.NodeID][fullAddress] = true

	if monitoring.IsMonitoringActive() {
		monitoring.SetCacheStorageMetrics(monitoring.TypeNodeAddressInfoCacheStorage, float64(nas.size()))
	}
	return nil
}

func (nas *NodeAddressInfoStorage) SetItems(item interface{}) error {
	return nil
}

func (nas *NodeAddressInfoStorage) GetItem(key, item interface{}) error {
	if !nas.isInTransaction {
		nas.RLock()
		defer nas.RUnlock()
	} else {
		nas.transactionalLock.RLock()
		defer nas.transactionalLock.RUnlock()
	}
	storageKey, ok := key.(NodeAddressInfoStorageKey)
	if !ok {
		return blocker.NewBlocker(blocker.ValidationErr, "WrongTypeKeyExpected:NodeAddressInfoStorageKey")
	}
	nodeAddresses, ok := item.(*[]*model.NodeAddressInfo)
	if !ok {
		return blocker.NewBlocker(blocker.ValidationErr, "WrongTypeKeyExpected:*[]*model.NodeAddressInfo")
	}
	// staus node address info is always required when getting node address info
	if len(storageKey.Statuses) == 0 {
		return blocker.NewBlocker(blocker.ValidationErr, "StatusNodeAddressInfoRequired")
	}
	switch {
	case storageKey.NodeID != 0:
		for _, status := range storageKey.Statuses {
			for fullAddressPort := range nas.nodeAddressInfoMapByStatus[status][storageKey.NodeID] {
				*nodeAddresses = nas.append(*nodeAddresses, nas.nodeAddressInfoMapByID[storageKey.NodeID][fullAddressPort])
			}
		}
	case storageKey.AddressPort != "":
		for _, status := range storageKey.Statuses {
			for nodeID := range nas.nodeAddressInfoMapByAddressPort[storageKey.AddressPort] {
				if nas.nodeAddressInfoMapByStatus[status][nodeID][storageKey.AddressPort] {
					*nodeAddresses = nas.append(*nodeAddresses, nas.nodeAddressInfoMapByID[nodeID][storageKey.AddressPort])
				}
			}
		}
	default:
		for _, status := range storageKey.Statuses {
			for nodeID, addressPortPotitions := range nas.nodeAddressInfoMapByStatus[status] {
				for addressPort := range addressPortPotitions {
					*nodeAddresses = nas.append(*nodeAddresses, nas.nodeAddressInfoMapByID[nodeID][addressPort])
				}
			}
		}

	}
	return nil
}

func (nas *NodeAddressInfoStorage) GetAllItems(item interface{}) error {
	if !nas.isInTransaction {
		nas.RLock()
		defer nas.RUnlock()
	} else {
		nas.transactionalLock.RLock()
		defer nas.transactionalLock.RUnlock()
	}
	nodeAddresses, ok := item.(*[]*model.NodeAddressInfo)
	if !ok {
		return blocker.NewBlocker(blocker.ValidationErr, "WrongTypeItemExpected:*[]*model.NodeAddressInfo")
	}
	for _, nodeAddressInfos := range nas.nodeAddressInfoMapByID {
		for _, nodeAddressInfo := range nodeAddressInfos {
			*nodeAddresses = nas.append(*nodeAddresses, nodeAddressInfo)
		}
	}
	return nil
}

func (nas *NodeAddressInfoStorage) GetTotalItems() int {
	nas.RLock()
	defer nas.RUnlock()
	var totalItems int
	for _, nodeIDs := range nas.nodeAddressInfoMapByAddressPort {
		totalItems += len(nodeIDs)
	}
	return totalItems
}

func (nas *NodeAddressInfoStorage) RemoveItem(key interface{}) error {
	nas.Lock()
	defer nas.Unlock()
	if key != nil {
		storageKey, ok := key.(NodeAddressInfoStorageKey)
		if !ok {
			return blocker.NewBlocker(blocker.ValidationErr, "WrongTypeKeyExpected:NodeAddressInfoStorageKey")
		}
		// to remove node AddressInfo is require status and node ID
		if len(storageKey.Statuses) == 0 {
			return blocker.NewBlocker(blocker.ValidationErr, "StatusNodeAddressInfoRequired")
		}
		if storageKey.NodeID == 0 {
			return blocker.NewBlocker(blocker.ValidationErr, "NodeIDNodeAddressInfoRequired")
		}
		for _, status := range storageKey.Statuses {
			for fullAddressPort := range nas.nodeAddressInfoMapByStatus[status][storageKey.NodeID] {
				delete(nas.nodeAddressInfoMapByID[storageKey.NodeID], fullAddressPort)
				delete(nas.nodeAddressInfoMapByAddressPort[fullAddressPort], storageKey.NodeID)
				delete(nas.nodeAddressInfoMapByStatus[status][storageKey.NodeID], fullAddressPort)
			}
		}
	}

	if monitoring.IsMonitoringActive() {
		monitoring.SetCacheStorageMetrics(monitoring.TypeNodeAddressInfoCacheStorage, float64(nas.size()))
	}
	return nil
}

func (nas *NodeAddressInfoStorage) size() int {
	var (
		nasBytes bytes.Buffer
		enc      = gob.NewEncoder(&nasBytes)
	)
	_ = enc.Encode(nas.nodeAddressInfoMapByID)
	_ = enc.Encode(nas.nodeAddressInfoMapByAddressPort)
	_ = enc.Encode(nas.nodeAddressInfoMapByStatus)
	_ = enc.Encode(nas.transactionalRemovedNodeAddressInfoMapByID)
	return nasBytes.Len()
}
func (nas *NodeAddressInfoStorage) GetSize() int64 {
	nas.RLock()
	defer nas.RUnlock()
	return int64(nas.size())
}

func (nas *NodeAddressInfoStorage) ClearCache() error {
	nas.Lock()
	defer nas.Unlock()
	nas.nodeAddressInfoMapByID = make(map[int64]map[string]model.NodeAddressInfo)
	nas.nodeAddressInfoMapByAddressPort = make(map[string]map[int64]bool)
	nas.nodeAddressInfoMapByStatus = make(map[model.NodeAddressStatus]map[int64]map[string]bool)
	nas.transactionalRemovedNodeAddressInfoMapByID = make(map[int64]map[string]bool)
	if monitoring.IsMonitoringActive() {
		monitoring.SetCacheStorageMetrics(monitoring.TypeNodeAddressInfoCacheStorage, float64(nas.size()))
	}
	return nil
}

// Transactional implementation

// Begin prepare data to begin doing transactional change to the cache, this implementation
// will never return error
func (nas *NodeAddressInfoStorage) Begin() error {
	nas.Lock()
	nas.transactionalLock.Lock()
	defer nas.transactionalLock.Unlock()
	nas.isInTransaction = true
	nas.transactionalRemovedNodeAddressInfoMapByID = make(map[int64]map[string]bool)
	return nil
}

func (nas *NodeAddressInfoStorage) Commit() error {
	// make sure isInTransaction is true
	if !nas.isInTransaction {
		return blocker.NewBlocker(blocker.ValidationErr, "BeginIsRequired")
	}
	nas.transactionalLock.Lock()
	defer func() {
		nas.isInTransaction = false
		nas.Unlock()
		nas.transactionalLock.Unlock()
	}()
	// Remove all node address info on transactional remove list
	for nodeID, nodePositionsByAddressPort := range nas.transactionalRemovedNodeAddressInfoMapByID {
		for fullAddress := range nodePositionsByAddressPort {
			status := nas.nodeAddressInfoMapByID[nodeID][fullAddress].Status
			delete(nas.nodeAddressInfoMapByStatus[status][nodeID], fullAddress)
			delete(nas.nodeAddressInfoMapByAddressPort[fullAddress], nodeID)
			delete(nas.nodeAddressInfoMapByID[nodeID], fullAddress)
		}
	}
	nas.transactionalRemovedNodeAddressInfoMapByID = make(map[int64]map[string]bool)
	if monitoring.IsMonitoringActive() {
		monitoring.SetCacheStorageMetrics(monitoring.TypeNodeAddressInfoCacheStorage, float64(nas.size()))
	}
	return nil
}

func (nas *NodeAddressInfoStorage) Rollback() error {
	// make sure isInTransaction is true
	if !nas.isInTransaction {
		return blocker.NewBlocker(blocker.ValidationErr, "BeginIsRequired")
	}
	nas.transactionalLock.Lock()
	defer func() {
		nas.isInTransaction = false
		nas.Unlock()
		nas.transactionalLock.Unlock()
	}()
	nas.transactionalRemovedNodeAddressInfoMapByID = make(map[int64]map[string]bool)
	if monitoring.IsMonitoringActive() {
		monitoring.SetCacheStorageMetrics(monitoring.TypeNodeAddressInfoCacheStorage, float64(nas.size()))
	}
	return nil
}

// TxSetItem currently doesn’t need to set in transactional
func (nas *NodeAddressInfoStorage) TxSetItem(id, item interface{}) error {
	return nil
}

// TxSetItems currently doesn’t need to set in transactional
func (nas *NodeAddressInfoStorage) TxSetItems(items interface{}) error {
	return nil
}

// TxRemoveItem will add node address info ID into transactional remove list
func (nas *NodeAddressInfoStorage) TxRemoveItem(key interface{}) error {
	nas.transactionalLock.Lock()
	defer nas.transactionalLock.Unlock()
	storageKey, ok := key.(NodeAddressInfoStorageKey)
	if !ok {
		return blocker.NewBlocker(blocker.ValidationErr, "WrongTypeKeyExpected:NodeAddressInfoStorageKey")
	}
	// to remove node AddressInfo is require status and node ID
	if len(storageKey.Statuses) == 0 {
		return blocker.NewBlocker(blocker.ValidationErr, "StatusNodeAddressInfoRequired")
	}
	if storageKey.NodeID == 0 {
		return blocker.NewBlocker(blocker.ValidationErr, "NodeIDNodeAddressInfoRequired")
	}
	// add removed item into transactional remove list
	for _, status := range storageKey.Statuses {
		for fullAddressPort := range nas.nodeAddressInfoMapByStatus[status][storageKey.NodeID] {
			if nas.transactionalRemovedNodeAddressInfoMapByID[storageKey.NodeID] == nil {
				nas.transactionalRemovedNodeAddressInfoMapByID[storageKey.NodeID] = make(map[string]bool)
			}
			nas.transactionalRemovedNodeAddressInfoMapByID[storageKey.NodeID][fullAddressPort] = true
		}
	}
	if monitoring.IsMonitoringActive() {
		monitoring.SetCacheStorageMetrics(monitoring.TypeNodeAddressInfoCacheStorage, float64(nas.size()))
	}
	return nil
}

func (nas *NodeAddressInfoStorage) append(
	nodeAddresses []*model.NodeAddressInfo,
	nodeAddress model.NodeAddressInfo,
) []*model.NodeAddressInfo {
	var copyNodeAddress = nodeAddress
	copy(copyNodeAddress.BlockHash, nodeAddress.BlockHash)
	copy(copyNodeAddress.Signature, nodeAddress.Signature)
	return append(nodeAddresses, &copyNodeAddress)
}
