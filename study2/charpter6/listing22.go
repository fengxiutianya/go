package main

import (
	"fmt"
	"sync"
	"time"
)

/**************************************
 *      Author : zhangke
 *      Date   : 2019-04-13 10:56
 *      email  : 398757724@qq.com
 *      Desc   :
 ***************************************/
var wg sync.WaitGroup

func main() {
  batton := make(chan int)

  wg.Add(1)

  go Runner(batton)
  batton <- 1
  wg.Wait()
}

func Runner(batton chan int)  {
	var newRunner int
	runner := <- batton

	fmt.Println(" Rnnner " ,runner," Running wit baton")

	if runner != 4{
		newRunner = runner+1
		fmt.Println(" Runner " ,newRunner," to the line")
		go Runner(batton)
	}

	time.Sleep(100 * time.Millisecond)

	if runner == 4{
		fmt.Println(" Runner ", runner," Finished，Race over")
		wg.Done()
		return
	}

	fmt.Println(" Runner ",runner," exchange with runner ",newRunner)
	batton <- newRunner
}