package main
import "fmt"
/**
map 映射键到值
map 在使用之前必须用make 而不是new来创建，值位nil的map是空的，并且不能赋值
**/
type Vertex struct{
	Lat ,Long float64
}

var m map[string]Vertex

func main(){
	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		40.68433,-74.39967,
	}

	fmt.Println(m["Bell Labs"])
}