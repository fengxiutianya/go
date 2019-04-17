package main

import (
	"fmt"
)

func main() {

	m1 := make(map[string]string)
	slice := make([]string, 3)
	slice = append(slice, "zhangke")
	slice = append(slice, "zhangke")
	slice = append(slice, "zhangke")
	m1["test"] = "test"
	if val, ok := m1["t"]; ok {
		fmt.Println(val)
	} else {
		m1["t"] = "add"
	}
	for key, _ := range m1 {
		fmt.Println(key)
	}
	for _, val1 := range slice {
		fmt.Println(val1)
	}

	mSlice := make(map[string][]string)
	mSlice["key"] = append(mSlice["key"], "zhangke", "zhank")
	key := "key"
	mSlice["key"] = append(mSlice[key], "zhanke")
	fmt.Println(mSlice)
}
