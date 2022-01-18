package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

/*
https://leetcode-cn.com/problems/minimum-time-difference/

539. 最小时间差
给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。

示例 1：

输入：timePoints = ["23:59","00:00"]
输出：1
示例 2：

输入：timePoints = ["00:00","23:59","00:00"]
输出：0

提示：

2 <= timePoints.length <= 2 * 10^4
timePoints[i] 格式为 "HH:MM"
*/
func main() {
	var tests = []struct {
		timePoints []string
		want       int
	}{
		{
			timePoints: []string{"23:59", "00:00"},
			want:       1,
		},
		{
			timePoints: []string{"00:00", "23:59", "00:00"},
			want:       0,
		},
		{
			timePoints: []string{"00:01", "00:02"},
			want:       1,
		},
	}

	for _, item := range tests {
		if ans := findMinDifference(item.timePoints); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func findMinDifference(timePoints []string) int {
	if len(timePoints) > MaxMinute {
		return 0
	}

	minutes := make([]int, 0, len(timePoints))
	for _, point := range timePoints {
		split := strings.Split(point, ":")
		h, _ := strconv.Atoi(split[0])
		m, _ := strconv.Atoi(split[1])
		minute := h*60 + m
		minutes = append(minutes, minute)
	}

	sort.Ints(minutes)
	min := MaxMinute
	for i := 0; i < len(minutes)-1; i++ {
		min = diff(minutes[i], minutes[i+1], min)
	}
	min = diff(minutes[len(minutes)-1], minutes[0]+MaxMinute, min)
	return min
}

const MaxMinute = 24 * 60

func diff(m1, m2, min int) int {
	d1 := m2 - m1
	if d1 < min {
		return d1
	} else {
		return min
	}
}
