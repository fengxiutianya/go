package main

import "fmt"

func main(){
  type user struct{
	  name string
	  email string
	  ext int
	  privileged bool
  }	

  lisa := user{
	  name :"lisa",
	  email:"398757724@qq.com",
	  ext: 123 ,
	  privileged :true,  //最后一行需要时也能 ，
  }

  zhangke :=user{"zhangke","@qq.com",134,false}  //最后不需要使用 ,
  fmt.Println(lisa,zhangke)

  type Duration int64
  var dur Duration
  // dur = int64(1000) //会出现编译错误，在go语言系统中，不认为这俩个类型是相同的
  fmt.Printf("%v\n",dur)
}