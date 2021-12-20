package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
https://leetcode-cn.com/problems/heaters/

475. 供暖器
冬季已经来临。 你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。
在加热器的加热半径范围内的每个房屋都可以获得供暖。
现在，给出位于一条水平线上的房屋 houses 和供暖器 heaters 的位置，请你找出并返回可以覆盖所有房屋的最小加热半径。
说明：所有供暖器都遵循你的半径标准，加热的半径也一样。

示例 1:

输入: houses = [1,2,3], heaters = [2]
输出: 1
解释: 仅在位置2上有一个供暖器。如果我们将加热半径设为1，那么所有房屋就都能得到供暖。
示例 2:

输入: houses = [1,2,3,4], heaters = [1,4]
输出: 1
解释: 在位置1, 4上有两个供暖器。我们需要将加热半径设为1，这样所有房屋就都能得到供暖。
示例 3：

输入：houses = [1,5], heaters = [2]
输出：3

提示：

1 <= houses.length, heaters.length <= 3 * 10^4
1 <= houses[i], heaters[i] <= 10^9
*/
func main() {
	var tests = []struct {
		houses  []int
		heaters []int
		want    int
	}{
		{
			houses:  []int{1, 2, 3},
			heaters: []int{2},
			want:    1,
		},
		{
			houses:  []int{1, 2, 3, 4},
			heaters: []int{1, 4},
			want:    1,
		},
		{
			houses:  []int{1, 5},
			heaters: []int{2},
			want:    3,
		},
		{
			houses:  []int{282475249, 622650073, 984943658, 144108930, 470211272, 101027544, 457850878, 458777923},
			heaters: []int{823564440, 115438165, 784484492, 74243042, 114807987, 137522503, 441282327, 16531729, 823378840, 143542612},
			want:    161834419,
		},
	}

	for _, item := range tests {
		if ans := findRadius(item.houses, item.heaters); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func findRadius(houses []int, heaters []int) int {
	minRadius := 0
	sort.Ints(heaters)
	for _, house := range houses {
		heater1, heater2 := findNearest(house, heaters)
		minGap := abs(house, heater1)
		if abs(house, heater2) < minGap {
			minGap = abs(house, heater2)
		}
		if minGap > minRadius {
			minRadius = minGap
		}
	}
	return minRadius
}

func abs(a, b int) int {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}

// 二分查找法
func binarySearch(arr []int, target int) (bool, int) {
	var mid int
	head, tail := 0, len(arr)-1
	for head <= tail {
		mid = (head + tail) / 2

		if arr[mid] == target {
			return true, mid
		} else if arr[mid] < target {
			head = mid + 1
		} else {
			tail = mid - 1
			if mid > 0 {
				mid -= 1 // 如果没有相等的值，arr[mid] 最后要小于 target，除非 mid == 0
			}
		}
	}
	return false, mid
}

func findNearest(house int, heaters []int) (int, int) {
	isMatch, mid := binarySearch(heaters, house)
	if isMatch {
		return heaters[mid], heaters[mid]
	} else {
		next := mid + 1
		if next >= len(heaters) {
			next -= 1
		}
		return heaters[mid], heaters[next]
	}
}
