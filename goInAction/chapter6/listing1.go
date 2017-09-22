package main

import (
	"fmt"
	"sync"
	"runtime"
)
func main(){

	//分配一个逻辑处理器给调度器使用
	// runtime.GOMAXPROCS(1)
	
	//分配每个可用的核心分配一个逻辑处理器
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Start Goroutine")
	
	//wg表示等待程序完成的数量，计数加2表示要等待俩个
	var wg sync.WaitGroup
	wg.Add(2)
	//声明一个匿名函数，并创建一个goroutine
	go func(){
		//在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()
		for count := 0;count < 3;count++{
			for char :='a'; char < 'a'+26;char++{
				fmt.Printf("%c ",char)
			}
		}
	}()
	go func(){
		defer wg.Done()

		for count := 0; count < 3;count++{
			for char :='A';char < 'A'+26;char++{
				fmt.Printf("%c ",char)
			}
		}
	}()
	//等待goroutine结束
	fmt.Println("Waiting To FInish")
	wg.Wait()

	fmt.Println("\nTerminating Pogram")

}