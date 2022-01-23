package w20220123

import (
	"reflect"
	"testing"
)

func Test_countElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{11, 7, 2, 15},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{-3, 3, 3, 90},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countElements(tt.args.nums); got != tt.want {
				t.Errorf("countElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rearrangeArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{nums: []int{3, 1, -2, -5, 2, -4}},
			want: []int{3, -2, 1, -5, 2, -4},
		},
		{
			args: args{nums: []int{-1, 1}},
			want: []int{1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rearrangeArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rearrangeArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLonely(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{[]int{10, 6, 5, 8}},
			want: []int{8, 10},
		},
		{
			args: args{[]int{1, 3, 5, 3}},
			want: []int{1, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLonely(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLonely() = %v, want %v", got, tt.want)
			}
		})
	}
}
