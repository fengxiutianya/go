package main
/**
示例程序展示的是内部类型和外部类型要实现同一个接口的做法
**/
import(
	"fmt"
)

type notifier1 interface{
	notify()
}

type user4 struct{
	name string
	email string
}
func (u *user4)notify(){
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,u.email)
}
type admin1 struct{
	user4
	level string
}
func (a *admin1)notify(){
	fmt.Printf("Sending admin email to %s<%s>\n",a.name,a.email)
}

func sendNotification(n notifier1){
	n.notify()
}
func main(){
	ad := admin1{
		user4:user4{
			name:"zhangke",
			email:"398757724@qq.com",
		},
		level:"super",
	}
	//给admin用户发一个通知，接口的嵌入的内部类型实现并没有提升的外部类型
	sendNotification(&ad)
	//我们可以直接访问内部类型的方法
	ad.user4.notify()
	//内部类型并没有提升到外部类型
	ad.notify()
}