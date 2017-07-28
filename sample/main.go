package main
import (
	"log"
	"os"

	_ "./matchers"
	"./search"
)
//init 在main之前调用
func init(){
	//将日志输出到标准输出
	log.SetOutput(os.Stdout)
}
//main 是真个程序的入口
func main(){
	//使用特点的项做搜索
	search.Run("president")
}