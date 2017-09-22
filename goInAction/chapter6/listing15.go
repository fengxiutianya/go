package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutDown int64


	wg sync.WaitGroup
)

func doWork(name string){
	defer wg.Done()
	
	for {
		fmt.Printf("Doing %s work\n",name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutDown) == 1{
			fmt.Printf("Shutting %s Down\n",name)
			break
		}
	}
}
func main(){
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(1 *time.Second)
	fmt.Println("Shutdown Now")
	
	atomic.StoreInt64(&shutDown,1)
	wg.Wait()
}