package main
/**
	这个示例程序显示goroutine 调度器是如何在单个线程上切分时间片
**/
import (
	"fmt"
	"runtime"
	"sync"
)

//wg用来等待时间的完成
var wg sync.WaitGroup

func main(){
	//分配一个逻辑处理器使用
	runtime.GOMAXPROCS(1)
	//计数加2，表示要等待俩个goroutine
	wg.Add(2)

	//创建俩个goroutine
	fmt.Println("Create Groutines")
	go printPrime("A");
	go printPrime("B")

	//等待goroutine结束
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("Terminating Program")

}

func printPrime(prefix string){
	//在函数退出时间调用Done来通知main函数工作已经完成
	defer wg.Done()
    next:
		for outer :=2;outer < 5000;outer++{
			for inner :=2;inner < outer;inner++{
				if outer % inner == 0{
					continue next
				}
			}
			fmt.Printf("%s:%d\n",prefix,outer)
		} 
		fmt.Println("completed",prefix)

}
