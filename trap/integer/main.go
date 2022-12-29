package main

import "fmt"

func main() {
	x := int(1e9 + 7)
	fmt.Printf("x=%d\n", x)
	fmt.Printf("?=1000000007\n")
	fmt.Println(x == 1000000007)
}
