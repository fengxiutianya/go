package test

import (
	"fmt"
	_ "fmt"
	"gopkg.in/russross/blackfriday.v2"
	_ "gopkg.in/russross/blackfriday.v2"
	"testing"
)

/**************************************
 *      Author : zhangke
 *      Date   : 2019-04-14 01:24
 *      email  : 398757724@qq.com
 *      Desc   :
 ***************************************/

func Test_main(t *testing.T) {

	//bytes := []byte("go gopher")
	//output := blackfriday.Run(bytes,blackfriday.WithNoExtensions())
	//str := string(output)
	//fmt.Println(str)
			head := blackfriday.HeadingData{
				Level:        1,
				HeadingID:    "test",
				IsTitleblock: true,
			}
			fmt.Println(head)

//	node := blackfriday.Node{
//		Type:blackfriday.Paragraph,
//		Literal:[]byte("testsssss"),
//		HeadingData :blackfriday.HeadingData{
//			Level:        1,
//			HeadingID:    "test",
//			IsTitleblock: true,
//		},
//	}
//fmt.Println(node.HeadingData.Level)
}
