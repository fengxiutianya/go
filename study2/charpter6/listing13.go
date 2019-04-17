package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/**************************************
 *      Author : zhangke
 *      Date   : 2019-04-13 10:34
 *      email  : 398757724@qq.com
 *      Desc   : 带有原子变量
 ***************************************/

var(
	counter int64
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incounter(1)
	go incounter(2)

	wg.Wait()
	fmt.Println(counter)
}
func incounter(id int)  {
	defer wg.Done()
	for count := 0;count < 2 ;count++  {
		atomic.AddInt64(&counter,1)
		runtime.Gosched()
	}
}