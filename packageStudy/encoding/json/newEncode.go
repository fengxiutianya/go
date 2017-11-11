package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	test, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	enc := json.NewEncoder(test)
	var JSON = `{
		"name": "Gopher",
		"title": "programmer",
		"contact": {
			"home": "415.333.3333",
			"cell": "415.555.5555"  
		  }
		}`
	enc.Encode(JSON)
	test.Close()
}
