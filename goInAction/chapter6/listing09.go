package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	//counter是所有goruntine都要增加其值得变量
	counter int
	
	//wg是用来等待程序结束
	wg sync.WaitGroup

)

//incCounter 增加包里面的counter变量的值
func incCounter(id int){
	defer wg.Done()
	for count := 0;count < 2;count++{
		// 捕获counter的值
		value := counter
		// fmt.Println()
		//当goruntine从线程中退出，并放回到队列
		runtime.Gosched()

		//增加本地value变量的值
		value++

		//将该值保存回counter
		counter = value
	}
}
func main(){
		wg.Add(2)
	
	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final Counter: ",counter)
}