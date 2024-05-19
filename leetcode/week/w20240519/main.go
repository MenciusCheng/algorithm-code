package w20240519

import "strconv"

// https://leetcode.cn/problems/permutation-difference-between-two-strings/
func findPermutationDifference(s string, t string) int {
	cnt := map[byte]int{}
	for i, b := range []byte(t) {
		cnt[b] = i
	}
	res := 0
	for i, b := range []byte(s) {
		res += abs(i, cnt[b])
	}
	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// https://leetcode.cn/problems/special-array-ii/
func isArraySpecial(nums []int, queries [][]int) []bool {
	a := make([]int, 0, len(nums))
	a = append(a, 0)
	for i := 0; i < len(nums)-1; i++ {
		v := 0
		if nums[i]%2 == nums[i+1]%2 {
			v = 1
		}
		a = append(a, v+a[i])
	}
	res := make([]bool, len(queries))
	for i, q := range queries {
		if a[q[0]] == a[q[1]] {
			res[i] = true
		}
	}
	return res
}

// https://leetcode.cn/problems/sum-of-digit-differences-of-all-pairs/
func sumDigitDifferences(nums []int) int64 {
	res := 0
	for i := 0; i < len(strconv.Itoa(nums[0])); i++ {
		cnt := make(map[int]int)
		for _, num := range nums {
			cnt[int(strconv.Itoa(num)[i]-'0')]++
		}
		pre := 0
		for _, v := range cnt {
			res += v * pre
			pre += v
		}
	}
	return int64(res)
}
