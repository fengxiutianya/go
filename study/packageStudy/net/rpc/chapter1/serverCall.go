package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	// if len(os.Args) != 2 {
	// 	fmt.Println("Usage: ", os.Args[0], "server")
	// 	os.Exit(1)
	// }
	// serverAddress := os.Args[1]

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &Args{17, 8}

	var reply int
	//远程调用，args是传给远程函数的参数，reply用来接收函数的结果
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	//asynchronous call 异步调用
	quot := new(Quotient)
	//远程调用，args是传给远程函数的参数，reply用来接收函数的结果
	divCall := client.Go("Arith.Divide", args, quot, nil)
	// if err != nil {
	// 	log.Fatal("arith error:", err)
	// }
	replyCall := <-divCall.Done

	if replyCall.Error != nil {
		log.Fatal(replyCall.Error)
	}
	fmt.Println(replyCall.ServiceMethod)
	quot, ok := replyCall.Reply.(*Quotient)
	if !ok {
		log.Fatal(err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
