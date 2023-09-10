package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/course-schedule-ii/

210. 课程表 II
提示
中等
849
相关企业
现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。

例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。

示例 1：

输入：numCourses = 2, prerequisites = [[1,0]]
输出：[0,1]
解释：总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。
示例 2：

输入：numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
输出：[0,2,1,3]
解释：总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
示例 3：

输入：numCourses = 1, prerequisites = []
输出：[0]

提示：
1 <= numCourses <= 2000
0 <= prerequisites.length <= numCourses * (numCourses - 1)
prerequisites[i].length == 2
0 <= ai, bi < numCourses
ai != bi
所有[ai, bi] 互不相同
*/
func main() {
	var tests = []struct {
		numCourses    int
		prerequisites [][]int
		want          []int
	}{
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          []int{0, 1},
		},
		{
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			want:          []int{0, 2, 1, 3},
		},
		{
			numCourses:    1,
			prerequisites: [][]int{},
			want:          []int{0},
		},
	}

	for _, item := range tests {
		if ans := findOrder(item.numCourses, item.prerequisites); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func findOrder(numCourses int, prerequisites [][]int) []int {

	inDegree := make(map[int]int)
	edge := make(map[int][]int)
	for _, item := range prerequisites {
		inDegree[item[0]]++
		edge[item[1]] = append(edge[item[1]], item[0])
	}

	start := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			start = append(start, i)
		}
	}

	res := make([]int, 0)
	for len(start) > 0 {
		ns := make([]int, 0)
		for _, i := range start {
			res = append(res, i)
			for _, j := range edge[i] {
				inDegree[j]--
				if inDegree[j] == 0 {
					ns = append(ns, j)
				}
			}
		}
		start = ns
	}

	for _, count := range inDegree {
		if count > 0 {
			return []int{}
		}
	}

	return res
}
