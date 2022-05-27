package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
https://leetcode.cn/problems/find-right-interval/

436. 寻找右区间
给你一个区间数组 intervals ，其中 intervals[i] = [starti, endi] ，且每个 starti 都 不同 。
区间 i 的 右侧区间 可以记作区间 j ，并满足 startj >= endi ，且 startj 最小化 。
返回一个由每个区间 i 的 右侧区间 的最小起始位置组成的数组。如果某个区间 i 不存在对应的 右侧区间 ，则下标 i 处的值设为 -1 。

示例 1：

输入：intervals = [[1,2]]
输出：[-1]
解释：集合中只有一个区间，所以输出-1。
示例 2：

输入：intervals = [[3,4],[2,3],[1,2]]
输出：[-1,0,1]
解释：对于 [3,4] ，没有满足条件的“右侧”区间。
对于 [2,3] ，区间[3,4]具有最小的“右”起点;
对于 [1,2] ，区间[2,3]具有最小的“右”起点。
示例 3：

输入：intervals = [[1,4],[2,3],[3,4]]
输出：[-1,2,-1]
解释：对于区间 [1,4] 和 [3,4] ，没有满足条件的“右侧”区间。
对于 [2,3] ，区间 [3,4] 有最小的“右”起点。

提示：

1 <= intervals.length <= 2 * 10^4
intervals[i].length == 2
-10^6 <= starti <= endi <= 10^6
每个间隔的起点都 不相同
*/
func main() {
	var tests = []struct {
		intervals [][]int
		want      []int
	}{
		{
			intervals: [][]int{{1, 2}},
			want:      []int{-1},
		},
		{
			intervals: [][]int{{3, 4}, {2, 3}, {1, 2}},
			want:      []int{-1, 0, 1},
		},
		{
			intervals: [][]int{{1, 4}, {2, 3}, {3, 4}},
			want:      []int{-1, 2, -1},
		},
	}

	for _, item := range tests {
		if ans := findRightInterval(item.intervals); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func findRightInterval(intervals [][]int) []int {
	starts := make([][]int, 0, len(intervals))
	ends := make([][]int, 0, len(intervals))
	for i, item := range intervals {
		starts = append(starts, []int{item[0], i})
		ends = append(ends, []int{item[1], i})
	}

	sort.Slice(starts, func(i, j int) bool {
		return starts[i][0] < starts[j][0]
	})
	sort.Slice(ends, func(i, j int) bool {
		return ends[i][0] < ends[j][0]
	})

	res := make([]int, len(intervals))
	for i := range res {
		res[i] = -1
	}

	j := 0
	for _, item := range ends {
		for ; j < len(starts); j++ {
			if starts[j][0] >= item[0] {
				res[item[1]] = starts[j][1]
				break
			}
		}
	}
	return res
}
