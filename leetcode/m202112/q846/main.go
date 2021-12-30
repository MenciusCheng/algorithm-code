package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
https://leetcode-cn.com/problems/hand-of-straights/

846. 一手顺子
Alice 手中有一把牌，她想要重新排列这些牌，分成若干组，使每一组的牌数都是 groupSize ，并且由 groupSize 张连续的牌组成。
给你一个整数数组 hand 其中 hand[i] 是写在第 i 张牌，和一个整数 groupSize 。如果她可能重新排列这些牌，返回 true ；否则，返回 false 。

示例 1：

输入：hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
输出：true
解释：Alice 手中的牌可以被重新排列为 [1,2,3]，[2,3,4]，[6,7,8]。
示例 2：

输入：hand = [1,2,3,4,5], groupSize = 4
输出：false
解释：Alice 手中的牌无法被重新排列成几个大小为 4 的组。

提示：

1 <= hand.length <= 10^4
0 <= hand[i] <= 10^9
1 <= groupSize <= hand.length
*/
func main() {
	var tests = []struct {
		hand      []int
		groupSize int
		want      bool
	}{
		{
			hand:      []int{1, 2, 3, 6, 2, 3, 4, 7, 8},
			groupSize: 3,
			want:      true,
		},
		{
			hand:      []int{1, 2, 3, 4, 5},
			groupSize: 4,
			want:      false,
		},
		{
			hand:      []int{1, 1, 2, 2, 3, 3},
			groupSize: 3,
			want:      true,
		},
	}

	for _, item := range tests {
		if ans := isNStraightHand(item.hand, item.groupSize); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func isNStraightHand(hand []int, groupSize int) bool {
	sort.Ints(hand)
	if len(hand)%groupSize != 0 {
		return false
	}

	size := groupSize
	for i := len(hand) - 1; i >= 0; i-- {
		if i == 0 || size == 1 || hand[i] == hand[i-1]+1 {
			size--
			if size == 0 {
				size = groupSize
			}
			continue
		}

		p := 2
		for i-p >= 0 {
			if hand[i] == hand[i-p]+1 {
				for j := p; j > 1; j-- {
					hand[i-j], hand[i-j+1] = hand[i-j+1], hand[i-j]
				}
				break
			}
			p++
		}
		if hand[i] != hand[i-1]+1 {
			return false
		}
		size--
		if size == 0 {
			size = groupSize
		}
	}
	return true
}
