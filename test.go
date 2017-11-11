package main

import (
	"fmt"
)

func main() {
	test := make([]int, 0)
	test = append(test, 12)
	test = append(test, 13)
	test = test[1:2]
	fmt.Println(test)
}
