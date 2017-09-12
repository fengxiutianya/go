package main

import "fmt"
/**
	当俩个或多个函数参数类型一致时可以在最后一个参数哦后面标注类型
**/
func add(x,y int)int {
	return x + y;
}

func main(){
	fmt.Println(add(42,13))
}