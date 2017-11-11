package main

import (
	"log"
	"net"
)

func main() {
	//监听一个端口
	ln, err := net.Listen("tcp", "127.0.0.1:9000")

	if err != nil {
		log.Fatal(err)
	}
	//循环进行处理接收到的请求
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		go func(c net.Conn) {
			addr := c.RemoteAddr()
			log.Println("test" + addr.String())
			c.Write([]byte("zhangke"))
			c.Close()
		}(conn)
	}
}
