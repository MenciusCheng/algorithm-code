package w20230304

func splitNum(num int) int {
	cnt := [9]int{}
	for num > 0 {
		v := num % 10
		if v > 0 {
			cnt[v-1]++
		}
		num /= 10
	}
	var a, b int
	t := 1
	i := 8
	for i >= 0 {
		if cnt[i] == 0 {
			i--
		} else if cnt[i] >= 2 {
			a += t * (i + 1)
			b += t * (i + 1)
			cnt[i] -= 2
			t *= 10
		} else {
			a += t * (i + 1)
			cnt[i]--
			for i >= 0 && cnt[i] == 0 {
				i--
			}
			if i < 0 {
				break
			}
			b += t * (i + 1)
			cnt[i]--
			t *= 10
		}

	}
	return a + b
}
