package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

type Binary struct {
	v int
}

func (b *Binary) String() string {
	return strconv.Itoa(b.v)
}

func main() {
	b := Binary{v: 18}
	//s := Stringer(b) // 编译报错
	s := Stringer(&b)
	fmt.Println(s.String())
}
