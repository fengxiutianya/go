package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	//访问一个地址
	conn, err := net.Dial("tcp", "127.0.0.1:9000")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	statu, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(statu)

}
