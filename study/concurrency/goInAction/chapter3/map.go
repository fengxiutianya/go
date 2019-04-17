package main

import "fmt"

func main(){
	//创建一个映射，键的类型是string，值得类型是int
	dict := make(map[string]int)
	fmt.Println(dict)
	//创建一个映射，检核值得类型都是string
	dict1 := map[string]string{"red":"#a333","Orange":"#e95a22"}
	
	fmt.Println(dict,dict1)
	//map的值是不能随意赋值的，必须map里面有这个键

	//获取Blue对应的值
	value,exists := dict1["red"]
	fmt.Println(value,exists)
	//在map中即使key不存在也会返回这个值对应的 nil值
	value1 := dict1["redd"]
	if value1 != ""{
		fmt.Println(value1)
	}
	fmt.Println()
	//使用delkete删除map中的一个键值对
	delete(dict1,"red")
	for value,index := range dict1{
		fmt.Println(value,index)
	}
}