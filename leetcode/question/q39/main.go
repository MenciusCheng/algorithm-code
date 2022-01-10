package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
https://leetcode-cn.com/problems/combination-sum/

39. 组合总和
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
对于给定的输入，保证和为 target 的不同组合数少于 150 个。

示例 1：

输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。
示例 2：

输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]]
示例 3：

输入: candidates = [2], target = 1
输出: []
示例 4：

输入: candidates = [1], target = 1
输出: [[1]]
示例 5：

输入: candidates = [1], target = 2
输出: [[1,1]]

提示：

1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都 互不相同
1 <= target <= 500
*/
func main() {
	var tests = []struct {
		candidates []int
		target     int
		want       [][]int
	}{
		{
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want: [][]int{
				{2, 2, 3}, {7},
			},
		},
		{
			candidates: []int{2, 3, 5},
			target:     8,
			want: [][]int{
				{2, 2, 2, 2}, {2, 3, 3}, {3, 5},
			},
		},
		{
			candidates: []int{2},
			target:     1,
			want:       [][]int{},
		},
		{
			candidates: []int{1},
			target:     1,
			want: [][]int{
				{1},
			},
		},
		{
			candidates: []int{1},
			target:     2,
			want: [][]int{
				{1, 1},
			},
		},
		{
			candidates: []int{2, 7, 6, 3, 5, 1},
			target:     9,
			want: [][]int{
				{1, 1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 2}, {1, 1, 1, 1, 1, 1, 3}, {1, 1, 1, 1, 1, 2, 2}, {1, 1, 1, 1, 2, 3}, {1, 1, 1, 1, 5}, {1, 1, 1, 2, 2, 2}, {1, 1, 1, 3, 3}, {1, 1, 1, 6}, {1, 1, 2, 2, 3}, {1, 1, 2, 5}, {1, 1, 7}, {1, 2, 2, 2, 2}, {1, 2, 3, 3}, {1, 2, 6}, {1, 3, 5}, {2, 2, 2, 3}, {2, 2, 5}, {2, 7}, {3, 3, 3}, {3, 6},
			},
		},
	}

	for _, item := range tests {
		if ans := combinationSum(item.candidates, item.target); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	for i, candidate := range candidates {
		if candidate == target {
			res = append(res, []int{candidate})
			break
		} else if candidate > target {
			break
		}
		if sub := combinationASum(candidate, candidates[i:], target, []int{candidate}); len(sub) > 0 {
			res = append(res, sub...)
		}
	}
	return res
}

func combinationASum(lSum int, candidates []int, target int, lArr []int) [][]int {
	res := make([][]int, 0)
	for i, candidate := range candidates {
		sum := lSum + candidate
		if sum == target {
			res = append(res, append(lArr, candidate))
			break
		} else if sum < target {
			if sub := combinationASum(sum, candidates[i:], target, copyAndAppend(lArr, candidate)); len(sub) > 0 {
				res = append(res, sub...)
			}
		} else {
			break
		}
	}
	return res
}

// 解决 append 会影响旧数组的陷阱
func copyAndAppend(arr []int, a int) []int {
	newArr := make([]int, 0, len(arr)+1)
	newArr = append(newArr, arr...)
	newArr = append(newArr, a)
	return newArr
}
