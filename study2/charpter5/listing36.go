package main

import "fmt"

/**************************************
 *      Author : zhangke
 *      Date   : 2019-04-12 14:13
 *      email  : 398757724@qq.com
 *      Desc   : 嵌入类型
 ***************************************/

type user struct {
	name string
	email string
}

func (u *user)notify()  {
	fmt.Printf("%s %s\n",u.name,u.email)
}

// 嵌入一个类型
type admin struct {
	user
	level string
}

func main() {
	ad := admin{user:user{"dd","ddd"},level:string("12")}
	ad.user.notify()
	ad.notify()
}