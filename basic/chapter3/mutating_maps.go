package main

import "fmt"
/**
修改map
在map m中插入或修改一个元素 ：
	 m[key] = elem
获得元素 ：
	elem = m[key]
删除元素
	delete(m,key)
通过双赋值检测某个键存在
	elem,ok = m[key]
如果key在m中，ok为true。否则ok为false，并且elem是map是map的元素类型的零值
同样的，当从map中读取某个不存在的键时，结果是map的元素类型的零值
**/
func main(){
	 m := make(map[string]int)

	 m["Answer"] = 42
	 fmt.Println("The value:",m["Answer"])

	 m["Answer"] = 48
	 fmt.Println("The value",m["ANswer"])

	 delete(m,"Answer")
	 fmt.Println("The value:",m["Answer"])

	 v,ok :=m["answer"]
	 fmt.Println("The value",v,"Present?",ok)
}