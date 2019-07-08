package handler

import (
	"context"

	"github.com/zoobc/zoobc-core/api/service"
	"github.com/zoobc/zoobc-core/common/chaintype"
	"github.com/zoobc/zoobc-core/common/model"
)

// BlockHandler to handle request related to Blocks from client
type BlockHandler struct {
	Service service.BlockServiceInterface // Use Blockservice Interface
}

// GetBlock handles request to get data of a single Block
func (bs *BlockHandler) GetBlock(ctx context.Context, req *model.GetBlockRequest) (*model.Block, error) {
	var blockResponse *model.Block
	var err error
	chainType := chaintype.GetChainType(req.ChainType)
	if req.ID != 0 {
		blockResponse, err = bs.Service.GetBlockByID(chainType, req.ID)
	}
	if req.BlockHeight != 0 {
		blockResponse, err = bs.Service.GetBlockByHeight(chainType, req.BlockHeight)
	}
	if err != nil {
		return nil, err
	}

	return blockResponse, nil
}

// GetBlocks handles request to get data of multiple blocks
func (bs *BlockHandler) GetBlocks(ctx context.Context, req *model.GetBlocksRequest) (*model.GetBlocksResponse, error) {
	chainType := chaintype.GetChainType(req.ChainType)
	blocksResponse, err := bs.Service.GetBlocks(chainType, req.BlockSize, req.BlockHeight)
	if err != nil {
		return nil, err
	}

	return blocksResponse, nil
}