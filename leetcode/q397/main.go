package main

import "fmt"

/*
https://leetcode-cn.com/problems/integer-replacement/

397. 整数替换

给定一个正整数 n ，你可以做如下操作：

如果 n 是偶数，则用 n / 2替换 n 。
如果 n 是奇数，则可以用 n + 1或n - 1替换 n 。
n 变为 1 所需的最小替换次数是多少？
*/

func main() {
	fmt.Println(integerReplacement(8) == 3)
	fmt.Println(integerReplacement(7) == 4)
	fmt.Println(integerReplacement(4) == 2)
}

func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}

	if n%2 == 0 {
		return 1 + integerReplacement(n/2)
	} else {
		a := integerReplacement(n - 1)
		b := integerReplacement(n + 1)

		if a <= b {
			return 1 + a
		} else {
			return 1 + b
		}
	}
}
