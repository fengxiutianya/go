package main

import "fmt"
/**
变量在定义是没有明确的初始化时会赋值为零值
**/
func main(){
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n",i,f,b,s)
}