package main

import (
	"fmt"
	"time"
)

//
//用于学习timer 定时器
//s
func main() {
	//newTimer()
	//afterFunc();
	//stop()
	reset()
}


//
// func NewTimer(d Duration) *Timer  在d时间之后创建一个timer对象
func newTimer() {

	// fmt.Println(time.Now())
	timer := time.NewTimer(time.Second * 1)

	fmt.Println(time.Now().Sub(<-timer.C))

}
//
//func AfterFunc(d Duration, f func()) *Timer
//在d时间之后调用f函数
func afterFunc(){
	current := time.Now();
	time.AfterFunc(time.Second * 1, func() {
		fmt.Println(time.Now().Sub(current))
	})
	time.Sleep(time.Second * 3)
}
//暂停还没有创建成功的timer
func stop(){
	//timer := time.NewTimer(time.Second * 10)
	timer :=time.AfterFunc(time.Second * 4, func() {
		fmt.Println(12)
	})
	if !timer.Stop(){
		fmt.Println(<-timer.C)
	}
	time.Sleep(time.Second * 10)
	//time.C 会等到channel里面的数据传送过来,
	//如果暂停成功，则C  channel就会关闭，所以就会读不出来
	//_,isClo := <-timer.C
	//fmt.Println(isClo)
}
//重新设置定时器
func reset(){
	current := time.Now()
	timer :=time.AfterFunc(time.Second * 4, func() {
		fmt.Println(time.Now().Sub(current))
	})
	time.Sleep(time.Second * 5)
	timer.Reset(time.Second * 3);
	time.Sleep(time.Second * 10)
}