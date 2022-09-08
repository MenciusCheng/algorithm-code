package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
https://leetcode.cn/problems/maximum-length-of-pair-chain/

646. 最长数对链
给出 n 个数对。 在每一个数对中，第一个数字总是比第二个数字小。
现在，我们定义一种跟随关系，当且仅当 b < c 时，数对(c, d) 才可以跟在 (a, b) 后面。我们用这种形式来构造一个数对链。
给定一个数对集合，找出能够形成的最长数对链的长度。你不需要用到所有的数对，你可以以任何顺序选择其中的一些数对来构造。

示例：

输入：[[1,2], [2,3], [3,4]]
输出：2
解释：最长的数对链是 [1,2] -> [3,4]

提示：

给出数对的个数在 [1, 1000] 范围内。
*/
func main() {
	var tests = []struct {
		pairs [][]int
		want  int
	}{
		{
			pairs: [][]int{{1, 2}, {2, 3}, {3, 4}},
			want:  2,
		},
		{
			pairs: [][]int{{1, 2}},
			want:  1,
		},
		{
			pairs: [][]int{{-10, -8}, {8, 9}, {-5, 0}, {6, 10}, {-6, -4}, {1, 7}, {9, 10}, {-4, 7}},
			want:  4,
		},
	}

	for _, item := range tests {
		if ans := findLongestChain(item.pairs); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})

	m := make(map[[2]int]int)

	var f func(pz, pi int) int
	f = func(pz, pi int) int {
		if v, ok := m[[2]int{pz, pi}]; ok {
			return v
		} else {
			if pi >= len(pairs) {
				return 0
			}
			if pi == len(pairs)-1 {
				if pz < pairs[pi][0] {
					m[[2]int{pz, pi}] = 1
					return 1
				} else {
					return 0
				}
			}

			a1 := 0
			if pz < pairs[pi][0] {
				a1 = 1 + f(pairs[pi][1], pi+1)
			}
			a2 := f(pz, pi+1)
			m[[2]int{pz, pi}] = max(a1, a2)
			return max(a1, a2)
		}
	}
	f(-2000, 0)

	return m[[2]int{-2000, 0}]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
