package w20230101

import (
	"math"
	"strconv"
)

func minimumPartition(s string, k int) int {
	cnt := make(map[string]int)

	var dfs func(ws, ks string) int
	dfs = func(ws, ks string) int {
		if _, ok := cnt[ws]; ok {
			return cnt[ws]
		}

		if len(ws) == 0 {
			return 0
		} else if len(ws) < len(ks) {
			cnt[ws] = 1
			return cnt[ws]
		}

		m := math.MaxInt
		if ws[:len(ks)] <= ks {
			t := dfs(ws[len(ks):], ks)
			if t != -1 && t < m {
				m = t
			}
		}
		if len(ks) >= 2 {
			t := dfs(ws[len(ks)-1:], ks)
			if t != -1 && t < m {
				m = t
			}
		}
		if m == math.MaxInt {
			cnt[ws] = -1
		} else {
			cnt[ws] = m + 1
		}

		return cnt[ws]
	}
	dfs(s, strconv.Itoa(k))

	return cnt[s]
}
