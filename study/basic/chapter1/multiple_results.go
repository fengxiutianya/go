package main
import "fmt"
/**
	多值函数返回
**/
func swap(x,y string)(string,string){
	return y,x
}

func main(){
	a,b := swap("hello","world")  //初始化并赋值
	fmt.Println(a,b)
}