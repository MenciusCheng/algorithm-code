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
