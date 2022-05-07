package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
https://leetcode-cn.com/problems/minimum-genetic-mutation/

433. 最小基因变化
基因序列可以表示为一条由 8 个字符组成的字符串，其中每个字符都是 'A'、'C'、'G' 和 'T' 之一。
假设我们需要调查从基因序列 start 变为 end 所发生的基因变化。一次基因变化就意味着这个基因序列中的一个字符发生了变化。
例如，"AACCGGTT" --> "AACCGGTA" 就是一次基因变化。
另有一个基因库 bank 记录了所有有效的基因变化，只有基因库中的基因才是有效的基因序列。
给你两个基因序列 start 和 end ，以及一个基因库 bank ，请你找出并返回能够使 start 变化为 end 所需的最少变化次数。如果无法完成此基因变化，返回 -1 。
注意：起始基因序列 start 默认是有效的，但是它并不一定会出现在基因库中。

示例 1：
输入：start = "AACCGGTT", end = "AACCGGTA", bank = ["AACCGGTA"]
输出：1

示例 2：
输入：start = "AACCGGTT", end = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA","AAACGGTA"]
输出：2

示例 3：
输入：start = "AAAAACCC", end = "AACCCCCC", bank = ["AAAACCCC","AAACCCCC","AACCCCCC"]
输出：3

提示：

start.length == 8
end.length == 8
0 <= bank.length <= 10
bank[i].length == 8
start、end 和 bank[i] 仅由字符 ['A', 'C', 'G', 'T'] 组成
*/
func main() {
	var tests = []struct {
		start string
		end   string
		bank  []string
		want  int
	}{
		{
			start: "AACCGGTT",
			end:   "AACCGGTA",
			bank:  []string{"AACCGGTA"},
			want:  1,
		},
		{
			start: "AACCGGTT",
			end:   "AAACGGTA",
			bank:  []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
			want:  2,
		},
		{
			start: "AAAAACCC",
			end:   "AACCCCCC",
			bank:  []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"},
			want:  3,
		},
	}

	for _, item := range tests {
		if ans := minMutation(item.start, item.end, item.bank); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func minMutation(start string, end string, bank []string) int {
	if start == end {
		return 0
	} else if len(bank) == 0 {
		return -1
	}

	res := math.MaxInt32
	for i := 0; i < len(bank); i++ {
		if canT(start, bank[i]) {
			nb := append([]string{}, bank[:i]...)
			nb = append(nb, bank[i+1:]...)
			c := minMutation(bank[i], end, nb)
			if c >= 0 {
				res = min(res, c+1)
			}
		}
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func canT(a, b string) bool {
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if diff >= 1 {
				return false
			}
			diff++
		}
	}
	return true
}
