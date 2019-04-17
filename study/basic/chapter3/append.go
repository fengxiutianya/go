package main
import  "fmt"
/**
 向slice添加元素
 向slice添加元素是一种常见的操作，因此GO 提供了一个內建函数 
   append 内建函数的文档对append、有详细介绍

   func append(s []T,vs ...T) []T
   append的第一个参数s是一个类型为T的数组，其余类型为T的值将会添加到slice
   append的结果是一个包含原slice所有元素加上新添加的元素的slice
   如果是的底层数组太小，而不能容纳所有值时，会分配一个更大的数组，
   返回的slice会指向这个新分配的数组
**/
func main(){
	 var a []int
	 printSlice("a",a)

	 //append works on nil slices
	 a = append(a,0)
	 printSlice("a",a)

	 //the slice grows as needed
	 a = append(a,1)
	 printSlice("a",a)

	 //we can add more than one element at a time
	 a = append(a,2,3,4)
	 printSlice("a",a)
}

func printSlice(s string,x []int){
	fmt.Printf("%s len=%d cap=%d %v\n",s,len(x),cap(x),x);
}