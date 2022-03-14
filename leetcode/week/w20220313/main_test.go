package w20220313

import (
	"reflect"
	"testing"
)

func Test_findKDistantIndices(t *testing.T) {
	type args struct {
		nums []int
		key  int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{3, 4, 9, 1, 3, 9, 5},
				key:  9,
				k:    1,
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			args: args{
				nums: []int{2, 2, 2, 2, 2},
				key:  2,
				k:    2,
			},
			want: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKDistantIndices(tt.args.nums, tt.args.key, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findKDistantIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_digArtifacts(t *testing.T) {
	type args struct {
		n         int
		artifacts [][]int
		dig       [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n:         2,
				artifacts: [][]int{{0, 0, 0, 0}, {0, 1, 1, 1}},
				dig:       [][]int{{0, 0}, {0, 1}},
			},
			want: 1,
		},
		{
			args: args{
				n:         2,
				artifacts: [][]int{{0, 0, 0, 0}, {0, 1, 1, 1}},
				dig:       [][]int{{0, 0}, {0, 1}, {1, 1}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digArtifacts(tt.args.n, tt.args.artifacts, tt.args.dig); got != tt.want {
				t.Errorf("digArtifacts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumTop(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{5, 2, 2, 4, 0, 6},
				k:    4,
			},
			want: 5,
		},
		{
			args: args{
				nums: []int{2},
				k:    1,
			},
			want: -1,
		},
		{
			args: args{
				nums: []int{35, 43, 23, 86, 23, 45, 84, 2, 18, 83, 79, 28, 54, 81, 12, 94, 14, 0, 0, 29, 94, 12, 13, 1, 48, 85, 22, 95, 24, 5, 73, 10, 96, 97, 72, 41, 52, 1, 91, 3, 20, 22, 41, 98, 70, 20, 52, 48, 91, 84, 16, 30, 27, 35, 69, 33, 67, 18, 4, 53, 86, 78, 26, 83, 13, 96, 29, 15, 34, 80, 16, 49},
				k:    15,
			},
			want: 94,
		},
		{
			args: args{
				nums: []int{91, 98, 17, 79, 15, 55, 47, 86, 4, 5, 17, 79, 68, 60, 60, 31, 72, 85, 25, 77, 8, 78, 40, 96, 76, 69, 95, 2, 42, 87, 48, 72, 45, 25, 40, 60, 21, 91, 32, 79, 2, 87, 80, 97, 82, 94, 69, 43, 18, 19, 21, 36, 44, 81, 99},
				k:    2,
			},
			want: 91,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTop(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumTop() = %v, want %v", got, tt.want)
			}
		})
	}
}
