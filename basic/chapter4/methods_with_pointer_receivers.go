package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X,Y float64
}
/**
接收者为指针的方法
方法可以与命名类型或者命名类型的指针关联
刚刚看到的俩个Abs方法。一个是在 *Vertex指针类型上，而另一个在MyFloat值类型上。有
俩个原因需要使用指针接收者。首先避免在每个方法调用中拷贝值（如果值类型是大的
结构体的话会更有效率）。其次方法可以修改接收者指向的值

**/
func (v *Vertex) Scale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex)Abs()float64{
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
func main(){
	v := &Vertex{3,4}
	v.Scale(5)
	fmt.Print(v,v.Abs())
}