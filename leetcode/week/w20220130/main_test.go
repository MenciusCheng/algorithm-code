package w20220130

import (
	"reflect"
	"testing"
)

func Test_findFinalValue(t *testing.T) {
	type args struct {
		nums     []int
		original int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:     []int{5, 3, 6, 1, 12},
				original: 3,
			},
			want: 24,
		},
		{
			args: args{
				nums:     []int{2, 7, 9},
				original: 4,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFinalValue(tt.args.nums, tt.args.original); got != tt.want {
				t.Errorf("findFinalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxScoreIndices(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{[]int{0, 0, 1, 0}},
			want: []int{2, 4},
		},
		{
			args: args{[]int{0, 0, 0}},
			want: []int{3},
		},
		{
			args: args{[]int{1, 1}},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScoreIndices(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxScoreIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subStrHash(t *testing.T) {
	type args struct {
		s         string
		power     int
		modulo    int
		k         int
		hashValue int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s:         "leetcode",
				power:     7,
				modulo:    20,
				k:         2,
				hashValue: 0,
			},
			want: "ee",
		},
		{
			args: args{
				s:         "fbxzaad",
				power:     31,
				modulo:    100,
				k:         3,
				hashValue: 32,
			},
			want: "fbx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subStrHash(tt.args.s, tt.args.power, tt.args.modulo, tt.args.k, tt.args.hashValue); got != tt.want {
				t.Errorf("subStrHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
