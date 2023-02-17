package w20230212

import (
	"fmt"
	"strconv"
)

func findTheArrayConcVal(nums []int) int64 {
	var res int64
	for len(nums) > 0 {
		if len(nums) == 1 {
			res += int64(nums[0])
			break
		} else {
			v, _ := strconv.Atoi(fmt.Sprintf("%d%d", nums[0], nums[len(nums)-1]))
			res += int64(v)
			nums = nums[1 : len(nums)-1]
		}
	}
	return res
}

func countFairPairs(nums []int, lower int, upper int) int64 {
	return 0
}

func find(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return i
}

func substringXorQueries(s string, queries [][]int) [][]int {
	res := make([][]int, len(queries))
	cnt := make(map[int][]int)
	maxV := 0
	for i, item := range queries {
		v := item[0] ^ item[1]
		cnt[v] = append(cnt[v], i)
		maxV = max(maxV, v)
	}

	for i := range s {
		if s[i] == '0' {
			if ids, ok := cnt[0]; ok && res[ids[0]] == nil {
				for _, id := range ids {
					res[id] = []int{i, i}
				}
			}
			continue
		}
		sum := 0
		for j := i; j < len(s); j++ {
			sum = sum*2 + int(s[j]-'0')
			if ids, ok := cnt[sum]; ok && res[ids[0]] == nil {
				for _, id := range ids {
					res[id] = []int{i, j}
				}
			}
			if sum > maxV {
				break
			}
		}
	}
	for i := range res {
		if res[i] == nil {
			res[i] = []int{-1, -1}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
