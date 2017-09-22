package main

//这个示例程序展示如何用无缓冲的通道来模拟4个roroutine间的接力比赛
import (
	"fmt"
	"sync"
	"time"
)

var wg  sync.WaitGroup

func Runner(baton chan int){
	var newRunner int

	//等待接力棒
	runner := <-baton

	//开始围绕着跑到跑步
	fmt.Printf("Runner %d Runing with Baton\n ",runner)

	//创建下一位跑步者
	if runner != 4{
		newRunner = runner + 1
		fmt.Printf("Runner %d ro the line\n",newRunner)
		go Runner(baton)
	}

	//围绕跑到跑
	time.Sleep(100 * time.Millisecond)

	//比赛结束了吗
	if runner == 4{
		fmt.Printf("Runner %d Finished,Race Over\n",runner)
		wg.Done()
		return
	}

	//将接力棒交给下一位跑步者
	fmt.Printf("Runner %d Exchange with Runner %d\n",runner,newRunner)
	
	baton <- newRunner

}

func main(){
	baton := make(chan int)
	//为最后一个怕不这将计数加1  我需要等几个，数子就是几
	wg.Add(1)
	//第一位跑步者持有接力棒
	go Runner(baton)

	//开始比赛
	baton <- 1

	//等待比赛结束
	wg.Wait()
}