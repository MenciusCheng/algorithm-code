package w20240519

import (
	"reflect"
	"testing"
)

func Test_sumDigitDifferences(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{nums: []int{50, 28, 48}},
			want: 5,
		}, {
			args: args{nums: []int{10, 10, 10, 10}},
			want: 0,
		}, {
			args: args{nums: []int{13, 23, 12}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumDigitDifferences(tt.args.nums); got != tt.want {
				t.Errorf("sumDigitDifferences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isArraySpecial(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			args: args{
				nums:    []int{1, 1},
				queries: [][]int{{0, 1}},
			},
			want: []bool{false},
		},
		{
			args: args{
				nums:    []int{3, 4, 1, 2, 6},
				queries: [][]int{{0, 4}},
			},
			want: []bool{false},
		},
		{
			args: args{
				nums:    []int{4, 3, 1, 6},
				queries: [][]int{{0, 2}, {2, 3}},
			},
			want: []bool{false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isArraySpecial(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isArraySpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}
