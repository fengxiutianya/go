package main
/**
一个类型嵌入到另一个类型，以及内部尅性和外部类型之间的关系
**/
import(
	"fmt"
)

type user2 struct{
	name string
	email string
}

//notify 实现了一个可以通过user类型值的指针调用的方法
func (u *user2)notify(){
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,u.email)
}

type admin struct{
	user2  //嵌入类型   use user2 属于结构体
	level string
}

func main(){
	ad := admin{
		user2:user2{
			name:"jhon",
			email:"sadfsd",
		},
		level:"super",
	}
	ad.user2.notify()
	ad.notify()
}