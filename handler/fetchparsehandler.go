package handler

import (
	"clmwallet/types"
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
	"clmwallet/model"
)

const (
	serviceHost = "http://localhost:7545"
)

type FecthParseHandler struct {
	client *rpc.Client
}

func InitFecthParserHandler() *FecthParseHandler {
	hdl := &FecthParseHandler{}
	return  hdl
}



func (f FecthParseHandler) FecthBlocks(localLastBlockNumber int, blocks chan *types.BlockNodeFecthed) {

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

	f.parseBlock(blockInfo) //解析交易信息到本地的表格中（如果有）


}

/// 解析区块，并记录跟本平台相关的交易表中
func (f FecthParseHandler) parseBlock(blockInfo map[string]interface{}) {
	if nil == blockInfo { return }

	//得到交易信息数组
	transInfoI := blockInfo["transactions"]
	if nil == transInfoI {return }
	transInfo,ok := transInfoI.([]map[string]string)
	if !ok {
		return
	}

	//解析每一个交易
	// WARNING:伪代码
	for _,m := range transInfo {
		t := &model.TransactionInfo{}
		t.FromAddress = m["from"]
		t.ToAddress = m["to"]
		t.TxHash = m["transactionHash"]
		//...
		if model.ExistReceiveTrans(t.ToAddress) { //表示有人把该帐号转帐
			t.Save()
		}
	}
}


