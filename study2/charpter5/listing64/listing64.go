package main

/**************************************
 *      Author : zhangke
 *      Date   : 2019-04-13 08:37
 *      email  : 398757724@qq.com
 *      Desc   :
 ***************************************/
import (
	"com.my/charpter5/listing64/counters"
	_ "com.my/charpter5/listing64/counters"
	"fmt"
)
func main() {
	// 下面这个操作需要俩个理由：
	// 1. 公开或未公开的标识符，不是一个值。
	// 2. 短声明操作符，有能力捕获引用的类型，并创建一个未公开的类型的变量
	//    永远不能显示穿件一个未公开的饿类型的变量，不过段变量声明操作符可以这么做
   alert :=counters.New(12)
   fmt.Println(alert)
}