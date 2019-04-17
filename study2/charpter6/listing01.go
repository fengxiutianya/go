package charpter6

import (
	"fmt"
	"runtime"
	"sync"
)

/**************************************
 *      Author : zhangke
 *      Date   : 2019-04-13 09:01
 *      email  : 398757724@qq.com
 *      Desc   : 简单的创建一个打印大小字母和小写字母的goroutine程序
 ***************************************/

func main() {
	// 分配一个逻辑处理器给调度器使用
	//runtime.GOMAXPROCS(1)

	runtime.GOMAXPROCS(runtime.NumCPU())
	// wg用来等待程序完成，计数加2，表示要等待俩个goroutine
	var wg sync.WaitGroup
	wg.Add(2)
	
	fmt.Println("goroutines 开始")
	
	// 声明一个匿名函数，并创建一个goroutine
	go func() {
		defer wg.Done()

		for counter := 0;counter<3 ;counter++  {
			for char := 'A';char < 'A'+26 ;char++  {
				fmt.Printf("%c",char)
			}
			fmt.Println()
		}
	}()

	go func() {
		defer wg.Done()

		for counter := 0;counter<3 ;counter++  {
			for char := 'a';char < 'a'+26 ;char++  {
				fmt.Printf("%c",char)
			}
			fmt.Println()
		}
	}()

	// 等待goroutine结束
	fmt.Println("等待goroutine结束")
	wg.Wait()
	fmt.Println("结束")
}
