package main

import (
	"fmt"
	"math"
)

/*
https://leetcode.cn/problems/k-th-symbol-in-grammar/

901. 股票价格跨度
中等

编写一个 StockSpanner 类，它收集某些股票的每日报价，并返回该股票当日价格的跨度。

今天股票价格的跨度被定义为股票价格小于或等于今天价格的最大连续日数（从今天开始往回数，包括今天）。

例如，如果未来7天股票的价格是 [100, 80, 60, 70, 60, 75, 85]，那么股票跨度将是 [1, 1, 1, 2, 1, 4, 6]。

示例：

输入：["StockSpanner","next","next","next","next","next","next","next"], [[],[100],[80],[60],[70],[60],[75],[85]]
输出：[null,1,1,1,2,1,4,6]
解释：
首先，初始化 S = StockSpanner()，然后：
S.next(100) 被调用并返回 1，
S.next(80) 被调用并返回 1，
S.next(60) 被调用并返回 1，
S.next(70) 被调用并返回 2，
S.next(60) 被调用并返回 1，
S.next(75) 被调用并返回 4，
S.next(85) 被调用并返回 6。

注意 (例如) S.next(75) 返回 4，因为截至今天的最后 4 个价格
(包括今天的价格 75) 小于或等于今天的价格。

提示：

调用 StockSpanner.next(int price) 时，将有 1 <= price <= 10^5。
每个测试用例最多可以调用  10000 次 StockSpanner.next。
在所有测试用例中，最多调用 150000 次 StockSpanner.next。
此问题的总时间限制减少了 50%。
*/
func main() {
	S := Constructor()
	//fmt.Println(S.Next(100) == 1) // 被调用并返回 1，
	//fmt.Println(S.Next(80) == 1)  // 被调用并返回 1，
	//fmt.Println(S.Next(60) == 1)  // 被调用并返回 1，
	//fmt.Println(S.Next(70) == 2)  // 被调用并返回 2，
	//fmt.Println(S.Next(60) == 1)  // 被调用并返回 1，
	//fmt.Println(S.Next(75) == 4)  // 被调用并返回 4，
	//fmt.Println(S.Next(85) == 6)  // 被调用并返回 6。

	// [[],[32],[82],[73],[99],[91]]
	// [null,1,2,1,4,1]
	fmt.Println(S.Next(32) == 1)
	fmt.Println(S.Next(82) == 2)
	fmt.Println(S.Next(73) == 1)
	fmt.Println(S.Next(99) == 4)
	fmt.Println(S.Next(91) == 1)
}

type StockSpanner struct {
	Stack [][2]int
	Idx   int
}

func Constructor() StockSpanner {
	return StockSpanner{
		Stack: [][2]int{{-1, math.MaxInt32}},
		Idx:   -1,
	}
}

func (this *StockSpanner) Next(price int) int {
	this.Idx++
	for price >= this.Stack[len(this.Stack)-1][1] {
		this.Stack = this.Stack[:len(this.Stack)-1]
	}
	this.Stack = append(this.Stack, [2]int{this.Idx, price})
	return this.Idx - this.Stack[len(this.Stack)-2][0]
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
