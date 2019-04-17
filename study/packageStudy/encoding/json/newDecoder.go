package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {

	test, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(test)
	for dec.More() {
		var m interface{}
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(m)
	}
}
