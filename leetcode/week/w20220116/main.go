package w20220116

import "bytes"

func divideString(s string, k int, fill byte) []string {
	fillN := len(s) % k
	if fillN != 0 {
		fillN = k - fillN
	}
	resB := make([]bytes.Buffer, 0)

	b := bytes.Buffer{}
	for i := 0; i < len(s); i++ {
		b.WriteByte(s[i])
		if i%k == k-1 || i == len(s)-1 {
			resB = append(resB, b)
			b = bytes.Buffer{}
		}
	}

	for i := 0; i < fillN; i++ {
		resB[len(resB)-1].WriteByte(fill)
	}

	res := make([]string, 0, len(resB))
	for i := 0; i < len(resB); i++ {
		res = append(res, resB[i].String())
	}
	return res
}

func minMoves(target int, maxDoubles int) int {
	step := 0
	current := target
	leftDoubles := maxDoubles
	for current > 1 {
		if leftDoubles == 0 {
			step += current - 1
			current = 1
		} else if current%2 == 0 {
			current /= 2
			step++
			leftDoubles--
		} else {
			current -= 1
			step++
		}
	}

	return step
}

func mostPoints(questions [][]int) int64 {
	maxPoints := make([]int64, len(questions))

	maxPoints[len(questions)-1] = int64(questions[len(questions)-1][0])
	for i := len(questions) - 2; i >= 0; i-- {
		next := i + questions[i][1] + 1
		current := int64(questions[i][0])

		if next > len(questions)-1 {
			if current > maxPoints[i+1] {
				maxPoints[i] = current
			} else {
				maxPoints[i] = maxPoints[i+1]
			}
		} else {
			current += maxPoints[next]
			if current > maxPoints[i+1] {
				maxPoints[i] = current
			} else {
				maxPoints[i] = maxPoints[i+1]
			}
		}
	}

	return maxPoints[0]
}
