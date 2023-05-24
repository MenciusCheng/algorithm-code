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
	Age  int    `json:"age"`
}

type RP struct {
	Name  string `json:"name"`
	BName string `json:"name"`
	Age   int    `json:"age"`
}

func main() {
	structToJsonToMap()
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

func mapToJsonToStruct() {
	param := make(map[string]interface{})
	param["name"] = "Good Dog"
	param["age"] = 12

	bytes, err := json.Marshal(param)
	if err != nil {
		panic(err)
	}

	dog := Dog{}
	err = json.Unmarshal(bytes, &dog)
	if err != nil {
		panic(err)
	}
	fmt.Printf("dog: %+v\n", dog)
}

func structToJsonToMap() {
	dog := Dog{
		Name: "Good Dog",
		Age:  12,
	}

	bytes, err := json.Marshal(dog)
	if err != nil {
		panic(err)
	}

	param := make(map[string]interface{})
	err = json.Unmarshal(bytes, &param)
	if err != nil {
		panic(err)
	}
	fmt.Printf("param: %+v\n", param)
}
