package main
/**
	可以通过赋值给 _ 来忽略序号和值。
	如果只需要索引值，去掉“, value”的部分即可
**/
import "fmt"

func main(){
	pow := make([]int,10)
	for i := range pow{
		fmt.Println(uint(i))
		pow[i] = 1 << uint(i)  //   num << num1  num向左移动num1 位
	}
	for i,value := range pow{
		fmt.Printf("%d %d\n",i,value)
	}
}