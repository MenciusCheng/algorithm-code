package w20220130

import "sort"

func findFinalValue(nums []int, original int) int {
	exist := false
	m := make([]int, 0)
	for _, num := range nums {
		if num == original {
			exist = true
		}
		if num%original == 0 {
			m = append(m, num)
		}
	}
	if !exist {
		return original
	}

	sort.Ints(m)
	for i := 0; i < len(m); i++ {
		if original == m[i] {
			original *= 2
		}
	}
	return original
}

func maxScoreIndices(nums []int) []int {
	res := []int{0}
	right := 0
	for _, num := range nums {
		if num == 1 {
			right++
		}
	}
	max := right

	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			left++
		} else {
			right--
		}
		sum := left + right
		if sum > max {
			max = sum
			res = []int{i + 1}
		} else if sum == max {
			res = append(res, i+1)
		}
	}

	return res
}

func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	arr := make([]byte, 0, k)
	for i := 0; i < k; i++ {
		arr = append(arr, s[i])
	}
	power %= modulo

	vs := make([]int, k)
	for i := 0; i < k; i++ {
		if i > 0 {
			vs[i] = (vs[i-1] * power) % modulo
		}
		vs[i] = (vs[i] + int(arr[k-i-1]-'a'+1)) % modulo
	}
	if vs[len(vs)-1] == hashValue {
		return string(arr)
	}

	for i := k; i < len(s); i++ {
		bs := arr[1:]
		arr = append(bs, s[i])

		nvs := make([]int, k)
		nvs[0] = int(arr[k-1]-'a'+1) % modulo
		for j := 1; j < k; j++ {
			nvs[j] = (nvs[j-1] * power) % modulo
		}
		for j := 1; j < k; j++ {
			nvs[j] = (nvs[j] + vs[j-1]) % modulo
		}
		if nvs[len(nvs)-1] == hashValue {
			return string(arr)
		}
		vs = nvs
	}

	return ""
}
