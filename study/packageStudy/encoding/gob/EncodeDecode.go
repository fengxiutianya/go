package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

//the Vector type has unexported fields,which the package cannot access.
//we therefore write a BinaryMarshal/BinaryUnMarshal method pair to allow us
//to send and receive the type with the gob package.
//These interfaces are defined in the "encoding" package.
//we could equivalently use the locally defined GobEncode/GobDecoder interfaces
//

type Vector struct {
	x, y, z int
}

func (v Vector) MarshalBinary() ([]byte, error) {
	//A simpe encoding:plain text
	var b bytes.Buffer
	fmt.Fprintln(&b, v.x, v.y, v.z)
	return b.Bytes(), nil
}

//unmarshalBinary modifies the receiver so it must take a pointer receiver
func (v *Vector) UnmarshalBinary(data []byte) error {
	//a simple encoding :plain text
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &v.x, &v.y, &v.z)
	return err
}

//This example transmits a value that implements the custom encoding and decoding methods
func main() {
	var netWork bytes.Buffer //stand-in for the network
	enc := gob.NewEncoder(&netWork)
	err := enc.Encode(Vector{3, 4, 5})
	if err != nil {
		log.Fatal("endoce:", err)
	}

	//create a decoder and receive a value
	dec := gob.NewDecoder(&netWork)
	var v Vector
	err = dec.Decode(&v)
	if err != nil {
		log.Fatal("decdoer:", err)
	}
	fmt.Println(v)
}
