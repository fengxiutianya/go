package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	constst()
}

func constst() {
	fmt.Println(rpc.DefaultDebugPath)
	fmt.Println(rpc.DefaultRPCPath)
	fmt.Println(rpc.DefaultServer)
	//获取一个默认的server
	var DefaultServer = rpc.NewServer()
	fmt.Println(DefaultServer)
	//获取默认的
}
