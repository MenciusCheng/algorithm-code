package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
https://leetcode.cn/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph/?envType=daily-question&envId=2024-04-04

2192. 有向无环图中一个节点的所有祖先
中等
给你一个正整数 n ，它表示一个 有向无环图 中节点的数目，节点编号为 0 到 n - 1 （包括两者）。
给你一个二维整数数组 edges ，其中 edges[i] = [fromi, toi] 表示图中一条从 fromi 到 toi 的单向边。
请你返回一个数组 answer，其中 answer[i]是第 i 个节点的所有 祖先 ，这些祖先节点 升序 排序。
如果 u 通过一系列边，能够到达 v ，那么我们称节点 u 是节点 v 的 祖先 节点。

示例 1：

输入：n = 8, edgeList = [[0,3],[0,4],[1,3],[2,4],[2,7],[3,5],[3,6],[3,7],[4,6]]
输出：[[],[],[],[0,1],[0,2],[0,1,3],[0,1,2,3,4],[0,1,2,3]]
解释：
上图为输入所对应的图。
- 节点 0 ，1 和 2 没有任何祖先。
- 节点 3 有 2 个祖先 0 和 1 。
- 节点 4 有 2 个祖先 0 和 2 。
- 节点 5 有 3 个祖先 0 ，1 和 3 。
- 节点 6 有 5 个祖先 0 ，1 ，2 ，3 和 4 。
- 节点 7 有 4 个祖先 0 ，1 ，2 和 3 。
示例 2：

输入：n = 5, edgeList = [[0,1],[0,2],[0,3],[0,4],[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
输出：[[],[0],[0,1],[0,1,2],[0,1,2,3]]
解释：
上图为输入所对应的图。
- 节点 0 没有任何祖先。
- 节点 1 有 1 个祖先 0 。
- 节点 2 有 2 个祖先 0 和 1 。
- 节点 3 有 3 个祖先 0 ，1 和 2 。
- 节点 4 有 4 个祖先 0 ，1 ，2 和 3 。

提示：

1 <= n <= 1000
0 <= edges.length <= min(2000, n * (n - 1) / 2)
edges[i].length == 2
0 <= fromi, toi <= n - 1
fromi != toi
图中不会有重边。
图是 有向 且 无环 的。
*/
func main() {
	var tests = []struct {
		n     int
		edges [][]int
		want  [][]int
	}{
		{
			n:     6,
			edges: [][]int{{0, 3}, {5, 0}, {2, 3}, {4, 3}, {5, 3}, {1, 3}, {2, 5}, {0, 1}, {4, 5}, {4, 2}, {4, 0}, {2, 1}, {5, 1}},
			want:  [][]int{{2, 4, 5}, {0, 2, 4, 5}, {4}, {0, 1, 2, 4, 5}, nil, {2, 4}},
		},
		{
			n:     8,
			edges: [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}},
			want:  [][]int{nil, nil, nil, {0, 1}, {0, 2}, {0, 1, 3}, {0, 1, 2, 3, 4}, {0, 1, 2, 3}},
		},
	}

	for _, item := range tests {
		if ans := getAncestors(item.n, item.edges); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func getAncestors(n int, edges [][]int) [][]int {
	inDegree := make([]int, n)
	chs := make([][]int, n)
	for _, edge := range edges {
		inDegree[edge[1]]++
		chs[edge[0]] = append(chs[edge[0]], edge[1])
	}

	ns := make([]int, 0)
	for i, d := range inDegree {
		if d == 0 {
			ns = append(ns, i)
		}
	}

	res := make([][]int, n)
	for len(ns) > 0 {
		i := ns[0]
		ns = ns[1:]
		vs := make(map[int]bool)
		for _, ch := range chs[i] {
			inDegree[ch]--
			if inDegree[ch] == 0 {
				ns = append(ns, ch)
			}

			sub := make([]int, 0)
			if !vs[ch] {
				sub = append(sub, ch)
			}
			vs[ch] = true
			for len(sub) > 0 {
				j := sub[0]
				sub = sub[1:]
				res[j] = append(res[j], i)
				for _, k := range chs[j] {
					if !vs[k] {
						sub = append(sub, k)
					}
					vs[k] = true
				}
			}
		}
	}
	for i := range res {
		sort.Ints(res[i])
	}

	return res
}
