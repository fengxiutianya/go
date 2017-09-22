package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	//counter 是所有goruntine都要增加其值得变量
	counter int64

	//
	wg sync.WaitGroup
)

func incCounter(id int){
	defer wg.Done()

	for count :=0;count < 2;count++{
		atomic.AddInt64(&counter,1)
		runtime.Gosched()
	}
}

func main(){
	wg.Add(2)

	incCounter(1)
	incCounter(2)
	
	wg.Wait()
	fmt.Println("Final Counter : " ,counter)
}