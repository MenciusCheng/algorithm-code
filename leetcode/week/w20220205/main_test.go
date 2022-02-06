package main

import (
	"reflect"
	"testing"
)

func Test_minimumSum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				num: 2932,
			},
			want: 52,
		},
		{
			args: args{
				num: 4009,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSum(tt.args.num); got != tt.want {
				t.Errorf("minimumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pivotArray(t *testing.T) {
	type args struct {
		nums  []int
		pivot int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums:  []int{9, 12, 5, 10, 14, 3, 10},
				pivot: 10,
			},
			want: []int{9, 5, 3, 10, 10, 12, 14},
		},
		{
			args: args{
				nums:  []int{-3, 4, 3, 2},
				pivot: 2,
			},
			want: []int{-3, 2, 4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotArray(tt.args.nums, tt.args.pivot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pivotArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minCostSetTime(t *testing.T) {
	type args struct {
		startAt       int
		moveCost      int
		pushCost      int
		targetSeconds int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				startAt:       1,
				moveCost:      2,
				pushCost:      1,
				targetSeconds: 600,
			},
			want: 6,
		},
		{
			args: args{
				startAt:       0,
				moveCost:      1,
				pushCost:      2,
				targetSeconds: 76,
			},
			want: 6,
		},
		{
			args: args{
				startAt:       0,
				moveCost:      1,
				pushCost:      4,
				targetSeconds: 9,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostSetTime(tt.args.startAt, tt.args.moveCost, tt.args.pushCost, tt.args.targetSeconds); got != tt.want {
				t.Errorf("minCostSetTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumDifference(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				nums: []int{3, 1, 2},
			},
			want: -1,
		},
		{
			args: args{
				nums: []int{7, 9, 5, 8, 1, 3},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDifference(tt.args.nums); got != tt.want {
				t.Errorf("minimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
