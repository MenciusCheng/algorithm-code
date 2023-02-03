package main

import "fmt"

func main() {
	initArray()
}

func initArray() {
	var arr1 [3]int
	arr2 := [3]int{}
	arr3 := [3]int{1, 2, 3}
	arr4 := [...]int{1, 2, 3}

	fmt.Printf("%+v, %+v, %+v, %+v", arr1, arr2, arr3, arr4)
}
