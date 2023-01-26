package offer_special

func alienOrder(words []string) string {
	graph := make(map[byte]map[byte]bool)
	indegree := make(map[byte]int)

	cnt := make(map[byte]bool)
	pm := 0
	for _, w := range words {
		pm = max(pm, len(w))
		for i := range w {
			cnt[w[i]] = true
		}
	}
	for k := range cnt {
		graph[k] = make(map[byte]bool)
		indegree[k] = 0
	}

	for i := 0; i < len(words)-1; i++ {
		w1 := words[i]
		for j := i + 1; j < len(words); j++ {
			w2 := words[j]
			l := min(len(w1), len(w2))
			if w1[:l] == w2[:l] && len(w1) > len(w2) {
				return ""
			}
		}
	}

	for k := 0; k < pm; k++ {
		for i := 0; i < len(words)-1; i++ {
			w1 := words[i]
			for j := i + 1; j < len(words); j++ {
				w2 := words[j]
				if k < len(w1) && k < len(w2) && w1[:k] == w2[:k] && w1[k] != w2[k] && !graph[w1[k]][w2[k]] {
					graph[w1[k]][w2[k]] = true
					indegree[w2[k]]++
				}
			}
		}
	}

	queue := make([]byte, 0)
	for k, v := range indegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}
	res := make([]byte, 0)
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		res = append(res, q)
		for ch := range graph[q] {
			indegree[ch]--
			if indegree[ch] == 0 {
				queue = append(queue, ch)
			}
		}
	}
	if len(res) != len(indegree) {
		return ""
	}
	return string(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
