package main

import "fmt"
/**
一个slice 会纸箱一个序列的值，并且包含了长度信息
[]T 是一个类型元素为T的slice
**/
func main(){
	p :=[]int{2,3,5,6,11,13}
	fmt.Println("p==",p);
	for i :=0;i< len(p);i++{
		fmt.Printf("p[%d] == %d\n",i,p[i])
	}
}