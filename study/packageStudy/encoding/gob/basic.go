package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

//This example shows the basic usage of the package:create an endocer
//transmit some values,receive them with a decoder

func main() {
	//Initialize the endocer and decoder. Normally enc and dec would be
	//bound to network connections and the encoder and decoder would
	//run in different processess.

	var network bytes.Buffer        //Stand-in for a nework connection
	enc := gob.NewEncoder(&network) //Will write to network
	dec := gob.NewDecoder(&network) //will read from network

	//Encode (send) some values
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	err = enc.Encode(P{1782, 1841, 1922, "Treehouse"})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	//Decode (receive) and print the values
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("%q:{%d, %d}\n", q.Name, *q.X, *q.Y)
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("%q:{%d,%d}\n", q.Name, *q.X, *q.Y)

}
