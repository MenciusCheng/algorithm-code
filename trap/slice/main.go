package main

import (
	"fmt"
	"strings"
)

func main() {
	splitArr()
}

func subSlice() {
	a := []int{0, 1, 2, 3, 4, 5}

	fmt.Println(a[0:1])
	fmt.Println(a[6:6])
	fmt.Println(a[5:6])
}

func sliceItemPointer() {
	arr := []int{1, 2, 3, 4, 5}
	res := make([]*int, 0)

	for _, item := range arr {
		fmt.Println("item=", item, " p=", &item)
		res = append(res, &item)
	}
	fmt.Println("res=", res)
}

func splitArr() {
	arr := strings.Split("", ",")
	fmt.Printf("arr: %+v, len: %d\n", arr, len(arr))
	arr = strings.Split("afwef", ",")
	fmt.Printf("arr: %+v, len: %d\n", arr, len(arr))
}

func findRepeat(arr []string) {
	m := make(map[string]bool)
	for i, str := range arr {
		if !m[str] {
			m[str] = true
		} else {
			fmt.Printf("%d) %s\n", i, str)
		}
	}
}
