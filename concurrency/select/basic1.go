package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

/**
select 当有一个case条件为真时就执行，如果有多个，就公平的从其中选择一个执行，
case条件只能是通道操作
**/

func receiveMsg() {
	c1 := make(chan string)
	c2 := make(chan string)
	t1 := time.Now()

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Millisecond * 150)
			c1 <- "msg 1 index " + strconv.Itoa(i)
		}
	}()
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Millisecond * 150)
			c2 <- "msg 2 index " + strconv.Itoa(i)
		}
	}()
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			log.Println(msg1)
		case msg2 := <-c2:
			log.Println(msg2)
		}
	}
	elapsed := time.Since(t1)
	log.Println("time elapsed", elapsed)
}
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func exeFibonacci() {
	c := make(chan int)
	quit := make(chan int)

	//produce data
	go func() {
		for i := 0; i < 10; i++ {
			log.Println("channel data item ", <-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
func main() {
	// receiveMsg()
	exeFibonacci()
}
