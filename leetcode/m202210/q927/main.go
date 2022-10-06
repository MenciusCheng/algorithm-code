package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/three-equal-parts/

927. 三等分
给定一个由 0 和 1 组成的数组 arr ，将数组分成  3 个非空的部分 ，使得所有这些部分表示相同的二进制值。

如果可以做到，请返回任何 [i, j]，其中 i+1 < j，这样一来：

arr[0], arr[1], ..., arr[i] 为第一部分；
arr[i + 1], arr[i + 2], ..., arr[j - 1] 为第二部分；
arr[j], arr[j + 1], ..., arr[arr.length - 1] 为第三部分。
这三个部分所表示的二进制值相等。
如果无法做到，就返回 [-1, -1]。

注意，在考虑每个部分所表示的二进制时，应当将其看作一个整体。例如，[1,1,0] 表示十进制中的 6，而不会是 3。此外，前导零也是被允许的，所以 [0,1,1] 和 [1,1] 表示相同的值。

示例 1：

输入：arr = [1,0,1,0,1]
输出：[0,3]
示例 2：

输入：arr = [1,1,0,1,1]
输出：[-1,-1]
示例 3:

输入：arr = [1,1,0,0,1]
输出：[0,2]


提示：

3 <= arr.length <= 3 * 10^4
arr[i] 是 0 或 1
*/
func main() {
	var tests = []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{0, 0, 0, 1, 0, 1, 1, 0, 1, 1, 1, 0, 0, 1, 0, 1, 1},
			want: []int{-1, -1},
		},
		{
			arr:  []int{1, 0, 1, 0, 1},
			want: []int{0, 3},
		},
		{
			arr:  []int{1, 1, 0, 1, 1},
			want: []int{-1, -1},
		},
		{
			arr:  []int{1, 1, 0, 0, 1},
			want: []int{0, 2},
		},
	}

	for _, item := range tests {
		if ans := threeEqualParts(item.arr); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func threeEqualParts(arr []int) []int {
	sum := 0
	for _, item := range arr {
		sum += item
	}
	if sum == 0 {
		return []int{0, len(arr) - 1}
	}
	if sum%3 != 0 {
		return []int{-1, -1}
	}
	avg := sum / 3

	count := 0
	var f1, f2, f3 int
	for i, item := range arr {
		if item == 1 {
			count++

			if count == 1 {
				f1 = i
			} else if count == avg+1 {
				f2 = i
			} else if count == avg*2+1 {
				f3 = i
				break
			}
		}
	}

	s3 := fmt.Sprintf("%v", arr[f3:])
	zeroCount3 := 0
	for i := f3; i < len(arr); i++ {
		if arr[i] == 0 {
			zeroCount3++
		}
	}

	zeroCount2 := 0
	for i := f2; i < f3; i++ {
		if arr[i] == 0 {
			zeroCount2++
		}
	}
	for zeroCount2 > zeroCount3 && arr[f3-1] == 0 {
		f3--
		zeroCount2--
	}
	if zeroCount2 != zeroCount3 {
		return []int{-1, -1}
	}
	s2 := fmt.Sprintf("%v", arr[f2:f3])

	zeroCount1 := 0
	for i := f1; i < f2; i++ {
		if arr[i] == 0 {
			zeroCount1++
		}
	}
	for zeroCount1 > zeroCount2 && arr[f2-1] == 0 {
		f2--
		zeroCount1--
	}
	if zeroCount1 != zeroCount2 {
		return []int{-1, -1}
	}
	s1 := fmt.Sprintf("%v", arr[f1:f2])

	if s1 == s2 && s2 == s3 {
		return []int{f2 - 1, f3}
	}
	return []int{-1, -1}
}
