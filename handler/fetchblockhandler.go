package handler

import (
	"clmwallet/types"
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
)

const (
	serviceHost = "http://localhost:7545"
)

type FecthBlockHandler struct {
}

func InitFecthHandler() *FecthBlockHandler {
	return &FecthBlockHandler{}
}

func (f FecthBlockHandler) FecthBlocks(localLastBlockNumber int, blocks chan *types.BlockNode) {

	client, err := rpc.Dial("http://localhost:7545")
	if err != nil {
		fmt.Println("rpc.Dial err", err)
		return
	}

	var blockInfo = make(map[string]interface{})
	err = client.Call(&blockInfo, "eth_getBlockByNumber", "latest", "true")
	if err != nil {
		fmt.Println("client.Call error", err)
		return
	}

	f.ParseBlock(blockInfo)

}

/// 解析区块，并记录跟本平台相关的帐号地址
func (f FecthBlockHandler) ParseBlock(blockInfo map[string]interface{}) {
	if nil == blockInfo {
		trans := blockInfo["transcations"]
		if nil != trans {

		}
		return
	}

}
