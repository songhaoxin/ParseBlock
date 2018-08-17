package handler

import "clmwallet/types"

type FethBlockHandler struct {
}

func InitFethHandler() *FethBlockHandler {
	return &FethBlockHandler{}
}

func (f FethBlockHandler) FetchBlocks(localLastBlockNumber int, blocks chan *types.BlockNode) (isFethed bool, chainLatestBlockNumber int) {
	for index := 0; index < 30; index++ {
		n := &types.BlockNode{index, "0x001"}
		blocks <- n
	}

	return false, 10
}
