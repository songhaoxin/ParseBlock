package main

import (
	"clmwallet/handler"
	"time"
)

func main() {
	hdl := handler.InitBlockHandler()
	hdl.FetchBlocksDelegate = handler.InitFethHandler()
	///for ; ;  {
		hdl.HandleTask()
	//}
	time.Sleep(1 * time.Second)


}
