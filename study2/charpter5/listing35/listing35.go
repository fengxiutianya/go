package main

import "fmt"

/**************************************
 *      Author : zhangke
 *      Date   : 2019-04-12 08:00
 *      email  : 398757724@qq.com
 *      Desc   : 简单的定义方法集
 ***************************************/

type user struct {
	name string
	email string
}

type notifier interface {
	notify()
}

func (u *user)notify()  {
	fmt.Println(u.email,u.name)
}

func sendNotifation(us notifier)  {
	us.notify()
}

func main()  {
	u := user{"12dd","333"}
	sendNotifation(&u)
}