package main

import (
	"fmt"
)

//notifier 是一个定义了通知类型为的接口
type notifier interface{
	notify()
}

//user 在程序里定义了一个用户类型
type user1 struct{
	name string
	email string
}

//notify是使用指针接收者实现的方法
func (u *user1) notify(){
	fmt.Printf("Sending user email to %s<%s>\n",
				u.name,u.email)
}
func main(){
	//创建一个user类型的值，并发送通知

	u := user1{"bill","bill@email.com"}
	sendNotification(&u);
}
//sendNotification 接受了一个实现了notifier接口的值
//并发送通知
func sendNotification(n notifier){
	n.notify()
}
