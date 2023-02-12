package w20230205

import (
	"reflect"
	"testing"
)

func Test_pickGifts(t *testing.T) {
	type args struct {
		gifts []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				gifts: []int{25, 64, 9, 4, 100},
				k:     4,
			},
			want: 29,
		},
		{
			args: args{
				gifts: []int{1, 1, 1, 1},
				k:     4,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pickGifts(tt.args.gifts, tt.args.k); got != tt.want {
				t.Errorf("pickGifts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vowelStrings(t *testing.T) {
	type args struct {
		words   []string
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				words:   []string{"aba", "bcb", "ece", "aa", "e"},
				queries: [][]int{{0, 2}, {1, 4}, {1, 1}},
			},
			want: []int{2, 3, 0},
		},
		{
			args: args{
				words:   []string{"a", "e", "i"},
				queries: [][]int{{0, 2}, {0, 1}, {2, 2}},
			},
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vowelStrings(tt.args.words, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("vowelStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minCapability(t *testing.T) {
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
				nums: []int{2, 3, 5, 9},
				k:    2,
			},
			want: 5,
		},
		{
			args: args{
				nums: []int{2, 7, 9, 3, 1},
				k:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCapability(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minCapability() = %v, want %v", got, tt.want)
			}
		})
	}
}
