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

type RP struct {
	Name  string `json:"name"`
	BName string `json:"name"`
	Age   int    `json:"age"`
}

func main() {
	testRP()
}

func testRP() {
	//rp := RP{
	//	Name: "aa",
	//	Age:  10,
	//}
	//bs, _ := json.Marshal(rp)
	//fmt.Printf("bs: %s\n", bs)

	bs := []byte(" {\"name\":\"heihei\",\"age\":10}")

	rp2 := RP{}
	_ = json.Unmarshal(bs, &rp2)
	fmt.Printf("rp2: %+v\n", rp2)
}

func testApple() {
	a := Apple{
		Age: 2,
	}
	bs, _ := json.Marshal(a)
	fmt.Printf("%s\n", bs)
}
