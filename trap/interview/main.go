package main

import "fmt"

func main() {
	newAndMake()
}

// new 和 make 的区别
func newAndMake() {
	slice := make([]int, 0, 100)
	hash := make(map[int]bool, 10)
	ch := make(chan int, 5)

	fmt.Printf("%+v, %+v, %+v\n", slice, hash, ch)

	a := new(int)
	var i int
	a2 := &i
	fmt.Printf("%+v, %+v\n", a, a2)
}
