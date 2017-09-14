package main

import (
	"fmt"
	"time"
)
/**
错误
 GO程序使用error值来表示错误状态
 与 fmt.Stringer 类似 error 类型是一个内奸接口
 通常函数会返回一个error值，调用的他的代码应当判断这个错误是否等于nil，
 来进行错误处理
 i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
}
fmt.Println("Converted integer:", i)
error 为 nil 时表示成功；非 nil 的 error 表示错误。
**/

type MyError struct{
	When time.Time
	what string
}

func (e *MyError)Error()string{
	return fmt.Sprintf("at %v,%s",e.When,e.what)
}
//此时定义了一个方法
func run()error{
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main(){
	if err:=run();err != nil{
		fmt.Println(err)
	}
}