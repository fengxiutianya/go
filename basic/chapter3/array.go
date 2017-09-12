package main

import "fmt"
/**
类型 [n]T 是一个有n个类型为T的值得数组
表达式
	var a[10] int

**/
func main(){
	var a[2]string
	a[0] = "hello"
	a[1] =  "world"
	fmt.Println(a[0],a[1])
	fmt.Println(a)
}