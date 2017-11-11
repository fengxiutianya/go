package main

import "fmt"
/**
	在GO语言里，数组是一个值。变量名代表整个数组。
	因此，同样类型的数组可以复制给另一个数组。
**/

func main(){
	//声明一个数组
	var arr1 [5]int
	fmt.Println(arr1)
	//声明并初始化一个数组
	arr2 := [5]int{10,20,30,40,50}
	fmt.Println(arr2)
	//由容量来初始化数组  不指定数组长度，有后面的元素来确定长度，必须用这个符号 ...
	arr3 :=[...]int{10,20,30,40}

	fmt.Println(arr3)
	//声明数组，并指定特定元素的值
	arr4 :=[5]int{1:10,2:20}
	fmt.Println(arr4)
	//访问指针元素的值
	var arr5 = [5]*int{0:new(int),1:new(int)}
	*arr5[0] = 10;
	*arr5[1] = 20;
	fmt.Println(arr5)
	//把同样类型的一个数组复制给另外一个数组
	var arr6 [5]string
	arr7 := [5]string{"Red","Blue","Green","Yellow","Pink"}

	arr6 = arr7;
	fmt.Println(arr6,arr7)

	//复制数组指针，只会复制指针的值，不会复制指针所指向的值
	var  arr8 [3]*string
	arr9 := [3]*string{new(string),new(string),new(string)}
	arr8 = arr9;
	*arr9[0] = "Red"
	*arr9[1] = "Blue"
	*arr9[2] = "Green"
	fmt.Println(arr8,arr9)
}