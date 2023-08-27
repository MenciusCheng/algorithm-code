package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
https://leetcode.cn/problems/merge-intervals/

56. 合并区间
中等
2.1K
相关企业
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

提示：

1 <= intervals.length <= 10^4
intervals[i].length == 2
0 <= starti <= endi <= 10^4
*/
func main() {
	var tests = []struct {
		intervals [][]int
		want      [][]int
	}{
		{
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			intervals: [][]int{{1, 4}, {4, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			intervals: [][]int{{1, 4}, {0, 4}},
			want:      [][]int{{0, 4}},
		},
	}

	for _, item := range tests {
		if ans := merge(item.intervals); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)

	el := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= el[1] {
			el[1] = max(el[1], intervals[i][1])
		} else {
			res = append(res, el)
			el = intervals[i]
		}
	}
	res = append(res, el)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
