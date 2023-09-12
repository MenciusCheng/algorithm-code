package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/course-schedule-iv/

1462. 课程表 IV
中等

你总共需要上 numCourses 门课，课程编号依次为 0 到 numCourses-1 。你会得到一个数组 prerequisite ，其中 prerequisites[i] = [ai, bi] 表示如果你想选 bi 课程，你 必须 先选 ai 课程。

有的课会有直接的先修课程，比如如果想上课程 1 ，你必须先上课程 0 ，那么会以 [0,1] 数对的形式给出先修课程数对。
先决条件也可以是 间接 的。如果课程 a 是课程 b 的先决条件，课程 b 是课程 c 的先决条件，那么课程 a 就是课程 c 的先决条件。
你也得到一个数组 queries ，其中 queries[j] = [uj, vj]。对于第 j 个查询，您应该回答课程 uj 是否是课程 vj 的先决条件。
返回一个布尔数组 answer ，其中 answer[j] 是第 j 个查询的答案。

示例 1：

输入：numCourses = 2, prerequisites = [[1,0]], queries = [[0,1],[1,0]]
输出：[false,true]
解释：课程 0 不是课程 1 的先修课程，但课程 1 是课程 0 的先修课程。
示例 2：

输入：numCourses = 2, prerequisites = [], queries = [[1,0],[0,1]]
输出：[false,false]
解释：没有先修课程对，所以每门课程之间是独立的。
示例 3：

输入：numCourses = 3, prerequisites = [[1,2],[1,0],[2,0]], queries = [[1,0],[1,2]]
输出：[true,true]

提示：

2 <= numCourses <= 100
0 <= prerequisites.length <= (numCourses * (numCourses - 1) / 2)
prerequisites[i].length == 2
0 <= ai, bi <= n - 1
ai != bi
每一对 [ai, bi] 都 不同
先修课程图中没有环。
1 <= queries.length <= 10^4
0 <= ui, vi <= n - 1
ui != vi
*/
func main() {
	var tests = []struct {
		numCourses    int
		prerequisites [][]int
		queries       [][]int
		want          []bool
	}{
		{
			numCourses:    5,
			prerequisites: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			queries:       [][]int{{0, 4}, {4, 0}, {1, 3}, {3, 0}},
			want:          []bool{true, false, true, false},
		},
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			queries:       [][]int{{0, 1}, {1, 0}},
			want:          []bool{false, true},
		},
		{
			numCourses:    2,
			prerequisites: [][]int{},
			queries:       [][]int{{1, 0}, {0, 1}},
			want:          []bool{false, false},
		},
		{
			numCourses:    3,
			prerequisites: [][]int{{1, 2}, {1, 0}, {2, 0}},
			queries:       [][]int{{1, 0}, {1, 2}},
			want:          []bool{true, true},
		},
	}

	for _, item := range tests {
		if ans := checkIfPrerequisite(item.numCourses, item.prerequisites, item.queries); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	inDegree := make(map[int]int)
	edge := make(map[int][]int)
	for _, item := range prerequisites {
		edge[item[0]] = append(edge[item[0]], item[1])
		inDegree[item[1]]++
	}

	start := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			start = append(start, i)
		}
	}

	parent := make(map[int]map[int]bool)
	for len(start) > 0 {
		ns := make([]int, 0)
		for _, i := range start {
			for _, j := range edge[i] {
				p, ok := parent[j]
				if !ok {
					p = make(map[int]bool)
				}
				for k := range parent[i] {
					p[k] = true
				}
				p[i] = true

				parent[j] = p
				inDegree[j]--
				if inDegree[j] == 0 {
					ns = append(ns, j)
				}
			}
		}
		start = ns
	}

	res := make([]bool, 0, len(queries))
	for _, query := range queries {
		p, ok := parent[query[1]]
		if ok && p[query[0]] {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}

	return res
}
