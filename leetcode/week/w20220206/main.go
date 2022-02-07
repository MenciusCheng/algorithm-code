package w20220206

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

func sortEvenOdd(nums []int) []int {
	if len(nums) <= 2 {
		return nums
	}
	a, b := make([]int, 0, len(nums)/2), make([]int, 0, len(nums)/2)
	for i, num := range nums {
		if i&1 == 0 {
			a = append(a, num)
		} else {
			b = append(b, num)
		}
	}
	sort.Ints(a)
	sort.Ints(b)
	res := make([]int, len(nums))
	index := 0
	for i := 0; i < len(a); i++ {
		res[index] = a[i]
		index += 2
	}
	index = 1
	for i := len(b) - 1; i >= 0; i-- {
		res[index] = b[i]
		index += 2
	}
	return res
}

func smallestNumber(num int64) int64 {
	if num == 0 {
		return 0
	}
	if num > 0 {
		s := fmt.Sprintf("%d", num)
		nums := make([]int, 0, len(s))
		for i := range s {
			nums = append(nums, int(s[i]-'0'))
		}
		sort.Ints(nums)
		sb := make([]byte, 0)
		for i := range nums {
			sb = append(sb, byte(nums[i])+'0')
		}
		if sb[0] == '0' {
			for i := 1; i < len(sb); i++ {
				if sb[i] != '0' {
					sb[0], sb[i] = sb[i], sb[0]
					break
				}
			}
		}
		res, _ := strconv.Atoi(string(sb))
		return int64(res)
	} else {
		num = -num
		s := fmt.Sprintf("%d", num)
		nums := make([]int, 0, len(s))
		for i := range s {
			nums = append(nums, int(s[i]-'0'))
		}
		sort.Ints(nums)
		sb := bytes.Buffer{}
		for i := len(nums) - 1; i >= 0; i-- {
			sb.WriteByte(byte(nums[i]) + '0')
		}
		res, _ := strconv.Atoi(sb.String())
		return int64(-res)
	}
}

type Bitset struct {
	arr  []byte
	size int
	one  int
	flip bool
}

func Constructor(size int) Bitset {
	b := Bitset{}
	b.arr = make([]byte, size)
	for i := range b.arr {
		b.arr[i] = '0'
	}
	b.size = size
	return b
}

func (this *Bitset) Fix(idx int) {
	if !this.flip {
		if this.arr[idx] != '1' {
			this.arr[idx] = '1'
			this.one++
		}
	} else {
		if this.arr[idx] != '0' {
			this.arr[idx] = '0'
			this.one++
		}
	}
}

func (this *Bitset) Unfix(idx int) {
	if !this.flip {
		if this.arr[idx] != '0' {
			this.arr[idx] = '0'
			this.one--
		}
	} else {
		if this.arr[idx] != '1' {
			this.arr[idx] = '1'
			this.one--
		}
	}
}

func (this *Bitset) Flip() {
	this.flip = !this.flip
	this.one = this.size - this.one
}

func (this *Bitset) All() bool {
	return this.one == this.size
}

func (this *Bitset) One() bool {
	return this.one > 0
}

func (this *Bitset) Count() int {
	return this.one
}

func (this *Bitset) ToString() string {
	if !this.flip {
		return string(this.arr)
	} else {
		res := make([]byte, len(this.arr))
		for i := 0; i < len(this.arr); i++ {
			if this.arr[i] == '0' {
				res[i] = '1'
			} else {
				res[i] = '0'
			}
		}
		return string(res)
	}
}
