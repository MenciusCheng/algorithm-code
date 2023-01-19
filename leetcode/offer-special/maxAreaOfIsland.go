package offer_special

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	visited := make(map[[2]int]bool)
	ds := [][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

	var dfs func(p [2]int) int
	dfs = func(p [2]int) int {
		if visited[p] {
			return 0
		}
		visited[p] = true
		v := 1
		for _, d := range ds {
			i := p[0] + d[0]
			j := p[1] + d[1]
			if i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == 1 && !visited[[2]int{i, j}] {
				v += dfs([2]int{i, j})
			}
		}
		return v
	}

	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				res = max(res, dfs([2]int{i, j}))
			}
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
