package w20211226

import "strconv"

/*
5963. 反转两次的数字
*/
func isSameAfterReversals(num int) bool {
	if num == 0 {
		return true
	}
	str := strconv.Itoa(num)

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '0' {
			return false
		} else {
			return true
		}
	}
	return true
}

/*
5964. 执行所有后缀指令
*/
func executeInstructions(n int, startPos []int, s string) []int {
	res := make([]int, 0, len(s))
	for i := 0; i < len(s); i++ {
		row, col := startPos[0], startPos[1]
		count := 0
		for j := i; j < len(s); j++ {
			isIn := true
			switch s[j] {
			case 'R':
				col++
				if col > n-1 {
					isIn = false
				}
			case 'L':
				col--
				if col < 0 {
					isIn = false
				}
			case 'U':
				row--
				if row < 0 {
					isIn = false
				}
			case 'D':
				row++
				if row > n-1 {
					isIn = false
				}
			}
			if !isIn {
				break
			}
			count++
		}
		res = append(res, count)
	}
	return res
}

/*
5965. 相同元素的间隔之和
*/
func getDistances(arr []int) []int64 {
	cnt := make(map[int][]int)
	for i, item := range arr {
		cnt[item] = append(cnt[item], i)
	}
	res := make([]int64, len(arr))

	for item, ints := range cnt {
		var absSum int64
		for _, index := range ints {
			absSum += abs(ints[0], index)
		}
		res[ints[0]] = absSum

		left := int64(0)
		right := int64(len(cnt[item]) - 1)
		for i := 1; i < len(ints); i++ {
			right--
			gap := int64(ints[i] - ints[i-1])
			absSum -= right * gap
			absSum += left * gap
			left++
			res[ints[i]] = absSum
		}
	}
	return res
}

func abs(a, b int) int64 {
	if a > b {
		return int64(a - b)
	} else {
		return int64(b - a)
	}
}
