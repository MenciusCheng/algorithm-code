package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/find-the-closest-palindrome/

564. 寻找最近的回文数
给定一个表示整数的字符串 n ，返回与它最近的回文整数（不包括自身）。如果不止一个，返回较小的那个。
“最近的”定义为两个整数差的绝对值最小。

示例 1:

输入: n = "123"
输出: "121"
示例 2:

输入: n = "1"
输出: "0"
解释: 0 和 2是最近的回文，但我们返回最小的，也就是 0。

提示:

1 <= n.length <= 18
n 只由数字组成
n 不含前导 0
n 代表在 [1, 10^18 - 1] 范围内的整数
*/
func main() {
	var tests = []struct {
		n    string
		want string
	}{
		{
			n:    "123",
			want: "121",
		},
		{
			n:    "1",
			want: "0",
		},
		{
			n:    "9",
			want: "8",
		},
		{
			n:    "11",
			want: "9",
		},
		{
			n:    "10",
			want: "9",
		},
		{
			n:    "1213",
			want: "1221",
		},
	}

	for _, item := range tests {
		if ans := nearestPalindromic(item.n); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func nearestPalindromic(n string) string {
	i := (len(n) - 1) / 2
	if isPalindromic(n) {
		a := palindr(add(n, i))
		b := minus(n, i)
		if len(b) < len(n) {
			sb := []byte(b)
			sb[i] = '9'
			b = string(sb)
		}
		b = palindr(b)
		if isLess(diff(a, n), diff(n, b)) {
			return a
		} else {
			return b
		}
	} else {
		a := palindr(n)
		if isLess(a, n) {
			b := palindr(add(n, i))
			if isLess(diff(b, n), diff(n, a)) {
				return b
			} else {
				return a
			}
		} else {
			b := minus(n, i)
			if len(b) < len(n) {
				sb := []byte(b)
				sb[i] = '9'
				b = string(sb)
			}
			b = palindr(b)
			if isLess(diff(a, n), diff(n, b)) {
				return a
			} else {
				return b
			}
		}
	}
}

func isLess(a, b string) bool {
	return len(a) < len(b) || len(a) == len(b) && a < b
}

func diff(as, bs string) string {
	a := []byte(as)
	b := []byte(bs)
	if isLess(as, bs) {
		a, b = b, a
	}
	sb := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		ai := len(a) - i - 1
		bi := len(b) - i - 1

		if a[ai] == '0'-1 {
			a[ai] = '9'
			a[ai-1]--
		}

		if bi >= 0 {
			if a[ai] >= b[bi] {
				sb[ai] = a[ai] - b[bi] + '0'
			} else {
				a[ai] += 10
				sb[ai] = a[ai] - b[bi] + '0'
				a[ai-1]--
			}
		} else {
			sb[ai] = a[ai]
		}
	}
	zero := -1
	for i := 0; i < len(sb); i++ {
		if sb[i] == '0' {
			zero = i
		} else {
			break
		}
	}
	if zero >= 0 {
		sb = sb[zero+1:]
	}

	return string(sb)
}

func isPalindromic(n string) bool {
	i := 0
	j := len(n) - 1
	for i < j {
		if n[i] != n[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func palindr(n string) string {
	sb := []byte(n)
	i := 0
	j := len(n) - 1
	for i < j {
		sb[j] = sb[i]
		i++
		j--
	}
	return string(sb)
}

func add(n string, index int) string {
	sb := []byte(n)
	if index < 0 {
		sb = append([]byte{'1'}, sb...)
	} else if sb[index] == '9' {
		sb[index] = '0'
		return add(string(sb), index-1)
	} else {
		sb[index]++
	}
	return string(sb)
}

func minus(n string, index int) string {
	sb := []byte(n)
	if sb[index] == '0' {
		sb[index] = '9'
		return minus(string(sb), index-1)
	} else if sb[index] == '1' && index == 0 && len(n) > 1 {
		sb = sb[1:]
	} else {
		sb[index]--
	}
	return string(sb)
}
