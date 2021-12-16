package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

/*
https://leetcode-cn.com/problems/maximum-number-of-visible-points/

1610. 可见点的最大数目
难度：困难
给你一个点数组 points 和一个表示角度的整数 angle ，你的位置是 location ，其中 location = [posx, posy] 且 points[i] = [xi, yi] 都表示 X-Y 平面上的整数坐标。
最开始，你面向东方进行观测。你 不能 进行移动改变位置，但可以通过 自转 调整观测角度。换句话说，posx 和 posy 不能改变。你的视野范围的角度用 angle 表示， 这决定了你观测任意方向时可以多宽。设 d 为你逆时针自转旋转的度数，那么你的视野就是角度范围 [d - angle/2, d + angle/2] 所指示的那片区域。
对于每个点，如果由该点、你的位置以及从你的位置直接向东的方向形成的角度 位于你的视野中 ，那么你就可以看到它。
同一个坐标上可以有多个点。你所在的位置也可能存在一些点，但不管你的怎么旋转，总是可以看到这些点。同时，点不会阻碍你看到其他点。
返回你能看到的点的最大数目。

示例 1：

输入：points = [[2,1],[2,2],[3,3]], angle = 90, location = [1,1]
输出：3
解释：阴影区域代表你的视野。在你的视野中，所有的点都清晰可见，尽管 [2,2] 和 [3,3]在同一条直线上，你仍然可以看到 [3,3] 。
示例 2：

输入：points = [[2,1],[2,2],[3,4],[1,1]], angle = 90, location = [1,1]
输出：4
解释：在你的视野中，所有的点都清晰可见，包括你所在位置的那个点。
示例 3：

输入：points = [[1,0],[2,1]], angle = 13, location = [1,1]
输出：1
解释：如图所示，你只能看到两点之一。

提示：

1 <= points.length <= 10^5
points[i].length == 2
location.length == 2
0 <= angle < 360
0 <= posx, posy, xi, yi <= 100
*/
func main() {
	var tests = []struct {
		points   [][]int
		angle    int
		location []int
		want     int
	}{
		{
			points: [][]int{
				{2, 1}, {2, 2}, {3, 4}, {1, 1},
			},
			angle:    90,
			location: []int{1, 1},
			want:     4,
		},
		{
			points: [][]int{
				{1, 0}, {2, 1},
			},
			angle:    13,
			location: []int{1, 1},
			want:     1,
		},
		{
			points: [][]int{
				{1, 1}, {2, 2}, {3, 3}, {4, 4}, {1, 2}, {2, 1},
			},
			angle:    0,
			location: []int{1, 1},
			want:     4,
		},
		{
			points: [][]int{
				{41, 7}, {22, 94}, {90, 53}, {94, 54}, {58, 50}, {51, 96}, {87, 88}, {55, 98}, {65, 62}, {36, 47}, {91, 61}, {15, 41}, {31, 94}, {82, 80}, {42, 73}, {79, 6}, {45, 4},
			},
			angle:    17,
			location: []int{6, 84},
			want:     5,
		},
	}

	for _, item := range tests {
		if ans := visiblePoints(item.points, item.angle, item.location); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func visiblePoints(points [][]int, angle int, location []int) int {
	pointAngles := make([]float64, 0, len(points)*2)
	zeroPointCount := 0
	for i := 0; i < len(points); i++ {
		ag := calRadian(points[i], location)
		if ag == -1 {
			zeroPointCount++
		} else {
			pointAngles = append(pointAngles, ag)
		}
	}
	sort.Float64s(pointAngles)
	for _, item := range pointAngles {
		pointAngles = append(pointAngles, item+DegreeToRadian(360))
	}

	var maxCount int
	var count int
	var radian = DegreeToRadian(float64(angle))
	for i, j := 0, 0; i < len(pointAngles); i++ {
		if i == 0 {
			count = 1
		} else {
			count--
		}
		for j+1 < len(pointAngles) && pointAngles[i]+radian >= pointAngles[j+1] {
			count++
			j++
		}
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount + zeroPointCount
}

func calRadian(point, location []int) float64 {
	// 以 location 为轴心
	switch {
	case point[0] == location[0] && point[1] == location[1]: // 轴心上
		return -1
	case point[0] == location[0] && point[1] > location[1]: // y正轴
		return DegreeToRadian(90)
	case point[0] == location[0] && point[1] < location[1]: // y负轴
		return DegreeToRadian(270)
	case point[0] > location[0] && point[1] == location[1]: // x正轴
		return DegreeToRadian(0)
	case point[0] < location[0] && point[1] == location[1]: // x负轴
		return DegreeToRadian(180)
	case point[0] > location[0] && point[1] > location[1]: // 第一象限
		return math.Atan(float64(point[1]-location[1]) / float64(point[0]-location[0]))
	case point[0] < location[0] && point[1] > location[1]: // 第二象限
		return DegreeToRadian(180) - math.Atan(float64(point[1]-location[1])/-float64(point[0]-location[0]))
	case point[0] < location[0] && point[1] < location[1]: // 第三象限
		return DegreeToRadian(180) + math.Atan(-float64(point[1]-location[1])/-float64(point[0]-location[0]))
	case point[0] > location[0] && point[1] < location[1]: // 第四象限
		return DegreeToRadian(360) - math.Atan(-float64(point[1]-location[1])/float64(point[0]-location[0]))
	default:
		panic("point cal angle error")
	}
}

// DegreeToRadian 角度转弧度
func DegreeToRadian(x float64) float64 {
	return x * (math.Pi / 180)
}
