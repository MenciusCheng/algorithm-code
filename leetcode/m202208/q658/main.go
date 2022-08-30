package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
https://leetcode.cn/problems/find-k-closest-elements/

658. 找到 K 个最接近的元素
给定一个 排序好 的数组 arr ，两个整数 k 和 x ，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。
整数 a 比整数 b 更接近 x 需要满足：

|a - x| < |b - x| 或者
|a - x| == |b - x| 且 a < b

示例 1：

输入：arr = [1,2,3,4,5], k = 4, x = 3
输出：[1,2,3,4]
示例 2：

输入：arr = [1,2,3,4,5], k = 4, x = -1
输出：[1,2,3,4]

提示：

1 <= k <= arr.length
1 <= arr.length <= 10^4
arr 按 升序 排列
-10^4 <= arr[i], x <= 10^4
*/
func main() {
	var tests = []struct {
		arr  []int
		k    int
		x    int
		want []int
	}{
		{
			arr:  []int{1, 1, 1, 10, 10, 10},
			k:    1,
			x:    9,
			want: []int{10},
		},
		{
			arr:  []int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8},
			k:    3,
			x:    5,
			want: []int{3, 3, 4},
		},
		{
			arr:  []int{1, 2, 3, 4, 5},
			k:    4,
			x:    3,
			want: []int{1, 2, 3, 4},
		},
		{
			arr:  []int{1, 2, 3, 4, 5},
			k:    4,
			x:    -1,
			want: []int{1, 2, 3, 4},
		},
	}

	for _, item := range tests {
		if ans := findClosestElements(item.arr, item.k, item.x); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func findClosestElements(arr []int, k int, x int) []int {
	res := make([]int, 0, k)
	if k == 0 {
		return res
	}
	var l, r int
	for i, item := range arr {
		if x <= item {
			l = i - 1
			r = i
			break
		}
	}
	if l == r {
		l = len(arr) - 1
		r = l + 1
	}

	for len(res) < k {
		if l < 0 {
			res = append(res, arr[r])
			r++
		} else if r >= len(arr) {
			res = append(res, arr[l])
			l--
		} else {
			d1 := x - arr[l]
			d2 := arr[r] - x
			if d1 <= d2 {
				res = append(res, arr[l])
				l--
			} else {
				res = append(res, arr[r])
				r++
			}
		}
	}
	sort.Ints(res)
	return res
}
