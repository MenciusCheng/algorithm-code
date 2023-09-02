package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/maximum-enemy-forts-that-can-be-captured/

2511. 最多可以摧毁的敌人城堡数目
提示
简单
42
相关企业
给你一个长度为 n ，下标从 0 开始的整数数组 forts ，表示一些城堡。forts[i] 可以是 -1 ，0 或者 1 ，其中：

-1 表示第 i 个位置 没有 城堡。
0 表示第 i 个位置有一个 敌人 的城堡。
1 表示第 i 个位置有一个你控制的城堡。
现在，你需要决定，将你的军队从某个你控制的城堡位置 i 移动到一个空的位置 j ，满足：

0 <= i, j <= n - 1
军队经过的位置 只有 敌人的城堡。正式的，对于所有 min(i,j) < k < max(i,j) 的 k ，都满足 forts[k] == 0 。
当军队移动时，所有途中经过的敌人城堡都会被 摧毁 。

请你返回 最多 可以摧毁的敌人城堡数目。如果 无法 移动你的军队，或者没有你控制的城堡，请返回 0 。

示例 1：

输入：forts = [1,0,0,-1,0,0,0,0,1]
输出：4
解释：
- 将军队从位置 0 移动到位置 3 ，摧毁 2 个敌人城堡，位置分别在 1 和 2 。
- 将军队从位置 8 移动到位置 3 ，摧毁 4 个敌人城堡。
4 是最多可以摧毁的敌人城堡数目，所以我们返回 4 。
示例 2：

输入：forts = [0,0,1,-1]
输出：0
解释：由于无法摧毁敌人的城堡，所以返回 0 。

提示：

1 <= forts.length <= 1000
-1 <= forts[i] <= 1
*/
func main() {
	var tests = []struct {
		forts []int
		want  int
	}{
		{
			forts: []int{1, 0, 0, -1, 0, 0, 0, 0, 1},
			want:  4,
		},
		{
			forts: []int{0, 0, 1, -1},
			want:  0,
		},
	}

	for _, item := range tests {
		if ans := captureForts(item.forts); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func captureForts(forts []int) int {
	res := 0
	count := 0
	t := 0
	for _, fort := range forts {
		switch fort {
		case -1:
			if t == 1 {
				res = max(res, count)
			}
			t = fort
			count = 0
		case 0:
			if t == 1 || t == -1 {
				count++
			}
		case 1:
			if t == -1 {
				res = max(res, count)
			}
			t = fort
			count = 0
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
