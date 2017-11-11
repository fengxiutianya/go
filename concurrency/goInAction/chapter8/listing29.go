package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//这个示例程序展示如何解码JSON字符串

//JSON包含要反序列化的样例字符串
var JSON1 = `{
	"name":"Gopher",
	"title":"programmer",
	"contact":{
			"home":"415.333.333.3333",
			"cell":"415.555.5555"
		}
	}`

func main() {
	//将JSON字符串序列化到map变量
	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON1), &c)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println("name: ", c["name"])
	fmt.Println("Title: ", c["title"])
	fmt.Println("Contact")
	fmt.Println("B: ", c["contact"].(map[string]interface{})["home"])
	fmt.Println("C: ", c["contact"].(map[string]interface{})["cell"])

}
