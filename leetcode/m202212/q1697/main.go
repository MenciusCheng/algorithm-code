package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
https://leetcode.cn/problems/checking-existence-of-edge-length-limited-paths/description/

1697. 检查边长度限制的路径是否存在
困难
给你一个 n 个点组成的无向图边集 edgeList ，其中 edgeList[i] = [ui, vi, disi] 表示点 ui 和点 vi 之间有一条长度为 disi 的边。请注意，两个点之间可能有 超过一条边 。
给你一个查询数组queries ，其中 queries[j] = [pj, qj, limitj] ，你的任务是对于每个查询 queries[j] ，判断是否存在从 pj 到 qj 的路径，且这条路径上的每一条边都 严格小于 limitj 。
请你返回一个 布尔数组 answer ，其中 answer.length == queries.length ，当 queries[j] 的查询结果为 true 时， answer 第 j 个值为 true ，否则为 false 。

示例 1：
输入：n = 3, edgeList = [[0,1,2],[1,2,4],[2,0,8],[1,0,16]], queries = [[0,1,2],[0,2,5]]
输出：[false,true]
解释：上图为给定的输入数据。注意到 0 和 1 之间有两条重边，分别为 2 和 16 。
对于第一个查询，0 和 1 之间没有小于 2 的边，所以我们返回 false 。
对于第二个查询，有一条路径（0 -> 1 -> 2）两条边都小于 5 ，所以这个查询我们返回 true 。

示例 2：
输入：n = 5, edgeList = [[0,1,10],[1,2,5],[2,3,9],[3,4,13]], queries = [[0,4,14],[1,4,13]]
输出：[true,false]
解释：上图为给定数据。

提示：

2 <= n <= 105
1 <= edgeList.length, queries.length <= 10^5
edgeList[i].length == 3
queries[j].length == 3
0 <= ui, vi, pj, qj <= n - 1
ui != vi
pj != qj
1 <= disi, limitj <= 10^9
两个点之间可能有 多条 边。
*/
func main() {
	var tests = []struct {
		n        int
		edgeList [][]int
		queries  [][]int
		want     []bool
	}{
		{
			n:        3,
			edgeList: [][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}},
			queries:  [][]int{{0, 1, 2}, {0, 2, 5}},
			want:     []bool{false, true},
		},
		{
			n:        5,
			edgeList: [][]int{{0, 1, 10}, {1, 2, 5}, {2, 3, 9}, {3, 4, 13}},
			queries:  [][]int{{0, 4, 14}, {1, 4, 13}},
			want:     []bool{true, false},
		},
	}

	for _, item := range tests {
		if ans := distanceLimitedPathsExist(item.n, item.edgeList, item.queries); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	for i := range queries {
		queries[i] = append(queries[i], i)
	}
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][2] < queries[j][2]
	})

	dsu := NewDsu(n)
	res := make([]bool, len(queries))

	k := 0
	for _, q := range queries {
		for ; k < len(edgeList) && edgeList[k][2] < q[2]; k++ {
			dsu.Union(edgeList[k][0], edgeList[k][1])
		}
		res[q[3]] = dsu.Find(q[0]) == dsu.Find(q[1])
	}
	return res
}

// 并查集 https://oi-wiki.org/ds/dsu/
type Dsu struct {
	pa []int
}

func NewDsu(n int) *Dsu {
	d := &Dsu{pa: make([]int, n)}
	for i := range d.pa {
		d.pa[i] = i
	}
	return d
}

func (d *Dsu) Find(x int) int {
	if d.pa[x] != x {
		d.pa[x] = d.Find(d.pa[x])
	}
	return d.pa[x]
}

func (d *Dsu) Union(x, y int) {
	d.pa[d.Find(x)] = d.Find(y)
}
