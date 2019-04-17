package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**************************************
 *      Author : zhangke
 *      Date   : 2019-04-13 10:48
 *      email  : 398757724@qq.com
 *      Desc   :
 ***************************************/

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	court := make(chan int)

	wg.Add(2)
	go player("Nadal",court)
	go player("Djokovic",court)
	// 发球
	court <- 1
	wg.Wait()
}
func player(name string,court chan int)  {
	defer wg.Done()
	for  {
		ball,ok := <-court
		if !ok {
			fmt.Println("player " , name," won ")
			return
		}
		if n:=rand.Intn(100);n %13 ==0 {
			fmt.Println("player ",name," missed\n")
			close(court)
			return
		}

		fmt.Println("player ", name," hits ",ball)
		ball++
		court<- ball
	}
}