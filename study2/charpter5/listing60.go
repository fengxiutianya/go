package main

import "fmt"

/**************************************
 *      Author : zhangke
 *      Date   : 2019-04-13 08:28
 *      email  : 398757724@qq.com
 *      Desc   : 外部接口定义同样的方法
 ***************************************/

type notifier interface {
	notify()
}
type user struct {
	name string
	email string
}

func (u *user)notify()  {
	fmt.Printf("sending user email to %s %s\n",u.name,u.email)
}

type admin struct {
	user
	level string
}

func (a *admin)notify()  {
	fmt.Printf("sending admin email to %s %s\n",a.name,a.email)
}

func main()  {
	ad :=admin{
		user:user{
			name:"jhon smith",
			email:"john@yahoo.com",
		},
		level:"super",
	}
	sendnotification(&ad)
	ad.user.notify()
	ad.notify()

}

func sendnotification( notifier2 notifier)  {
	notifier2.notify()
}