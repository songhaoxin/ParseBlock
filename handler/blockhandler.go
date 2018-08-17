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
	hdl.blockPool = types.InitBlockPool()
	hdl.pendingPool = types.InitPendingPool()
	hdl.blockChan = make(chan *types.BlockNode, chanCap)
	return hdl
}
func (hdl *BlockHandler) HandleTask() {
	//接收区块
	if nil == hdl.FetchBlocksDelegate || nil == hdl.blockPool {
		return
	}

	go hdl.FetchBlocksDelegate.FetchBlocks(hdl.latestBlockNumber, hdl.blockChan)
	//更新区块池区块
	go hdl.blockPool.ReceiveBlocks(hdl.blockChan)
	//处理确认过的交易

	//处理需要重新处理的交易

}
