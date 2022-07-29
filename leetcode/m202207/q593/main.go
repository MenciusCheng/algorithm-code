package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
https://leetcode.cn/problems/valid-square/

593. 有效的正方形
给定2D空间中四个点的坐标 p1, p2, p3 和 p4，如果这四个点构成一个正方形，则返回 true 。
点的坐标 pi 表示为 [xi, yi] 。输入 不是 按任何顺序给出的。
一个 有效的正方形 有四条等边和四个等角(90度角)。

示例 1:

输入: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,1]
输出: True
示例 2:

输入：p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,12]
输出：false
示例 3:

输入：p1 = [1,0], p2 = [-1,0], p3 = [0,1], p4 = [0,-1]
输出：true

提示:

p1.length == p2.length == p3.length == p4.length == 2
-10^4 <= xi, yi <= 10^4
*/
func main() {
	var tests = []struct {
		p1   []int
		p2   []int
		p3   []int
		p4   []int
		want bool
	}{
		{
			p1:   []int{0, 0},
			p2:   []int{1, 1},
			p3:   []int{1, 0},
			p4:   []int{0, 1},
			want: true,
		},
		{
			p1:   []int{0, 0},
			p2:   []int{1, 1},
			p3:   []int{1, 0},
			p4:   []int{0, 12},
			want: false,
		},
		{
			p1:   []int{1, 0},
			p2:   []int{-1, 0},
			p3:   []int{0, 1},
			p4:   []int{0, -1},
			want: true,
		},
	}

	for _, item := range tests {
		if ans := validSquare(item.p1, item.p2, item.p3, item.p4); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	ds1 := []int{calD(p1, p2), calD(p1, p3), calD(p1, p4)}
	sort.Ints(ds1)
	if !(ds1[0] == ds1[1] && ds1[2] > ds1[0]) {
		return false
	}
	ds2 := []int{calD(p2, p1), calD(p2, p3), calD(p2, p4)}
	sort.Ints(ds2)
	if !(ds2[0] == ds2[1] && ds2[2] > ds2[0]) {
		return false
	}
	ds3 := []int{calD(p3, p1), calD(p3, p2), calD(p3, p4)}
	sort.Ints(ds3)
	if !(ds3[0] == ds3[1] && ds3[2] > ds3[0]) {
		return false
	}
	ds4 := []int{calD(p4, p1), calD(p4, p2), calD(p4, p3)}
	sort.Ints(ds4)
	if !(ds4[0] == ds4[1] && ds4[2] > ds4[0]) {
		return false
	}

	return ds1[2] == ds2[2] && ds2[2] == ds3[2] && ds3[2] == ds4[2]
}

func calD(p1, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}
