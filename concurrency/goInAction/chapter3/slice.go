package main

import "fmt"

/**
	切片的底层还是数组
	切片包含三个元素 指向底层数组的指针、切片访问的元素个数（长度）、切片允许增长到的元素个数（容量）
**/

func main(){
	//通过make来创建切片
	//创建一个字符串切片，其长度和容量都是5个元素
	slice := make([]string,5)
	fmt.Println(slice)
	//创建一个整型切片，其长度为3个元素，容量为5个元素
	slice1 := make([]int,3,5)
	fmt.Println(slice1)

	//通过切片字面量来声明切片
	//如果在【】运算符里指定了一个值，那么创建的就是数组而不是切片
	slice2 := []string{"Red","Blue","Green","yellow","Pink"}
	fmt.Println(slice2)

	//创建一个整型切片，其长度和容量都是3各元素
	slice3 := []int{10,20,30}
	fmt.Println(slice3)
	//使用空字符串初始化第100个元素
	slice4 :=[]string{99:""}
	fmt.Println(slice4)

	//创建3各元素的整型数组
	arr :=[3]int{10,20,30}
	//创建容量和长度都是3个元素的切片
	sli :=[]int{10,20,30}

	fmt.Println(arr,sli)


}