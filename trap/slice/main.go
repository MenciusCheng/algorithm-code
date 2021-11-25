package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	res := make([]*int, 0)

	for _, item := range arr {
		fmt.Println("item=", item, " p=", &item)
		res = append(res, &item)
	}
	fmt.Println("res=", res)
}
