package main

import (
	"encoding/json"
	"fmt"
)

type Apple struct {
	Age int `json:"age,omitempty"`
	H   int `json:"h,omitempty"`
	D   Dog `json:"d,omitempty"`
}

type Dog struct {
	Name string `json:"name"`
}

func main() {

	a := Apple{
		Age: 2,
	}
	bs, _ := json.Marshal(a)
	fmt.Printf("%s\n", bs)
}
