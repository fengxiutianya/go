package main

import (
	"fmt"
)

//
//用于学习类型装换
//

func main(){
	basic()
	assert()
	typeSwitch()
}

//
//转换常量值
//转换变量值
//
func basic(){
	//将一个值x转换成特定类型，格式 T(x) 非常简单，类型加小括号即可
	//如果类型T以 *、<-、func(不带结果；列表),为避免造成歧义，需要类型括号包裹起来：(T)(x)
	//常量值转换成T 类型的值，需要满足下面的条件之一
	//1. x 可以表达为T的值
	//2. x 是浮点数值，T 是浮点类型
	//3. x 是一个整数而T是字符串类型

	fmt.Println(float32(1.23456))

	//对于一个常量值x 如果能转换成T类型的值，它需要满足下面的条件之一:
	// x可以赋值给 T
	// x的类型和T的底层类型 类型一致
	// x类型和T 都是未命名的指针类型，它们的指针指向的对象类型 类型一致
	// x的类型和T都是整数或者浮点数
	// x的类型和T都是复数
	// x是整数、slice of byte、slice of rune, T是字符串类型
	// x是字符串， T是slice of byte 或者slice of rune
}

//
//类型断言 type assertion
//
func assert(){
	// 和上节的类型转换不同，类型断言是将接口类型的值x，转换成类型T。
	//格式
	// x.(T)
	// v := x.(T)
	// v, ok := x.(T)

	var i interface{} = 10
	v ,ok := i.(int)
	if ok{
		fmt.Println(v)
	}
	
}

//
// 类型切换 type switch
//
func typeSwitch(){
	// 对于type switch，它检查的是值x的类型T是否匹配某个类型。
	// 格式如下，类型类型断言，但是括号内的不是某个具体的类型，而是单词type:
	//格式
	// switch x.(type) {
	//  cases
	// }
	var x interface{} 
	switch x.(type){
	case nil:
		fmt.Println("x is nil")
	case int:
		fmt.Println("x's type is int")
	default:
		fmt.Println("don't know the type")
	}

}