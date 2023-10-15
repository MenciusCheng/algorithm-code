package main

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"reflect"
)

var json2 = jsoniter.ConfigCompatibleWithStandardLibrary

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

func UnmarshalPromotionV2Data(extension string) error {
	data := PromotionV2Data{}
	err := json.Unmarshal([]byte(extension), &data)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", data)
	return nil
}

func Unmarshal2PromotionV2Data(extension string) error {
	data := PromotionV2Data{}
	err := json2.Unmarshal([]byte(extension), &data)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", data)
	return nil
}

func Unmarshal2PromotionV2DataXX(extension string) error {
	data := PromotionV2Data{}
	err := json2.Unmarshal([]byte(extension), &data)
	if err != nil {
		return err
	}
	fmt.Printf("data: %+v\n", data)

	thirdData := &data
	extensionByte, err := json2.Marshal(thirdData)
	if err != nil {
		return err
	}
	fmt.Printf("extensionByte: %s\n", extensionByte)

	return nil
}

// 把 Map 序列化为 JSON，再反序列化为 Map ，再序列化为 JSON
// 测试序列化与反序列化过程中，哪些字段类型会被改变
func MapToJsonToMapToJson() {
	// https://developer.mozilla.org/zh-CN/docs/Glossary/JSON
	param := make(map[string]interface{})
	param["name"] = "cat"
	param["age"] = 12
	param["length"] = 15.78
	param["null"] = nil
	param["bool"] = true
	param["sub"] = map[string]interface{}{
		"subname": "wei",
		"age":     15,
	}
	param["ints"] = []int{1, 2}
	param["sp"] = "<div>cat&dog</div>" //  replaces "<", ">", "&", U+2028, and U+2029 are escaped to "\u003c","\u003e", "\u0026", "\u2028", and "\u2029".

	bytes, err := json.Marshal(param)
	if err != nil {
		panic(err)
	}

	unmarshalParam := make(map[string]interface{})
	err = json.Unmarshal(bytes, &unmarshalParam)
	if err != nil {
		panic(err)
	}

	unmarshalBytes, err := json.Marshal(unmarshalParam)
	if err != nil {
		panic(err)
	}

	fmt.Printf("param: %+v, age type: %s\n", param, reflect.TypeOf(param["age"]).String())
	fmt.Printf("bytes: %s\n", string(bytes))
	// json 反序列化后，整型变成了浮点数
	fmt.Printf("unmarshalParam: %+v, age type: %s\n", unmarshalParam, reflect.TypeOf(unmarshalParam["age"]).String())
	fmt.Printf("unmarshalBytes: %s\n", string(unmarshalBytes))
}
