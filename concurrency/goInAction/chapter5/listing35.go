package main
/**
这个示例程序展示bytes.Buffer 也可以
用于io.Copy函数
**/

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main(){
	var b bytes.Buffer
	//将字符串写入Buffer
	b.Write([]byte("hello"))
	fmt.Fprintf(&b," World!\n")
	io.Copy(os.Stdout,&b)
}