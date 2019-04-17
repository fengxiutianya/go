package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte("appended some data\n")); err != nil {
		log.Fatal(err)
	}
	fmt.Println(12)
	con := make([]byte, 10)

	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Read(con)

	if err != nil {
		log.Fatal(err)
	}
	for ; 
	fmt.Println(string(con))
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
