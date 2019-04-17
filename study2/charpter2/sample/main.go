package main

import (
	_ "com.my/sample/matchers"
	"com.my/sample/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
