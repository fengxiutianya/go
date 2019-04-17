package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

/**
 学习net包下面一些基础函数的使用
**/
func main() {

	// prinVar()
	// InterfaceAddrs()
	// Interfaces()
	// JoinHostPort()
	// LookupAddr()
	// LookupTXT()
	// splitAHostPort()
	// addrError()
	dial()
}

func prinVar() {
	//打印包常量
	fmt.Println(net.IPv6zero)
	fmt.Println(net.IPv4len)
}
func InterfaceAddrs() {

	//使用InterfaceAddrs
	addr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	for key, val := range addr {
		fmt.Println(key, val)
	}

}
func Interfaces() {
	rets, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}
	for key, val := range rets {
		fmt.Println(key, " : ", val)
	}
}

func JoinHostPort() {
	host1 := net.JoinHostPort("127.0.0.1", "90")
	fmt.Println(host1)
	//当ip地址里面有 ：时，此时就会把ip当成IPV6 地址
	host2 := net.JoinHostPort("127.0.0.1:1", "900")
	fmt.Println(host2)
}

func LookupAddr() {
	test, err := net.LookupAddr("127.0.0.1")
	print(err)
	for key, val := range test {
		fmt.Println(key, "  : ", val)
	}
}
func printErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
func LookupTXT() {
	txts, err := net.LookupTXT("127.0.0.1")
	printErr(err)
	for _, val := range txts {
		fmt.Println(val)
	}
}
func splitAHostPort() {
	host, port, err := net.SplitHostPort("127.0.0.1:80")
	printErr(err)
	fmt.Println(host)
	fmt.Println(port)
}
func addrError() {
	var addrE = net.AddrError{"错误", "127.0.0.1"}
	fmt.Println(addrE.Error())
	fmt.Println(addrE.Timeout())
	fmt.Println(addrE.Temporary())
}

func dial() {
	conn, err := net.DialTimeout("tcp", "golang.org:http", time.Duration(10)*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	addr := conn.LocalAddr()
	fmt.Println(addr.Network(), addr.String())
	addr = conn.RemoteAddr()
	fmt.Println(addr.Network(), addr.String())
	var bytes = make([]byte, 1024)

	for {
		_, err := conn.Read(bytes)
		if err != nil {
			break
		}
		fmt.Println(string(bytes))
	}
	conn.Close()
}
