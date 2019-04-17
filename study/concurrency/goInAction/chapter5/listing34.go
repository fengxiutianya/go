package main

import(
	"fmt"
	"io"
	"net/http"
	"os"
)

//init 在main函数之前调用
func init(){
	if len(os.Args) != 2{
		fmt.Println("Usage: ./example2 <url>")
		os.Exit(-1)
	}
}

//main是应用程序的入口
func main(){
	//从web服务器得到响应
	r,err :=http.Get(os.Args[1])
	if err != nil{
		fmt.Println(err)
		return
	}

	//从body复制到stdout
	io.Copy(os.Stdout,r.Body)
	if err := r.Body.Close();err != nil{
		fmt.Println(err)
	}
}