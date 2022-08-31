package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/validate-stack-sequences/

946. 验证栈序列
给定 pushed 和 popped 两个序列，每个序列中的 值都不重复，只有当它们可能是在最初空栈上进行的推入 push 和弹出 pop 操作序列的结果时，返回 true；否则，返回 false 。

示例 1：

输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
输出：true
解释：我们可以按以下顺序执行：
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
示例 2：

输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
输出：false
解释：1 不能在 2 之前弹出。

提示：

1 <= pushed.length <= 1000
0 <= pushed[i] <= 1000
pushed 的所有元素 互不相同
popped.length == pushed.length
popped 是 pushed 的一个排列
*/
func main() {
	var tests = []struct {
		pushed []int
		popped []int
		want   bool
	}{
		{
			pushed: []int{1, 2, 3, 4, 5},
			popped: []int{4, 5, 3, 2, 1},
			want:   true,
		},
		{
			pushed: []int{1, 2, 3, 4, 5},
			popped: []int{4, 3, 5, 1, 2},
			want:   false,
		},
	}

	for _, item := range tests {
		if ans := validateStackSequences(item.pushed, item.popped); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0, len(pushed))
	var i, j int
	for i < len(pushed) || j < len(popped) {
		if len(stack) > 0 && j < len(popped) && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		} else {
			if i >= len(pushed) {
				return false
			}
			stack = append(stack, pushed[i])
			i++
		}
	}

	return true
}

//if len(stack) == 0 || stack[len(stack)-1] != popped[j] {
//	if i >= len(pushed) {
//		return false
//	}
//	stack = append(stack, pushed[i])
//	i++
//} else {
//	stack = stack[:len(stack)-1]
//	j++
//}
