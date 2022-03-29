package w20220327

import "strconv"

func findDifference(nums1 []int, nums2 []int) [][]int {
	cnt1 := make(map[int]bool)
	for _, num := range nums1 {
		cnt1[num] = true
	}
	cnt2 := make(map[int]bool)
	for _, num := range nums2 {
		cnt2[num] = true
	}

	answer1 := make([]int, 0)
	ansM1 := make(map[int]bool)
	for _, num := range nums1 {
		if !cnt2[num] && !ansM1[num] {
			answer1 = append(answer1, num)
			ansM1[num] = true
		}
	}
	answer2 := make([]int, 0)
	ansM2 := make(map[int]bool)
	for _, num := range nums2 {
		if !cnt1[num] && !ansM2[num] {
			answer2 = append(answer2, num)
			ansM2[num] = true
		}
	}
	answer := make([][]int, 0)
	answer = append(answer, answer1)
	answer = append(answer, answer2)
	return answer
}

func minDeletion(nums []int) int {
	n := len(nums)
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			res++
		} else {
			i++
		}
	}
	if (n-res)%2 == 1 {
		res++
	}
	return res
}

func kthPalindrome(queries []int, intLength int) []int64 {
	headLen := intLength / 2
	if intLength%2 > 0 {
		headLen++
	}
	start := 1
	end := 9
	if headLen > 1 {
		for i := 1; i < headLen; i++ {
			start *= 10
			end = end*10 + 9
		}
	}

	res := make([]int64, 0)
	for _, query := range queries {
		i := query + start - 1
		if i > end {
			res = append(res, -1)
		} else {
			res = append(res, fillPalindrome(i, intLength))
		}
	}
	return res
}

func fillPalindrome(a int, intLength int) int64 {
	s := strconv.Itoa(a)
	sb := []byte(s)
	index := len(sb) - 1
	if intLength%2 > 0 {
		index--
	}
	for i := index; i >= 0; i-- {
		sb = append(sb, s[i])
	}
	res, _ := strconv.Atoi(string(sb))
	return int64(res)
}

/*
5269. 从栈中取出 K 个硬币的最大面值和

一张桌子上总共有 n 个硬币 栈 。每个栈有 正整数 个带面值的硬币。
每一次操作中，你可以从任意一个栈的 顶部 取出 1 个硬币，从栈中移除它，并放入你的钱包里。
给你一个列表 piles ，其中 piles[i] 是一个整数数组，分别表示第 i 个栈里 从顶到底 的硬币面值。同时给你一个正整数 k ，请你返回在 恰好 进行 k 次操作的前提下，你钱包里硬币面值之和 最大为多少 。

示例 1：

输入：piles = [[1,100,3],[7,8,9]], k = 2
输出：101
解释：
上图展示了几种选择 k 个硬币的不同方法。
我们可以得到的最大面值为 101 。
示例 2：

输入：piles = [[100],[100],[100],[100],[100],[100],[1,1,1,1,1,1,700]], k = 7
输出：706
解释：
如果我们所有硬币都从最后一个栈中取，可以得到最大面值和。

提示：

n == piles.length
1 <= n <= 1000
1 <= piles[i][j] <= 10^5
1 <= k <= sum(piles[i].length) <= 2000
*/
func maxValueOfCoins(piles [][]int, k int) int {
	dp := make([][]int, len(piles))
	for i := 0; i < len(piles); i++ {
		dp[i] = make([]int, k+1)
		for j := 1; j < len(piles[i]); j++ {
			piles[i][j] += piles[i][j-1]
		}

		if i == 0 {
			for j := 0; j < len(piles[i]) && j < k; j++ {
				dp[i][j+1] = piles[i][j]
			}
		} else {
			for j := k; j > 0; j-- {
				dp[i][j] = dp[i-1][j]
				for w, v := range piles[i] {
					if j-w-1 < 0 {
						break
					}
					dp[i][j] = max(dp[i][j], dp[i-1][j-w-1]+v)
				}
			}
		}
	}
	return dp[len(piles)-1][k]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
