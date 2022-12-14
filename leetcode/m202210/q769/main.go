package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/max-chunks-to-make-sorted/

769. 最多能完成排序的块
中等

给定一个长度为 n 的整数数组 arr ，它表示在 [0, n - 1] 范围内的整数的排列。
我们将 arr 分割成若干 块 (即分区)，并对每个块单独排序。将它们连接起来后，使得连接的结果和按升序排序后的原数组相同。
返回数组能分成的最多块数量。

示例 1:

输入: arr = [4,3,2,1,0]
输出: 1
解释:
将数组分成2块或者更多块，都无法得到所需的结果。
例如，分成 [4, 3], [2, 1, 0] 的结果是 [3, 4, 0, 1, 2]，这不是有序的数组。
示例 2:

输入: arr = [1,0,2,3,4]
输出: 4
解释:
我们可以把它分成两块，例如 [1, 0], [2, 3, 4]。
然而，分成 [1, 0], [2], [3], [4] 可以得到最多的块数。

提示:

n == arr.length
1 <= n <= 10
0 <= arr[i] < n
*/
func main() {
	var tests = []struct {
		arr  []int
		want int
	}{
		{
			arr:  []int{4, 3, 2, 1, 0},
			want: 1,
		},
		{
			arr:  []int{1, 0, 2, 3, 4},
			want: 4,
		},
	}

	for _, item := range tests {
		if ans := maxChunksToSorted(item.arr); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func maxChunksToSorted(arr []int) int {
	res := 0
	need := 0
	for i, item := range arr {
		need = max(need, item)
		if i >= need {
			res++
			need = 0
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}