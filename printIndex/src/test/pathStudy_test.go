package test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

/**************************************
 *      Author : zhangke
 *      Date   : 2019-04-14 10:35
 *      email  : 398757724@qq.com
 *      Desc   :
 ***************************************/
func Test_getPath(t *testing.T) {
	path := "Users/zhangke/Documents/书籍"
	files,err := ioutil.ReadDir(path)
	if err != nil {
		t.Fatal(err)
	}
	// 获取所有的path
	//paths := []os.FileInfo
	for _, value := range files {
		if 	!strings.HasPrefix(strings.TrimSpace(value.Name()),".")  &&
			value.IsDir()	{
			fmt.Println(value.Name())
		}
	}
}