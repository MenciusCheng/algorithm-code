package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/jump-game-iv/

1345. 跳跃游戏 IV
给你一个整数数组 arr ，你一开始在数组的第一个元素处（下标为 0）。

每一步，你可以从下标 i 跳到下标：

i + 1 满足：i + 1 < arr.length
i - 1 满足：i - 1 >= 0
j 满足：arr[i] == arr[j] 且 i != j
请你返回到达数组最后一个元素的下标处所需的 最少操作次数 。

注意：任何时候你都不能跳到数组外面。

示例 1：

输入：arr = [100,-23,-23,404,100,23,23,23,3,404]
输出：3
解释：那你需要跳跃 3 次，下标依次为 0 --> 4 --> 3 --> 9 。下标 9 为数组的最后一个元素的下标。
示例 2：

输入：arr = [7]
输出：0
解释：一开始就在最后一个元素处，所以你不需要跳跃。
示例 3：

输入：arr = [7,6,9,6,9,6,9,7]
输出：1
解释：你可以直接从下标 0 处跳到下标 7 处，也就是数组的最后一个元素处。
示例 4：

输入：arr = [6,1,9]
输出：2
示例 5：

输入：arr = [11,22,7,7,7,7,7,7,7,22,13]
输出：3

提示：

1 <= arr.length <= 5 * 10^4
-10^8 <= arr[i] <= 10^8
*/
func main() {
	var tests = []struct {
		arr  []int
		want int
	}{
		{
			arr:  []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404},
			want: 3,
		},
		{
			arr:  []int{7},
			want: 0,
		},
		{
			arr:  []int{7, 6, 9, 6, 9, 6, 9, 7},
			want: 1,
		},
		{
			arr:  []int{6, 1, 9},
			want: 2,
		},
		{
			arr:  []int{11, 22, 7, 7, 7, 7, 7, 7, 7, 22, 13},
			want: 3,
		},
	}

	for _, item := range tests {
		if ans := minJumps(item.arr); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func minJumps(arr []int) int {
	valueIndexMap := make(map[int][]int)
	for i, v := range arr {
		valueIndexMap[v] = append(valueIndexMap[v], i)
	}

	visited := make(map[int]bool)
	jumpers := []int{0}
	step := 0
	for len(jumpers) > 0 {
		nextList := make([]int, 0, len(jumpers))
		for _, jumper := range jumpers {
			if jumper == len(arr)-1 {
				return step
			}

			nextIndex := jumper + 1
			if nextIndex < len(arr) && !visited[nextIndex] {
				nextList = append(nextList, nextIndex)
				visited[nextIndex] = true
			}

			nextIndex = jumper - 1
			if nextIndex > 0 && !visited[nextIndex] {
				nextList = append(nextList, nextIndex)
				visited[nextIndex] = true
			}

			if len(valueIndexMap[arr[jumper]]) > 0 {
				for _, v := range valueIndexMap[arr[jumper]] {
					if v != jumper && !visited[v] {
						nextList = append(nextList, v)
						visited[v] = true
					}
				}
				valueIndexMap[arr[jumper]] = []int{}
			}
		}

		step++
		jumpers = nextList
	}

	return step
}
