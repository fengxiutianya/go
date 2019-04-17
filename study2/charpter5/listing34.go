package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/**************************************
 *      Author : zhangke
 *      Date   : 2019-04-11 23:15
 *      email  : 398757724@qq.com
 *      Desc   : 实现简单的curl
 ***************************************/

func main()  {
	r,err := http.Get("http://www.baidu.com")
	if err != nil{
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout,r.Body)
	if err:=r.Body.Close();err!= nil {
		fmt.Println(err)
	}

}