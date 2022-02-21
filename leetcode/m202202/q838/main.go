package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/push-dominoes/

838. 推多米诺
n 张多米诺骨牌排成一行，将每张多米诺骨牌垂直竖立。在开始时，同时把一些多米诺骨牌向左或向右推。
每过一秒，倒向左边的多米诺骨牌会推动其左侧相邻的多米诺骨牌。同样地，倒向右边的多米诺骨牌也会推动竖立在其右侧的相邻多米诺骨牌。
如果一张垂直竖立的多米诺骨牌的两侧同时有多米诺骨牌倒下时，由于受力平衡， 该骨牌仍然保持不变。
就这个问题而言，我们会认为一张正在倒下的多米诺骨牌不会对其它正在倒下或已经倒下的多米诺骨牌施加额外的力。

给你一个字符串 dominoes 表示这一行多米诺骨牌的初始状态，其中：
dominoes[i] = 'L'，表示第 i 张多米诺骨牌被推向左侧，
dominoes[i] = 'R'，表示第 i 张多米诺骨牌被推向右侧，
dominoes[i] = '.'，表示没有推动第 i 张多米诺骨牌。
返回表示最终状态的字符串。

示例 1：

输入：dominoes = "RR.L"
输出："RR.L"
解释：第一张多米诺骨牌没有给第二张施加额外的力。
示例 2：

输入：dominoes = ".L.R...LR..L.."
输出："LL.RR.LLRRLL.."

提示：

n == dominoes.length
1 <= n <= 105
dominoes[i] 为 'L'、'R' 或 '.'
*/
func main() {
	var tests = []struct {
		dominoes string
		want     string
	}{
		{
			dominoes: "RR.L",
			want:     "RR.L",
		},
		{
			dominoes: ".L.R...LR..L..",
			want:     "LL.RR.LLRRLL..",
		},
	}

	for _, item := range tests {
		if ans := pushDominoes(item.dominoes); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func pushDominoes(dominoes string) string {
	rs := make([]int, 0)
	ls := make([]int, 0)

	for i := 0; i < len(dominoes); i++ {
		if dominoes[i] == 'R' {
			rs = append(rs, i)
		} else if dominoes[i] == 'L' {
			ls = append(ls, i)
		}
	}

	for len(rs) > 0 || len(ls) > 0 {
		sb := []byte(dominoes)

		nrs := make([]int, 0)
		for i := 0; i < len(rs); i++ {
			ri := rs[i] + 1
			if ri == len(dominoes) || dominoes[ri] != '.' {
				continue
			} else {
				rii := ri + 1
				if rii == len(dominoes) {
					sb[ri] = 'R'
				} else if dominoes[rii] == '.' {
					sb[ri] = 'R'
					nrs = append(nrs, ri)
				} else if dominoes[rii] == 'R' {
					sb[ri] = 'R'
				}
			}
		}
		rs = nrs

		lrs := make([]int, 0)
		for i := 0; i < len(ls); i++ {
			li := ls[i] - 1
			if li == -1 || dominoes[li] != '.' {
				continue
			} else {
				lii := li - 1
				if lii == -1 {
					sb[li] = 'L'
				} else if dominoes[lii] == '.' {
					sb[li] = 'L'
					lrs = append(lrs, li)
				} else if dominoes[lii] == 'L' {
					sb[li] = 'L'
				}
			}
		}
		ls = lrs

		dominoes = string(sb)
	}

	return dominoes
}
