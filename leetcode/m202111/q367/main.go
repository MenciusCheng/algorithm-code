package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/valid-perfect-square/
func main() {
	fmt.Println(isPerfectSquare(4))
}

func isPerfectSquare(num int) bool {
	i := 1
	for true {
		s := i * i
		if s == num {
			return true
		} else if s > num {
			return false
		}
		i += 1
	}

	return false
}
