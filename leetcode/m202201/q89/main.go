package main

import "fmt"

/*
https://leetcode-cn.com/problems/gray-code/

89. 格雷编码
n 位格雷码序列 是一个由 2^n 个整数组成的序列，其中：
每个整数都在范围 [0, 2^n - 1] 内（含 0 和 2^n - 1）
第一个整数是 0
一个整数在序列中出现 不超过一次
每对 相邻 整数的二进制表示 恰好一位不同 ，且
第一个 和 最后一个 整数的二进制表示 恰好一位不同
给你一个整数 n ，返回任一有效的 n 位格雷码序列 。

示例 1：

输入：n = 2
输出：[0,1,3,2]
解释：
[0,1,3,2] 的二进制表示是 [00,01,11,10] 。
- 00 和 01 有一位不同
- 01 和 11 有一位不同
- 11 和 10 有一位不同
- 10 和 00 有一位不同
[0,2,3,1] 也是一个有效的格雷码序列，其二进制表示是 [00,10,11,01] 。
- 00 和 10 有一位不同
- 10 和 11 有一位不同
- 11 和 01 有一位不同
- 01 和 00 有一位不同
示例 2：

输入：n = 1
输出：[0,1]

提示：

1 <= n <= 16
*/
func main() {
	//var tests = []struct {
	//	n    int
	//	want []int
	//}{
	//	{},
	//}
	//
	//for _, item := range tests {
	//	if ans := grayCode(item.n); reflect.DeepEqual(ans, item.want) {
	//		fmt.Println(true)
	//	} else {
	//		fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
	//	}
	//}

	for i := 1; i <= 4; i++ {
		fmt.Println(grayCode(i))
	}
	//grayCode(2)
}

func grayCode(n int) []int {
	all := make([][]int, n+1)
	allI := make([]int, n+1)
	count := 0
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = '0'
	}
	pn := pow(n)
	for i := 0; i < pn; i++ {
		if i > 0 {
			count = bfIncr(bytes)
		}
		all[count] = append(all[count], i)
	}

	res := make([]int, 0, pn)
	curr := 0
	for i := 0; i < pn; i++ {
		res = append(res, all[curr][allI[curr]])
		allI[curr]++
		if curr+1 > n {
			curr--
		} else if allI[curr+1] >= len(all[curr+1]) {
			curr--
		} else {
			curr++
		}
	}

	return res
}

func pow(n int) int {
	sum := 1
	for i := 0; i < n; i++ {
		sum *= 2
	}
	return sum
}

func bfIncr(bytes []byte) int {
	count := 0
	added := false
	for i, b := range bytes {
		if added {
			if b == '1' {
				count++
			}
		} else {
			if b == '0' {
				bytes[i] = '1'
				count++
				added = true
			} else {
				bytes[i] = '0'
			}
		}
	}
	return count
}
