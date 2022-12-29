package offer_special

import "math"

func divide(a int, b int) int {
	nag := false
	if a > 0 {
		a = -a
		nag = !nag
	}
	if b > 0 {
		b = -b
		nag = !nag
	}
	if a > b {
		return 0
	}
	res := 0
	c := 1
	for a <= (b << 1) {
		b <<= 1
		c <<= 1
	}
	for a < 0 && c >= 1 {
		if a <= b {
			a -= b
			res += c
		}
		b >>= 1
		c >>= 1
	}
	if nag {
		res = -res
	}
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res
}
