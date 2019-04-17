package main

import (
	"fmt"
)
// notifier 是一个定义了通知类行为的接口

type notifier interface{
	notify()
}

type user3 struct{
	name string
	email string
}

func (u *user3) notify(){
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,u.email)
}

type admin struct{
	user3
	level string
}

func main(){
	ad  :=admin{
		user3:user3{
			name:"john smith",
			email:"john@yahoo.com",
		},
		level:"super",
	}
	sendNotification(&ad)
}

func sendNotification(n notifier){
	n.notify()
}