package main

import "clmwallet/modles/blocklist"

func main() {
	/*
		hdl := handler.InitBlockHandler()
		hdl.FetchBlocksDelegate = handler.InitFethHandler()
		///for ; ;  {
			hdl.HandleTask()
		//}
		time.Sleep(1 * time.Second)
	*/
	//fethHdl := handler.InitFecthParserHandler()
	//chains := make(chan *types.BlockNode, 100)
	//fethHdl.FecthBlocks(0

	node := &blocklist.BlockNodeInfo{}
	node.BlockNumber = "abc"
	node.BlockHash = "xxx"
	node.Store()
	print(node)

}
