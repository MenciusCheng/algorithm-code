package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/number-of-boomerangs/description/

447. 回旋镖的数量
中等
相关标签
相关企业
给定平面上 n 对 互不相同 的点 points ，其中 points[i] = [xi, yi] 。回旋镖 是由点 (i, j, k) 表示的元组 ，其中 i 和 j 之间的距离和 i 和 k 之间的欧式距离相等（需要考虑元组的顺序）。
返回平面上所有回旋镖的数量。

示例 1：

输入：points = [[0,0],[1,0],[2,0]]
输出：2
解释：两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]
示例 2：

输入：points = [[1,1],[2,2],[3,3]]
输出：2
示例 3：

输入：points = [[1,1]]
输出：0

提示：

n == points.length
1 <= n <= 500
points[i].length == 2
-10^4 <= xi, yi <= 10^4
所有点都 互不相同
*/
func main() {
	var tests = []struct {
		points [][]int
		want   int
	}{
		{
			points: [][]int{{0, 0}, {1, 0}, {2, 0}},
			want:   2,
		},
		{
			points: [][]int{{1, 1}, {2, 2}, {3, 3}},
			want:   2,
		},
		{
			points: [][]int{{1, 1}},
			want:   0,
		},
	}

	for _, item := range tests {
		if ans := numberOfBoomerangs(item.points); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func numberOfBoomerangs(points [][]int) int {
	res := 0
	for i, p1 := range points {
		cnt := make(map[int]int)
		for j, p2 := range points {
			if i == j {
				continue
			}
			d := (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
			cnt[d]++
		}
		for _, v := range cnt {
			res += v * (v - 1)
		}
	}

	return res
}
