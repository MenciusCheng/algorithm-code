package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/maximal-network-rank/

1615. 最大网络秩
中等
n 座城市和一些连接这些城市的道路 roads 共同组成一个基础设施网络。每个 roads[i] = [ai, bi] 都表示在城市 ai 和 bi 之间有一条双向道路。
两座不同城市构成的 城市对 的 网络秩 定义为：与这两座城市 直接 相连的道路总数。如果存在一条道路直接连接这两座城市，则这条道路只计算 一次 。
整个基础设施网络的 最大网络秩 是所有不同城市对中的 最大网络秩 。
给你整数 n 和数组 roads，返回整个基础设施网络的 最大网络秩 。

示例 1：
输入：n = 4, roads = [[0,1],[0,3],[1,2],[1,3]]
输出：4
解释：城市 0 和 1 的网络秩是 4，因为共有 4 条道路与城市 0 或 1 相连。位于 0 和 1 之间的道路只计算一次。
示例 2：
输入：n = 5, roads = [[0,1],[0,3],[1,2],[1,3],[2,3],[2,4]]
输出：5
解释：共有 5 条道路与城市 1 或 2 相连。
示例 3：

输入：n = 8, roads = [[0,1],[1,2],[2,3],[2,4],[5,6],[5,7]]
输出：5
解释：2 和 5 的网络秩为 5，注意并非所有的城市都需要连接起来。

提示：

2 <= n <= 100
0 <= roads.length <= n * (n - 1) / 2
roads[i].length == 2
0 <= ai, bi <= n-1
ai != bi
*/
func main() {
	var tests = []struct {
		n     int
		roads [][]int
		want  int
	}{
		{
			n:     4,
			roads: [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}},
			want:  4,
		},
		{
			n:     5,
			roads: [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}, {2, 3}, {2, 4}},
			want:  5,
		},
		{
			n:     8,
			roads: [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}, {5, 6}, {5, 7}},
			want:  5,
		},
	}

	for _, item := range tests {
		if ans := maximalNetworkRank(item.n, item.roads); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func maximalNetworkRank(n int, roads [][]int) int {
	cnt := make(map[int]map[int]bool)
	for _, road := range roads {
		if cnt[road[0]] == nil {
			cnt[road[0]] = make(map[int]bool)
		}
		cnt[road[0]][road[1]] = true
		if cnt[road[1]] == nil {
			cnt[road[1]] = make(map[int]bool)
		}
		cnt[road[1]][road[0]] = true
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			s := len(cnt[i]) + len(cnt[j])
			if cnt[i][j] {
				s--
			}
			res = max(res, s)
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
