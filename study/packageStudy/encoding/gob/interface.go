package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"math"
)

//
//This is example shows how to encode an interface value.
//The key distinction from reqular types types is to register the
//concrete type that implemetnts the interface
//
type Point struct {
	X, Y int
}

func (p Point) Hypotenuse() float64 {
	return math.Hypot(float64(p.X), float64(p.Y))
}

type Pythagoras interface {
	Hypotenuse() float64
}

func main() {
	var network bytes.Buffer //stand-in for the network
	//We must register the concrete type for the encoder and decoder(which would
	//normally be on a separate machine from the encoder).
	//On each end,this tells the engine which concrete type  is being sent that implements the interface.
	gob.Register(Point{})

	//Create an encoder and send some values
	enc := gob.NewEncoder(&network)
	for i := 0; i <= 3; i++ {
		interfaceEncode(enc, Point{3 * i, 4 * i})
	}

	//Create a decoder and receive some values.
	dec := gob.NewDecoder(&network)
	for i := 0; i <= 3; i++ {
		result := interfaceDecode(dec)
		fmt.Println(result.Hypotenuse())
	}
}

//interfaceEncode encodes the interface value into the encoder.
func interfaceEncode(enc *gob.Encoder, p Pythagoras) {
	//the encode will fail unless  the concrete type  has been
	//registered.We registered it in the calling function.

	//Pass pointer to interface so Encode sees(and hence sends) a value of interface type.
	//If we passed p directly it would see the concrete typoe instead.
	//See the blog post,"The laws of Reflection"for background

	err := enc.Encode(&p)
	if err != nil {
		log.Fatal("Encode:", err)
	}
}

//interfaceDecode decodes the next interface value from the stream and returns it.
func interfaceDecode(dec *gob.Decoder) Pythagoras {
	//The decode will fail unless the concrete type on the wire has been registered.
	//we registered it in the calling function
	var p Pythagoras
	err := dec.Decode(&p)
	if err != nil {
		log.Fatal("Decode:", err)
	}
	return p
}
