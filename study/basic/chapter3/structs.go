package main

import "fmt"
/**
结构体 一个机构体就是一个字段的集合
而type 的韩式是跟其字面意思相符
*/
type Vertex struct{
	X int
	Y int
}

func main(){
	fmt.Println(Vertex{1,2})
}