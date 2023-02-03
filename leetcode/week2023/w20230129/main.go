package w20230129

import "math"

func monkeyMove(n int) int {
	mod := int(1e9 + 7)
	cnt := make(map[int]int)

	c := 1
	v := 2
	cnt[c] = v
	for c < n {
		c = c * 2
		v = (v * v) % mod
		cnt[c] = v
	}

	res := 1
	for n > 0 {
		for c > n {
			c /= 2
		}
		n -= c
		if n == 0 {
			res = (res*cnt[c] - 2) % mod
		} else {
			res = (res * cnt[c]) % mod
		}
	}

	return res
}

func putMarbles(weights []int, k int) int64 {
	maxm := make(map[[2]int]int)
	var dfsMax func(j, k2 int) int
	dfsMax = func(j, k2 int) int {
		res := 0
		for j2 := j; j2 >= k2-1 && j2 >= 0; j2-- {
			v := weights[j] + weights[j2] + dfsMax(j2-1, k2-1)
			if v > res {
				res = v
			}
		}
		maxm[[2]int{j, k}] = res
		return res
	}
	maxRes := dfsMax(len(weights)-1, k)

	minm := make(map[[2]int]int)
	var dfsMin func(j, k2 int) int
	dfsMin = func(j, k2 int) int {
		res := math.MaxInt
		for j2 := j; j2 >= k2-1 && j2 >= 0; j2-- {
			v := weights[j] + weights[j2] + dfsMin(j2-1, k2-1)
			if v < res {
				res = v
			}
		}
		minm[[2]int{j, k2}] = res
		return res
	}
	minRes := dfsMin(len(weights)-1, k)

	return int64(maxRes - minRes)
}
