package w20230115

func countGood(nums []int, k int) int64 {
	var res int64
	var left, nk int
	cnt := make(map[int]int)
	for i, num := range nums {
		if cnt[num] >= 1 {
			nk += cnt[num]
		}
		cnt[num]++
		for nk >= k {
			res += int64(len(nums) - i)
			cnt[nums[left]]--
			if cnt[nums[left]] >= 1 {
				nk -= cnt[nums[left]]
			}
			left++
		}
	}
	return res
}

func maxOutput(n int, edges [][]int, price []int) int64 {
	medge := make(map[int][]int)
	for _, edge := range edges {
		medge[edge[0]] = append(medge[edge[0]], edge[1])
		medge[edge[1]] = append(medge[edge[1]], edge[0])
	}

	mvalue := make(map[int][]int)
	var calPrice func(i int, fa int) int
	calPrice = func(i int, fa int) int {
		if len(mvalue[i]) == 0 {
			mvalue[i] = make([]int, len(medge[i]))
		}
		v := 0
		for j, ii := range medge[i] {
			if ii == fa {
				continue
			}
			if mvalue[i][j] == 0 {
				mvalue[i][j] = calPrice(ii, i)
			}
			v = max(v, mvalue[i][j])
		}
		return price[i] + v
	}

	res := 0
	for k, v := range medge {
		if len(v) > 1 {
			continue
		}
		res = max(res, calPrice(v[0], k))
	}

	return int64(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
