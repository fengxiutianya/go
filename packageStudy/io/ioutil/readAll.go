package main

//读取文件的武安不内容
import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	r := strings.NewReader("Go is a general-purpose language designed with systems programing in mind.")
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(nil)
	}
	fmt.Printf("%s", b)
}
