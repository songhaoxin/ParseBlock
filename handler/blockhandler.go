package handler

import (
	"clmwallet/types"
)

type BlockHandler struct {
	AffiremTransDelegate types.AffiremTransInterface
	ResendTransDelegate  types.ReSendTransInterface
	FetchBlocksDelegate  types.FecthBlocksInterface

	blockPool         *types.BlockPool
	pendingPool       *types.PendingPool
	blockChan         chan *types.BlockNode
	latestBlockNumber int
}

const chanCap = 100

func InitBlockHandler() *BlockHandler {
	hdl := &BlockHandler{}
	hdl.blockChan = make(chan *types.BlockNode, chanCap)
	return hdl
}
func (hdl *BlockHandler) HandleTask() {
	//接收区块
	if nil == hdl.FetchBlocksDelegate || nil == hdl.blockPool {
		return
	}
	/*
		_, latestBlockNumber := go hdl.FetchBlocksDelegate.FetchBlocks(hdl.blockPool.LatestNumber(), hdl.blockChan)
		if latestBlockNumber != -1 {
			hdl.latestBlockNumber = latestBlockNumber
		}
	*/
	go hdl.FetchBlocksDelegate.FetchBlocks(hdl.latestBlockNumber, hdl.blockChan)
	//更新区块池区块
	//处理确认过的交易
	//处理需要重新处理的交易

}
