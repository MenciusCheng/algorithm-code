package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/minimum-height-trees/

310. 最小高度树
树是一个无向图，其中任何两个顶点只通过一条路径连接。 换句话说，一个任何没有简单环路的连通图都是一棵树。
给你一棵包含 n 个节点的树，标记为 0 到 n - 1 。给定数字 n 和一个有 n - 1 条无向边的 edges 列表（每一个边都是一对标签），其中 edges[i] = [ai, bi] 表示树中节点 ai 和 bi 之间存在一条无向边。
可选择树中任何一个节点作为根。当选择节点 x 作为根节点时，设结果树的高度为 h 。在所有可能的树中，具有最小高度的树（即，min(h)）被称为 最小高度树 。
请你找到所有的 最小高度树 并按 任意顺序 返回它们的根节点标签列表。
树的 高度 是指根节点和叶子节点之间最长向下路径上边的数量。

示例 1：

输入：n = 4, edges = [[1,0],[1,2],[1,3]]
输出：[1]
解释：如图所示，当根是标签为 1 的节点时，树的高度是 1 ，这是唯一的最小高度树。
示例 2：

输入：n = 6, edges = [[3,0],[3,1],[3,2],[3,4],[5,4]]
输出：[3,4]

提示：

1 <= n <= 2 * 10^4
edges.length == n - 1
0 <= ai, bi < n
ai != bi
所有 (ai, bi) 互不相同
给定的输入 保证 是一棵树，并且 不会有重复的边
*/
func main() {
	var tests = []struct {
		n     int
		edges [][]int
		want  []int
	}{
		{
			n:     4,
			edges: [][]int{{1, 0}, {1, 2}, {1, 3}},
			want:  []int{1},
		},
		{
			n:     6,
			edges: [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}},
			want:  []int{3, 4},
		},
		{
			n:     7,
			edges: [][]int{{0, 1}, {1, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 6}},
			want:  []int{1, 2},
		},
	}

	for _, item := range tests {
		if ans := findMinHeightTrees(item.n, item.edges); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

//func findMinHeightTrees3(n int, edges [][]int) []int {
//	if n == 1 {
//		return []int{0}
//	}
//
//	hes := make([][]int, n)
//	for _, item := range edges {
//		hes[item[0]] = append(hes[item[0]], item[1])
//		hes[item[1]] = append(hes[item[1]], item[0])
//	}
//
//	sp := make([]int, 0)
//	visited := make(map[int]bool)
//	for i := 0; i < n; i++ {
//		if len(hes[i]) == 1 {
//			sp = append(sp, i)
//			visited[i] = true
//		}
//	}
//	for len(visited) < n {
//		nsp := make([]int, 0)
//		for _, item := range sp {
//			for _, s := range hes[item] {
//				if !visited[s] {
//					nsp = append(nsp, s)
//					visited[s] = true
//				}
//			}
//		}
//		sp = nsp
//	}
//	return sp
//}
//
//func findMinHeightTrees2(n int, edges [][]int) []int {
//	hes := make([][]int, n)
//	for _, item := range edges {
//		hes[item[0]] = append(hes[item[0]], item[1])
//		hes[item[1]] = append(hes[item[1]], item[0])
//	}
//	res := make([]int, 0)
//	minDep := n
//	for i := 0; i < n; i++ {
//		dep := calDep(i, hes)
//		if dep < minDep {
//			minDep = dep
//			res = []int{i}
//		} else if dep == minDep {
//			res = append(res, i)
//		}
//	}
//	return res
//}
//
//func calDep(i int, hes [][]int) int {
//	dep := 0
//	sub := []int{i}
//	visited := make(map[int]bool)
//	visited[i] = true
//	for len(sub) > 0 {
//		nsub := make([]int, 0)
//		for _, item := range sub {
//			for _, h := range hes[item] {
//				if !visited[h] {
//					nsub = append(nsub, h)
//					visited[h] = true
//				}
//			}
//		}
//		sub = nsub
//		if len(sub) > 0 {
//			dep++
//		}
//	}
//	return dep
//}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	parents := make([]int, n)
	bfs := func(start int) (x int) {
		vis := make([]bool, n)
		vis[start] = true
		q := []int{start}
		for len(q) > 0 {
			x, q = q[0], q[1:]
			for _, y := range g[x] {
				if !vis[y] {
					vis[y] = true
					parents[y] = x
					q = append(q, y)
				}
			}
		}
		return
	}
	x := bfs(0) // 找到与节点 0 最远的节点 x
	y := bfs(x) // 找到与节点 x 最远的节点 y

	path := []int{}
	parents[x] = -1
	for y != -1 {
		path = append(path, y)
		y = parents[y]
	}
	m := len(path)
	if m%2 == 0 {
		return []int{path[m/2-1], path[m/2]}
	}
	return []int{path[m/2]}
}
