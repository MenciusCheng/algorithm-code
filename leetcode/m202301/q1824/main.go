package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/minimum-sideway-jumps/

1824. 最少侧跳次数
中等
83
相关企业
给你一个长度为 n 的 3 跑道道路 ，它总共包含 n + 1 个 点 ，编号为 0 到 n 。一只青蛙从 0 号点第二条跑道 出发 ，它想要跳到点 n 处。然而道路上可能有一些障碍。

给你一个长度为 n + 1 的数组 obstacles ，其中 obstacles[i] （取值范围从 0 到 3）表示在点 i 处的 obstacles[i] 跑道上有一个障碍。如果 obstacles[i] == 0 ，那么点 i 处没有障碍。任何一个点的三条跑道中 最多有一个 障碍。

比方说，如果 obstacles[2] == 1 ，那么说明在点 2 处跑道 1 有障碍。
这只青蛙从点 i 跳到点 i + 1 且跑道不变的前提是点 i + 1 的同一跑道上没有障碍。为了躲避障碍，这只青蛙也可以在 同一个 点处 侧跳 到 另外一条 跑道（这两条跑道可以不相邻），但前提是跳过去的跑道该点处没有障碍。

比方说，这只青蛙可以从点 3 处的跑道 3 跳到点 3 处的跑道 1 。
这只青蛙从点 0 处跑道 2 出发，并想到达点 n 处的 任一跑道 ，请你返回 最少侧跳次数 。

注意：点 0 处和点 n 处的任一跑道都不会有障碍。
*/
func main() {
	var tests = []struct {
		obstacles []int
		want      int
	}{
		{
			obstacles: []int{0, 0, 3, 1, 0, 1, 0, 2, 3, 1, 0},
			want:      2,
		},
		{
			obstacles: []int{0, 1, 2, 3, 0},
			want:      2,
		},
		{
			obstacles: []int{0, 1, 1, 3, 3, 0},
			want:      0,
		},
		{
			obstacles: []int{0, 2, 1, 0, 3, 0},
			want:      2,
		},
	}

	for _, item := range tests {
		if ans := minSideJumps(item.obstacles); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func minSideJumps(obstacles []int) int {
	ds := []int{1, 0, 1}

	for i := 1; i < len(obstacles); i++ {
		ds2 := make([]int, 3)
		if obstacles[i] > 0 {
			ds2[obstacles[i]-1] = int(1e6)
			if obstacles[i] == 1 {
				ds2[1] = min(ds[1], ds[2]+1)
				ds2[2] = min(ds[2], ds[1]+1)
			} else if obstacles[i] == 2 {
				ds2[0] = min(ds[0], ds[2]+1)
				ds2[2] = min(ds[2], ds[0]+1)
			} else if obstacles[i] == 3 {
				ds2[1] = min(ds[1], ds[0]+1)
				ds2[0] = min(ds[0], ds[1]+1)
			}
		} else {
			ds2[0] = min3(ds[0], ds[1]+1, ds[2]+1)
			ds2[1] = min3(ds[1], ds[0]+1, ds[2]+1)
			ds2[2] = min3(ds[2], ds[0]+1, ds[1]+1)
		}
		ds = ds2
		//fmt.Printf("i:%d, ds: %+v\n", i, ds)
	}
	return min3(ds[0], ds[1], ds[2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	return min(min(a, b), c)
}
