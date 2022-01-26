package main

import (
	"fmt"
)

/*
https://leetcode-cn.com/problems/detect-squares/

2013. 检测正方形
给你一个在 X-Y 平面上的点构成的数据流。设计一个满足下述要求的算法：

添加 一个在数据流中的新点到某个数据结构中。可以添加 重复 的点，并会视作不同的点进行处理。
给你一个查询点，请你从数据结构中选出三个点，使这三个点和查询点一同构成一个 面积为正 的 轴对齐正方形 ，统计 满足该要求的方案数目。
轴对齐正方形 是一个正方形，除四条边长度相同外，还满足每条边都与 x-轴 或 y-轴 平行或垂直。

实现 DetectSquares 类：

DetectSquares() 使用空数据结构初始化对象
void add(int[] point) 向数据结构添加一个新的点 point = [x, y]
int count(int[] point) 统计按上述方式与点 point = [x, y] 共同构造 轴对齐正方形 的方案数。

示例：

输入：
["DetectSquares", "add", "add", "add", "count", "count", "add", "count"]
[[], [[3, 10]], [[11, 2]], [[3, 2]], [[11, 10]], [[14, 8]], [[11, 2]], [[11, 10]]]
输出：
[null, null, null, null, 1, 0, null, 2]

解释：
DetectSquares detectSquares = new DetectSquares();
detectSquares.add([3, 10]);
detectSquares.add([11, 2]);
detectSquares.add([3, 2]);
detectSquares.count([11, 10]); // 返回 1 。你可以选择：
                               //   - 第一个，第二个，和第三个点
detectSquares.count([14, 8]);  // 返回 0 。查询点无法与数据结构中的这些点构成正方形。
detectSquares.add([11, 2]);    // 允许添加重复的点。
detectSquares.count([11, 10]); // 返回 2 。你可以选择：
                               //   - 第一个，第二个，和第三个点
                               //   - 第一个，第三个，和第四个点

提示：

point.length == 2
0 <= x, y <= 1000
调用 add 和 count 的 总次数 最多为 5000
*/
func main() {
	detectSquares := Constructor()
	detectSquares.Add([]int{3, 10})
	detectSquares.Add([]int{11, 2})
	detectSquares.Add([]int{3, 2})
	fmt.Println(detectSquares.Count([]int{11, 10}) == 1) // 返回 1 。你可以选择：
	//   - 第一个，第二个，和第三个点
	fmt.Println(detectSquares.Count([]int{14, 8}) == 0)  // 返回 0 。查询点无法与数据结构中的这些点构成正方形。
	detectSquares.Add([]int{11, 2})                      // 允许添加重复的点。
	fmt.Println(detectSquares.Count([]int{11, 10}) == 2) // 返回 2 。你可以选择：
}

type DetectSquares struct {
	Points       map[[2]int]int
	PointYMapByX map[int][]int
}

func Constructor() DetectSquares {
	ds := DetectSquares{}
	ds.Points = make(map[[2]int]int)
	ds.PointYMapByX = make(map[int][]int)
	return ds
}

func (this *DetectSquares) Add(point []int) {
	this.Points[[2]int{point[0], point[1]}] += 1
	this.PointYMapByX[point[0]] = append(this.PointYMapByX[point[0]], point[1])
}

func (this *DetectSquares) Count(point []int) int {
	x, y := point[0], point[1]
	ys := this.PointYMapByX[point[0]]
	sum := 0
	for _, item := range ys {
		if item < y {
			leftX := x - (y - item)
			if this.Points[[2]int{leftX, y}] > 0 && this.Points[[2]int{leftX, item}] > 0 {
				sum += this.Points[[2]int{leftX, y}] * this.Points[[2]int{leftX, item}]
			}

			rightX := x + y - item
			if this.Points[[2]int{rightX, y}] > 0 && this.Points[[2]int{rightX, item}] > 0 {
				sum += this.Points[[2]int{rightX, y}] * this.Points[[2]int{rightX, item}]
			}
		} else if item > y {
			leftX := x - (item - y)
			if this.Points[[2]int{leftX, y}] > 0 && this.Points[[2]int{leftX, item}] > 0 {
				sum += this.Points[[2]int{leftX, y}] * this.Points[[2]int{leftX, item}]
			}

			rightX := x + item - y
			if this.Points[[2]int{rightX, y}] > 0 && this.Points[[2]int{rightX, item}] > 0 {
				sum += this.Points[[2]int{rightX, y}] * this.Points[[2]int{rightX, item}]
			}
		}
	}
	return sum
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
