package w20220206

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_sortEvenOdd(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{4, 1, 2, 3},
			},
			want: []int{2, 3, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortEvenOdd(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortEvenOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_smallestNumber(t *testing.T) {
	type args struct {
		num int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				num: 310,
			},
			want: 103,
		},
		{
			args: args{
				num: -7605,
			},
			want: -7650,
		},
		{
			args: args{
				num: 4099,
			},
			want: 4099,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestNumber(tt.args.num); got != tt.want {
				t.Errorf("smallestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor(t *testing.T) {
	bs := Constructor(5) // bitset = "00000".
	fmt.Println(bs.ToString())
	bs.Fix(3) // 将 idx = 3 处的值更新为 1 ，此时 bitset = "00010" 。
	fmt.Println(bs.ToString())
	bs.Fix(1) // 将 idx = 1 处的值更新为 1 ，此时 bitset = "01010" 。
	fmt.Println(bs.ToString())
	bs.Flip() // 翻转每一位上的值，此时 bitset = "10101" 。
	fmt.Println(bs.ToString())
	fmt.Println(bs.All() == false) // 返回 False ，bitset 中的值不全为 1 。
	bs.Unfix(0)                    // 将 idx = 0 处的值更新为 0 ，此时 bitset = "00101" 。
	fmt.Println(bs.ToString())
	bs.Flip() // 翻转每一位上的值，此时 bitset = "11010" 。
	fmt.Println(bs.ToString())
	fmt.Println(bs.One() == true) // 返回 True ，至少存在一位的值为 1 。
	bs.Unfix(0)                   // 将 idx = 0 处的值更新为 0 ，此时 bitset = "01010" 。
	fmt.Println(bs.ToString())
	fmt.Println(bs.Count() == 2) // 返回 2 ，当前有 2 位的值为 1 。
	fmt.Println(bs.ToString())   // 返回 "01010" ，即 bitset 的当前组成情况。
}
