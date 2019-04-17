package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**************************************
 *      Author : zhangke
 *      Date   : 2019-04-13 09:13
 *      email  : 398757724@qq.com
 *      Desc   : 竞争状态
 ***************************************/
var (
	counter int
	wg sync.WaitGroup
)

func main() {
	
	wg.Add(2)
	go inCounter(1)
	go inCounter(2)

	wg.Wait()
	fmt.Println("Final counter: ",counter)
}

func inCounter(it int)  {
	defer wg.Done()

	for count := 0;count < 2 ;count++  {
		value := counter
		// 当前goroutine从线程退出，并放回到队列
		runtime.Gosched()

		value++

		counter = value
	}
}