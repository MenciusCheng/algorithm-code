package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/jewels-and-stones/

771. 宝石与石头
提示
简单
779
相关企业
 给你一个字符串 jewels 代表石头中宝石的类型，另有一个字符串 stones 代表你拥有的石头。 stones 中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。
字母区分大小写，因此 "a" 和 "A" 是不同类型的石头。

示例 1：

输入：jewels = "aA", stones = "aAAbbbb"
输出：3
示例 2：

输入：jewels = "z", stones = "ZZ"
输出：0

提示：

1 <= jewels.length, stones.length <= 50
jewels 和 stones 仅由英文字母组成
jewels 中的所有字符都是 唯一的
*/
func main() {
	var tests = []struct {
		jewels string
		stones string
		want   int
	}{
		{},
	}

	for _, item := range tests {
		if ans := numJewelsInStones(item.jewels, item.stones); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func numJewelsInStones(jewels string, stones string) int {
	m := make(map[byte]bool)
	for i := range jewels {
		m[jewels[i]] = true
	}
	res := 0
	for i := range stones {
		if m[stones[i]] {
			res++
		}
	}
	return res
}
