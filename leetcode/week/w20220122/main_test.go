package w20220122

import (
	"reflect"
	"testing"
)

func Test_minimumCost(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{[]int{1, 2, 3}},
			want: 5,
		},
		{
			args: args{[]int{6, 5, 7, 9, 2, 2}},
			want: 23,
		},
		{
			args: args{[]int{5, 5}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumCost(tt.args.cost); got != tt.want {
				t.Errorf("minimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberOfArrays(t *testing.T) {
	type args struct {
		differences []int
		lower       int
		upper       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				differences: []int{1, -3, 4},
				lower:       1,
				upper:       6,
			},
			want: 2,
		},
		{
			args: args{
				differences: []int{3, -4, 5, 1, -2},
				lower:       -4,
				upper:       5,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArrays(tt.args.differences, tt.args.lower, tt.args.upper); got != tt.want {
				t.Errorf("numberOfArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_highestRankedKItems(t *testing.T) {
	type args struct {
		grid    [][]int
		pricing []int
		start   []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				grid: [][]int{
					{1, 2, 0, 1}, {1, 3, 0, 1}, {0, 2, 5, 1},
				},
				pricing: []int{2, 5},
				start:   []int{0, 0},
				k:       3,
			},
			want: [][]int{
				{0, 1}, {1, 1}, {2, 1},
			},
		},
		{
			args: args{
				grid: [][]int{
					{1, 2, 0, 1}, {1, 3, 3, 1}, {0, 2, 5, 1},
				},
				pricing: []int{2, 3},
				start:   []int{2, 3},
				k:       2,
			},
			want: [][]int{
				{2, 1}, {1, 2},
			},
		},
		{
			args: args{
				grid: [][]int{
					{1, 1, 1}, {0, 0, 1}, {2, 3, 4},
				},
				pricing: []int{2, 3},
				start:   []int{0, 0},
				k:       3,
			},
			want: [][]int{
				{2, 1}, {2, 0},
			},
		},
		{
			args: args{
				grid: [][]int{
					{0, 2, 0},
				},
				pricing: []int{2, 2},
				start:   []int{0, 1},
				k:       1,
			},
			want: [][]int{
				{0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := highestRankedKItems(tt.args.grid, tt.args.pricing, tt.args.start, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("highestRankedKItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
