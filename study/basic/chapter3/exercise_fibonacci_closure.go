package main

import "fmt"

//fibonacci 函数会返回一个返回int的函数

func fibonacci()func() int{
	 i := 0;
	 n := 1;
	return func() int{
		 j := n + i;
		 i = n;
		 n = j;
		 return j
	}
}

func main(){
	f := fibonacci()
	for i :=0;i < 10;i++{
		fmt.Println(f())
	}
}