package main
/**
函数的闭包
go函数可以是闭包的，闭包是一个函数值，他来自函数体的外部的变量引用。
函数可以对这个引用值进行访问和赋值，换句话说这个函数被绑定在这个变量上

类如，函数adder返回一个闭包。每个闭包都被绑定到其各自的sum变量上
**/
import "fmt"

func adder()func(int)int{
	sum := 0
	return func(x int)int{
		sum +=x
		return sum
	}
}

func main(){
	pos,neg := adder(),adder()
	for i :=0;i < 10;i++{
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}