package main
/**
nil slice 
slice 的零值是 nil
一个nil的slice的长度和容量是0
**/
import "fmt"

func main(){
	var z []int
	fmt.Println(z,len(z),cap(z))

	if z == nil{
		fmt.Println("nil!");
	}
}
