package main

import (
	"clmwallet/handler"
	"clmwallet/types"
)

func main() {
	/*
		hdl := handler.InitBlockHandler()
		hdl.FetchBlocksDelegate = handler.InitFethHandler()
		///for ; ;  {
			hdl.HandleTask()
		//}
		time.Sleep(1 * time.Second)
	*/
	fethHdl := handler.InitFecthHandler()
	chains := make(chan *types.BlockNode, 100)
	fethHdl.FecthBlocks(0, chains)
}
