package w20230204

func maximizeWin(prizePositions []int, k int) int {
	pi := make([]int, 0)
	pc := make([]int, 0)

	count := 0
	for i, v := range prizePositions {
		count++
		if i == len(prizePositions)-1 || v != prizePositions[i+1] {
			pi = append(pi, v)
			pc = append(pc, count)
			count = 0
		}
	}
	if len(pc) == 1 {
		return pc[0]
	}

	maxRight := make([]int, len(pc))
	j := len(pc) - 1
	sum := pc[j]
	maxRight[len(pc)-1] = sum
	for i := len(pc) - 2; i >= 1; i-- {
		sum += pc[i]
		for pi[j]-pi[i] > k {
			sum -= pc[j]
			j--
		}
		maxRight[i] = max(maxRight[i+1], sum)
	}

	maxLeft := make([]int, len(pc))
	j = 0
	sum = pc[j]

	res := sum + maxRight[1]

	maxLeft[0] = sum
	for i := 1; i <= len(pc)-2; i++ {
		sum += pc[i]
		for pi[i]-pi[j] > k {
			sum -= pc[j]
			j++
		}
		maxLeft[i] = max(maxLeft[i-1], sum)
		res = max(res, maxLeft[i]+maxRight[i+1])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isPossibleToCutPath(grid [][]int) bool {
	m := len(grid)
	n := len(grid[0])
	fin := [2]int{m - 1, n - 1}
	fcount := 0
	queue := [][2]int{{0, 0}}
	visited := make(map[[2]int]bool)
	for len(queue) > 0 {
		next := make([][2]int, 0)
		for _, k := range queue {
			r := [2]int{k[0], k[1] + 1}
			if k[1]+1 < n && grid[k[0]][k[1]+1] == 1 && !visited[r] {
				if r != fin {
					next = append(next, r)
				} else {
					fcount++
				}
			}
			d := [2]int{k[0] + 1, k[1]}
			if d[0] < m && grid[d[0]][d[1]] == 1 && !visited[d] {
				visited[d] = true
				if d != fin {
					next = append(next, d)
				} else {
					fcount++
				}
			}
		}
		if len(next) <= 1 {
			return fcount <= 0
		}
		queue = next
	}
	return false
}
