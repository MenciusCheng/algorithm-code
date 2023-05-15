package main

import "fmt"

func main() {
	fmtFloat64()
}

func fmtInt() {
	x := int(1e9 + 7)
	fmt.Printf("x=%d\n", x)
	fmt.Printf("?=1000000007\n")
	fmt.Println(x == 1000000007)
}

func fmtFloat64() {
	fmt.Printf("%.2f\n", float64(100.0))
	fmt.Printf("%.2f\n", float64(100.23))
	fmt.Printf("%v\n", float64(100.0))
	fmt.Printf("%v\n", float64(100.23))
}
