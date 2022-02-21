package w20220219

import (
	"reflect"
	"testing"
)

func Test_countPairs(t *testing.T) {
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
				nums: []int{3, 1, 2, 2, 2, 1, 3},
				k:    2,
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{1, 2, 3, 4},
				k:    1,
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{10, 2, 3, 4, 9, 6, 3, 10, 3, 6, 3, 9, 1},
				k:    4,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumEvenSplit(t *testing.T) {
	type args struct {
		finalSum int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			args: args{
				finalSum: 12,
			},
			want: []int64{2, 4, 6},
		},
		{
			args: args{
				finalSum: 7,
			},
			want: []int64{},
		},
		{
			args: args{
				finalSum: 28,
			},
			want: []int64{2, 4, 6, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumEvenSplit(tt.args.finalSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumEvenSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
