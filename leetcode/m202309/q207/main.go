package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/course-schedule/

207. 课程表
提示
中等
1.7K
相关企业
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

示例 1：

输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
示例 2：

输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。

提示：

1 <= numCourses <= 2000
0 <= prerequisites.length <= 5000
prerequisites[i].length == 2
0 <= ai, bi < numCourses
prerequisites[i] 中的所有课程对 互不相同
*/
func main() {
	var tests = []struct {
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          true,
		},
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			want:          false,
		},
	}

	for _, item := range tests {
		if ans := canFinish(item.numCourses, item.prerequisites); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	edge := make(map[int][]int)
	inDegree := make(map[int]int)
	for _, item := range prerequisites {
		inDegree[item[0]]++
		edge[item[1]] = append(edge[item[1]], item[0])
	}
	start := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 && len(edge[i]) > 0 {
			start = append(start, i)
		}
	}

	for len(start) > 0 {
		ns := make([]int, 0)
		for _, num := range start {
			for _, ch := range edge[num] {
				inDegree[ch]--
				if inDegree[ch] == 0 && len(edge[ch]) > 0 {
					ns = append(ns, ch)
				}
			}
		}
		start = ns
	}

	for _, v := range inDegree {
		if v > 0 {
			return false
		}
	}

	return true
}
