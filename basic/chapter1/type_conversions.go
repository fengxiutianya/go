package main

import (
	"fmt"
	"math"
)
/**
类型转换 
表达式 T(v) 将值转换为T

与 c不同的是GO的在不同类型之间的项目赋值是需要显示装换

**/
func main(){
	var x,y int =3,4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println(x,y,z)
}