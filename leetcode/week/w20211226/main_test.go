package w20211226

import (
	"reflect"
	"testing"
)

func Test_isSameAfterReversals(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{num: 526},
			want: true,
		},
		{
			args: args{num: 1800},
			want: false,
		},
		{
			args: args{num: 0},
			want: true,
		},
		{
			args: args{num: 609576},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameAfterReversals(tt.args.num); got != tt.want {
				t.Errorf("isSameAfterReversals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_executeInstructions(t *testing.T) {
	type args struct {
		n        int
		startPos []int
		s        string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				n:        3,
				startPos: []int{0, 1},
				s:        "RRDDLU",
			},
			want: []int{1, 5, 4, 3, 1, 0},
		},
		{
			args: args{
				n:        2,
				startPos: []int{1, 1},
				s:        "LURD",
			},
			want: []int{4, 1, 0, 0},
		},
		{
			args: args{
				n:        1,
				startPos: []int{0, 0},
				s:        "LRUD",
			},
			want: []int{0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := executeInstructions(tt.args.n, tt.args.startPos, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("executeInstructions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDistances(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			args: args{
				arr: []int{2, 1, 3, 1, 2, 3, 3},
			},
			want: []int64{4, 2, 7, 2, 4, 4, 5},
		},
		{
			args: args{
				arr: []int{10, 5, 10, 10},
			},
			want: []int64{5, 0, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDistances(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDistances() = %v, want %v", got, tt.want)
			}
		})
	}
}
