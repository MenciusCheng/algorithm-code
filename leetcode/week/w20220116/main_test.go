package w20220116

import (
	"reflect"
	"testing"
)

func Test_divideString(t *testing.T) {
	type args struct {
		s    string
		k    int
		fill byte
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				s:    "abcdefghi",
				k:    3,
				fill: 'x',
			},
			want: []string{
				"abc", "def", "ghi",
			},
		},
		{
			args: args{
				s:    "abcdefghij",
				k:    3,
				fill: 'x',
			},
			want: []string{"abc", "def", "ghi", "jxx"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divideString(tt.args.s, tt.args.k, tt.args.fill); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divideString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minMoves(t *testing.T) {
	type args struct {
		target     int
		maxDoubles int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	args: args{
		//		target:     5,
		//		maxDoubles: 0,
		//	},
		//	want: 4,
		//},
		{
			args: args{
				target:     19,
				maxDoubles: 2,
			},
			want: 7,
		},
		{
			args: args{
				target:     10,
				maxDoubles: 4,
			},
			want: 4,
		},
		//{
		//	args: args{
		//		target:     766972377,
		//		maxDoubles: 92,
		//	},
		//	want: 4,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoves(tt.args.target, tt.args.maxDoubles); got != tt.want {
				t.Errorf("minMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mostPoints(t *testing.T) {
	type args struct {
		questions [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{questions: [][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}},
			want: 5,
		},
		{
			args: args{questions: [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}},
			want: 7,
		},
		{
			args: args{questions: [][]int{{21, 5}, {92, 3}, {74, 2}, {39, 4}, {58, 2}, {5, 5}, {49, 4}, {65, 3}}},
			want: 157,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostPoints(tt.args.questions); got != tt.want {
				t.Errorf("mostPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
