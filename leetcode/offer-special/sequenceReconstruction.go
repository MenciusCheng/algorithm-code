package offer_special

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	graph := make(map[int][]int)
	indegreee := make([]int, len(nums)+1)
	for _, s := range sequences {
		for i := 0; i < len(s)-1; i++ {
			graph[s[i]] = append(graph[s[i]], s[i+1])
			indegreee[s[i+1]]++
		}
	}

	queue := make([]int, 0)
	for i := 1; i < len(indegreee); i++ {
		if indegreee[i] == 0 {
			queue = append(queue, i)
		}
	}

	i := 0
	for len(queue) > 0 {
		if len(queue) > 1 {
			return false
		}
		q := queue[0]
		queue = queue[1:]
		if nums[i] != q {
			return false
		}
		i++
		for _, v := range graph[q] {
			indegreee[v]--
			if indegreee[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return true
}
