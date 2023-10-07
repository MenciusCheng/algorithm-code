package main

import "fmt"

/*
https://leetcode.cn/problems/online-stock-span/description/

901. 股票价格跨度

中等

设计一个算法收集某些股票的每日报价，并返回该股票当日价格的 跨度 。
当日股票价格的 跨度 被定义为股票价格小于或等于今天价格的最大连续日数（从今天开始往回数，包括今天）。
例如，如果未来 7 天股票的价格是 [100,80,60,70,60,75,85]，那么股票跨度将是 [1,1,1,2,1,4,6] 。

实现 StockSpanner 类：
StockSpanner() 初始化类对象。
int next(int price) 给出今天的股价 price ，返回该股票当日价格的 跨度 。

示例：

输入：
["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]
[[], [100], [80], [60], [70], [60], [75], [85]]
输出：
[null, 1, 1, 1, 2, 1, 4, 6]

解释：
StockSpanner stockSpanner = new StockSpanner();
stockSpanner.next(100); // 返回 1
stockSpanner.next(80);  // 返回 1
stockSpanner.next(60);  // 返回 1
stockSpanner.next(70);  // 返回 2
stockSpanner.next(60);  // 返回 1
stockSpanner.next(75);  // 返回 4 ，因为截至今天的最后 4 个股价 (包括今天的股价 75) 都小于或等于今天的股价。
stockSpanner.next(85);  // 返回 6

提示：

1 <= price <= 10^5
最多调用 next 方法 10^4 次
*/
func main() {
	stockSpanner := Constructor()
	fmt.Println(stockSpanner.Next(100)) // 返回 1
	fmt.Println(stockSpanner.Next(80))  // 返回 1
	fmt.Println(stockSpanner.Next(60))  // 返回 1
	fmt.Println(stockSpanner.Next(70))  // 返回 2
	fmt.Println(stockSpanner.Next(60))  // 返回 1
	fmt.Println(stockSpanner.Next(75))  // 返回 4 ，因为截至今天的最后 4 个股价 (包括今天的股价 75) 都小于或等于今天的股价。
	fmt.Println(stockSpanner.Next(85))  // 返回 6
}

type StockSpanner struct {
	Stack [][]int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	count := 1
	for len(this.Stack) > 0 && price >= this.Stack[len(this.Stack)-1][0] {
		count += this.Stack[len(this.Stack)-1][1]
		this.Stack = this.Stack[:len(this.Stack)-1]
	}
	this.Stack = append(this.Stack, []int{price, count})
	return count
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
