package w20220305

import (
	"reflect"
	"testing"
)

func Test_mostFrequent(t *testing.T) {
	type args struct {
		nums []int
		key  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{1, 100, 200, 1, 100},
				key:  1,
			},
			want: 100,
		},
		{
			args: args{
				nums: []int{2, 2, 2, 2, 3},
				key:  2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostFrequent(tt.args.nums, tt.args.key); got != tt.want {
				t.Errorf("mostFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cellsInRange(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				s: "K1:L2",
			},
			want: []string{"K1", "K2", "L1", "L2"},
		},
		{
			args: args{
				s: "A1:F1",
			},
			want: []string{"A1", "B1", "C1", "D1", "E1", "F1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cellsInRange(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cellsInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimalKSum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				nums: []int{1, 4, 25, 10, 25},
				k:    2,
			},
			want: 5,
		},
		{
			args: args{
				nums: []int{5, 6},
				k:    6,
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimalKSum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minimalKSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createBinaryTree(t *testing.T) {
	type args struct {
		descriptions [][]int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{
				descriptions: [][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}},
			},
		},
		{
			args: args{
				descriptions: [][]int{{1, 2, 1}, {2, 3, 0}, {3, 4, 1}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createBinaryTree(tt.args.descriptions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
