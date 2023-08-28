package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/insert-interval/

57. 插入区间
中等
786
相关企业
给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

示例 1：

输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
输出：[[1,5],[6,9]]
示例 2：

输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出：[[1,2],[3,10],[12,16]]
解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
示例 3：

输入：intervals = [], newInterval = [5,7]
输出：[[5,7]]
示例 4：

输入：intervals = [[1,5]], newInterval = [2,3]
输出：[[1,5]]
示例 5：

输入：intervals = [[1,5]], newInterval = [2,7]
输出：[[1,7]]

提示：

0 <= intervals.length <= 10^4
intervals[i].length == 2
0 <= intervals[i][0] <= intervals[i][1] <= 10^5
intervals 根据 intervals[i][0] 按 升序 排列
newInterval.length == 2
0 <= newInterval[0] <= newInterval[1] <= 10^5
*/
func main() {
	var tests = []struct {
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{[][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
		{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{[][]int{}, []int{5, 7}, [][]int{{5, 7}}},
		{[][]int{{1, 5}}, []int{2, 3}, [][]int{{1, 5}}},
		{[][]int{{1, 5}}, []int{2, 7}, [][]int{{1, 7}}},
	}

	for _, item := range tests {
		if ans := insert(item.intervals, item.newInterval); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)

	if len(intervals) == 0 {
		res = append(res, newInterval)
		return res
	}

	for _, interval := range intervals {
		if newInterval == nil {
			res = append(res, interval)
			continue
		}
		if newInterval[0] <= interval[0] {
			if newInterval[1] >= interval[0] {
				newInterval[1] = max(newInterval[1], interval[1])
			} else {
				res = append(res, newInterval)
				newInterval = nil
				res = append(res, interval)
			}
		} else if newInterval[0] <= interval[1] {
			newInterval[0] = interval[0]
			newInterval[1] = max(newInterval[1], interval[1])
		} else {
			res = append(res, interval)
		}
	}

	if newInterval != nil {
		res = append(res, newInterval)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
