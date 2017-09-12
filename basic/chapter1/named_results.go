package main
import "fmt"
/**
命名返回值 
go的返回值可以被命名，并且向变量那样使用
返回值的名称应当具有一定的意义，可以作为文档使用
没有参数的return语句返回结果的当前值，也就是直接返回
直接返回语句应当用在想下面这样的端函数中。在场函数中他们会影响代码的可读性
**/
func split(sum int)(x,y int){
	x = sum * 4 /9
	y = sum - x
	return 
}

func main(){
	// x,y := split(17)
	// fmt.Println(x,y)
	fmt.Println(split(17))
}