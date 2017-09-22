package main
import "fmt"

//user 在程序里定义了一个用户类型
type user struct{
	name string
	email string
}
/**
func 和 函数名之间的参数被称为接收者
将函数与接收者的类型绑定在一起
如果一个函数有接收者，这个函数就被称为方法
**/
//nottify 使用值接收者实现了一个方法
func (u user)nottify(){
	fmt.Printf("Sending User Email TO %s<%s>\n",u.name,u.email)
}

//changemail 使用指针接收者实现了一个方法
func (u *user)changemail(email string){
	u.email = email
}

func main(){
	//user类型的值可以用来调用
	//使用值接收者声明的方法
	bill :=user{"bill","zhangke@qq.com"}
	bill.nottify();

	//指向user类型的指针也可以用来调用
	lisa := &user{"Lisa","lisa@qq.com"}
	lisa.nottify()
	
	//user类型的值也可以用来调用
	//使用值接收者声明的方法
	bill.changemail("zhangke@162.com")  //notify操作的也是一个副本，只不过这次操作的是从lisa指针指向的值得副本
	bill.nottify()

	//使用user类型值得指针可以用来调用
	//使用指针接收者声明的方法
	lisa.changemail("lisa@163.com")
	lisa.nottify()
}