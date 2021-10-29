package main

import "fmt"

// https://leetcode-cn.com/problems/reordered-power-of-2/
// 869. 重新排序得到 2 的幂
func main() {
	fmt.Println(reorderedPowerOf2(1))
	fmt.Println(reorderedPowerOf2(10))
	fmt.Println(reorderedPowerOf2(16))
	fmt.Println(reorderedPowerOf2(24))
	fmt.Println(reorderedPowerOf2(46))
}

func reorderedPowerOf2(n int) bool {
	return testMap[countKey(n)]
}

var testMap = initMap(10e9)

func countKey(n int) (arr [10]int) {
	for n > 0 {
		arr[n%10] += 1
		n /= 10
	}
	return
}

func initMap(max int) map[[10]int]bool {
	res := make(map[[10]int]bool)

	for i := 1; i <= max; i *= 2 {
		res[countKey(i)] = true
	}

	return res
}
