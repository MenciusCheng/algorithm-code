package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/permutations/

46. 全排列
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]

提示：

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同
*/
func main() {
	var tests = []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{1, 2, 3},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			nums: []int{0, 1},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			nums: []int{1},
			want: [][]int{{1}},
		},
	}

	for _, item := range tests {
		if ans := permute(item.nums); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	res := make([][]int, 0)
	for i, num := range nums {
		arr := make([]int, 0)
		arr = append(arr, nums[:i]...)
		arr = append(arr, nums[i+1:]...)
		arr2 := permute(arr)
		for j := range arr2 {
			arr2[j] = append(arr2[j], num)
		}
		res = append(res, arr2...)
	}

	return res
}
