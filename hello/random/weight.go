package random

import (
	"math/rand"
	"sort"
	"time"
)

// WeightedRandomIdx 随机选择算法
func WeightedRandomIdx(weights []int) int {
	if len(weights) == 0 {
		return -1
	}
	rand.Seed(time.Now().UnixNano())
	sum := 0
	var sumWeight []int
	for _, v := range weights {
		weight := v
		if weight < 0 {
			weight = 0
		}
		sum += weight
		sumWeight = append(sumWeight, sum)
	}
	r := rand.Intn(sum) + 1
	idx := sort.SearchInts(sumWeight, r)
	return idx
}
