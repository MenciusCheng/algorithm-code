package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/count-special-quadruplets/

1995. 统计特殊四元组
给你一个 下标从 0 开始 的整数数组 nums ，返回满足下述条件的 不同 四元组 (a, b, c, d) 的 数目 ：
nums[a] + nums[b] + nums[c] == nums[d] ，且
a < b < c < d

示例 1：

输入：nums = [1,2,3,6]
输出：1
解释：满足要求的唯一一个四元组是 (0, 1, 2, 3) 因为 1 + 2 + 3 == 6 。
示例 2：

输入：nums = [3,3,6,4,5]
输出：0
解释：[3,3,6,4,5] 中不存在满足要求的四元组。
示例 3：

输入：nums = [1,1,1,3,5]
输出：4
解释：满足要求的 4 个四元组如下：
- (0, 1, 2, 3): 1 + 1 + 1 == 3
- (0, 1, 3, 4): 1 + 1 + 3 == 5
- (0, 2, 3, 4): 1 + 1 + 3 == 5
- (1, 2, 3, 4): 1 + 1 + 3 == 5

提示：

4 <= nums.length <= 50
1 <= nums[i] <= 100
*/
func main() {
	var tests = []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 3, 6},
			want: 1,
		},
		{
			nums: []int{3, 3, 6, 4, 5},
			want: 0,
		},
		{
			nums: []int{1, 1, 1, 3, 5},
			want: 4,
		},
		{
			nums: []int{28, 8, 49, 85, 37, 90, 20, 8},
			want: 1,
		},
	}

	for _, item := range tests {
		if ans := countQuadruplets(item.nums); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func countQuadruplets(nums []int) int {
	var count int
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			for k := j + 1; k < len(nums)-1; k++ {
				sum := nums[i] + nums[j] + nums[k]
				for l := k + 1; l < len(nums); l++ {
					if nums[l] == sum {
						count++
					}
				}
			}
		}
	}

	return count
}
