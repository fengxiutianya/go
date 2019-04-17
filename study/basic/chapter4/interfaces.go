package main
import (
	"fmt"
	"math"
)
// 借口类型室友一组方法定义的集合
//借口类型的值可以存放实现这些方法的任何值
//注意  a =v 这行代码是一个错误。由于 Abs只定义在 *Vertex 指针类型上 ，
// 所以Vertex（值类型） 不满足 Abser

type Abser interface{
	Abs() float64
}
func main(){
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3,4}

	a = f   //a Myfloat 实现了Abser
	a = &v  //a *Vertex 实现了Abser

	//下面一行 v是一个Vertex（而不是*vertex）
	//所以没有实现 Abser
	// a = v

	fmt.Println(a.Abs())
}