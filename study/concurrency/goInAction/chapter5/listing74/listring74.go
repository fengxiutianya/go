package main

import (
	"fmt"
	"./entries"
)
func main(){
	//嵌入类型 提升到外部，嵌入类型的公开的也就变成了public类型，但不能用字面常量赋值
	ad := entries.Admin{
		Rights:"test",   
	}
	//可以在外部通过提升到外部字面常量来赋值
	ad.Name = "zhangke"
	ad.Email= "zhangke@email.com"
	fmt.Println(ad)

	u := entries.User8{
		Name :"zhangke",
		//小写的变量是私有的
	}
	fmt.Println(u)


}