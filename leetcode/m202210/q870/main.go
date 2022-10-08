package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
https://leetcode.cn/problems/advantage-shuffle/

870. 优势洗牌
给定两个大小相等的数组 nums1 和 nums2，nums1 相对于 nums 的优势可以用满足 nums1[i] > nums2[i] 的索引 i 的数目来描述。
返回 nums1 的任意排列，使其相对于 nums2 的优势最大化。

示例 1：

输入：nums1 = [2,7,11,15], nums2 = [1,10,4,11]
输出：[2,11,7,15]
示例 2：

输入：nums1 = [12,24,8,32], nums2 = [13,25,32,11]
输出：[24,32,8,12]

提示：

1 <= nums1.length <= 10^5
nums2.length == nums1.length
0 <= nums1[i], nums2[i] <= 10^9
*/
func main() {
	var tests = []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			nums1: []int{5621, 1743, 5532, 3549, 9581},
			nums2: []int{913, 9787, 4121, 5039, 1481},
			want:  []int{1743, 9581, 5532, 5621, 3549},
		},
		{
			nums1: []int{2, 7, 11, 15},
			nums2: []int{1, 10, 4, 11},
			want:  []int{2, 11, 7, 15},
		},
		{
			nums1: []int{12, 24, 8, 32},
			nums2: []int{13, 25, 32, 11},
			want:  []int{24, 32, 8, 12},
		},
	}

	for _, item := range tests {
		if ans := advantageCount(item.nums1, item.nums2); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func advantageCount(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	sort.Ints(nums1)
	indexs := make([]int, len(nums2))
	for i := range indexs {
		indexs[i] = i
	}
	sort.Slice(indexs, func(i, j int) bool {
		return nums2[indexs[i]] < nums2[indexs[j]]
	})

	left := make([]int, 0, len(nums1))
	j := 0
	for _, i := range indexs {
		for j < len(nums1) && nums1[j] <= nums2[i] {
			left = append(left, nums1[j])
			j++
		}

		if j < len(nums1) {
			res[i] = nums1[j]
		} else {
			res[i] = left[len(left)-1]
			left = left[:len(left)-1]
		}
		j++
	}

	return res
}
