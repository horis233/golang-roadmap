package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/horis233/golang-roadmap/day1/prototext"
)

func main() {

	text := &prototext.Test{
		Name:   "Jiaming",
		Weight: []int32{120, 125, 198, 180, 150, 195},
		Height: 180,
		Motto:  "",
	}

	fmt.Println(text)
	// proto encoding
	data, err := proto.Marshal(text)
	if err != nil {
		fmt.Println("Encoding failed")
	}
	// Print after encoding
	fmt.Println(data)

	newtext := &prototext.Test{}

	// proto decoding
	err = proto.Unmarshal(data, newtext)
	if err != nil {
		fmt.Println("Decoding failed")
	}
	// Print after decoding
	fmt.Println(newtext)
	fmt.Println(newtext.String())

	fmt.Println("name:", newtext.Name)
	fmt.Println("height:", newtext.Height)
	fmt.Println("Weight:", newtext.Weight)
	fmt.Println("Motto:", newtext.Motto)

}
