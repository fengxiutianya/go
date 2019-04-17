package main
import "fmt"
/**
slice 可以重新切片 创建一个新的slice值指向相同的数组
表达式
  s[li:hi]  表示从lo 到hi-1 的slice元素，含俩端。因此
  s[lo:lo]  是空的而 s[lo:lo+1]有一个元素
**/
func main(){
	p := []int{2,3,5,7,11,13}
	fmt.Println("p == ",p);
	fmt.Println("p[1:4] == ",p[1:4])

	//省略下表代表从0开始
	fmt.Println("p[:3] == ",p[:3])

	//省略上标代表到len(s)结束
	fmt.Println("p[4:] == ",p[4:])
}