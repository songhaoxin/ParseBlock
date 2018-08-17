package main

import (
	"clmwallet/handler"
	"fmt"
)

func main() {
	hdl := handler.InitBlockHandler()
	for {
		hdl.HandleTask()
		fmt.Println("handing task ...")
	}
}
