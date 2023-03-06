package w20230226

import "sort"

func divisibilityArray(word string, m int) []int {
	res := make([]int, len(word))
	v := 0
	for i := range word {
		v = v*10 + int(word[i]-'0')
		v = v % m
		if v == 0 {
			res[i] = 1
		}
	}
	return res
}

func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	cnt := make(map[int]bool)
	l := len(nums) - 2
	for j := len(nums) - 1; j >= 0 && l >= 0; j-- {
		if cnt[j] {
			continue
		}
		v := nums[j] / 2
		for l >= 0 && nums[l] > v {
			l--
		}
		if l >= 0 {
			cnt[j] = true
			cnt[l] = true
			l--
		}
	}
	return len(cnt)
}

func minimumTime(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	fin := [2]int{m - 1, n - 1}
	t := 0
	p := map[[2]int]bool{
		[2]int{0, 0}: true,
	}
	for len(p) > 0 {
		t++
		next := make(map[[2]int]bool)
		for k := range p {
			if k == fin {
				return t - 1
			}
			i := k[0]
			j := k[1] + 1
			if i >= 0 && i < m && j >= 0 && j < n && t >= grid[i][j] {
				next[[2]int{i, j}] = true
			}
			i = k[0] + 1
			j = k[1]
			if i >= 0 && i < m && j >= 0 && j < n && t >= grid[i][j] {
				next[[2]int{i, j}] = true
			}
			i = k[0] - 1
			j = k[1]
			if i >= 0 && i < m && j >= 0 && j < n && t >= grid[i][j] {
				next[[2]int{i, j}] = true
			}
			i = k[0]
			j = k[1] - 1
			if i >= 0 && i < m && j >= 0 && j < n && t >= grid[i][j] {
				next[[2]int{i, j}] = true
			}
		}
		p = next
	}
	return -1
}
