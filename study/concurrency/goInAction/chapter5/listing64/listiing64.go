package main

import "fmt"
import "./counters"
func main(){
	counter := counters.New(10)
	fmt.Printf("COunter:%d\n",counter)

}