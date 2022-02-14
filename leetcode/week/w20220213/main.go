package w20220213

import "sort"

func countOperations(num1 int, num2 int) int {
	res := 0
	for num1 != 0 && num2 != 0 {
		if num1 >= num2 {
			num1 = num1 - num2
		} else {
			num2 = num2 - num1
		}
		res++
	}
	return res
}

/*
6005. 使数组变成交替数组的最少操作数
*/
func minimumOperations(nums []int) int {
	var len0, len1 int
	m0 := make(map[int]int)
	m1 := make(map[int]int)
	for i, num := range nums {
		if i%2 == 0 {
			m0[num]++
			len0++
		} else {
			m1[num]++
			len1++
		}
	}

	a0 := make([][2]int, 0, len(m0))
	a1 := make([][2]int, 0, len(m1))
	for k, v := range m0 {
		a0 = append(a0, [2]int{k, v})
	}
	for k, v := range m1 {
		a1 = append(a1, [2]int{k, v})
	}

	if len(a0) == 1 && len(a1) == 1 {
		if a0[0][0] != a1[0][0] {
			return 0
		} else {
			return min(len0, len1)
		}
	} else if len(a0) <= 1 && len(a1) <= 1 {
		return 0
	}

	sort.Slice(a0, func(i, j int) bool {
		return a0[i][1] > a0[j][1]
	})
	sort.Slice(a1, func(i, j int) bool {
		return a1[i][1] > a1[j][1]
	})

	if len(a0) >= 2 {
		return cp(a0, a1, len0, len1)
	} else {
		return cp(a1, a0, len1, len0)
	}
}

func cp(a0, a1 [][2]int, len0, len1 int) int {
	if len(a1) >= 2 {
		ch0 := len0 - a0[0][1]
		ch1 := len1 - a1[0][1]
		if a0[0][0] != a1[0][0] {
			return ch0 + ch1
		} else {
			ch01 := len0 - a0[1][1]
			ch11 := len1 - a1[1][1]
			return min(ch0+ch11, ch1+ch01)
		}
	} else if len(a1) == 1 {
		ch0 := len0 - a0[0][1]
		if a0[0][0] != a1[0][0] {
			return ch0
		} else {
			ch01 := len0 - a0[1][1]
			return min(ch01, ch0+len1)
		}
	} else {
		return len0 - a0[0][1]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
