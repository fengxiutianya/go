package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"github.com/goinaction/code/chapter3/words"
)

func main(){
	filename :=os.Args[1]
	counters,err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println(err)
		return
	}

	text := string(counters)
	count :=
}