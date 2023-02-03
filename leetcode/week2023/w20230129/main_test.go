package w20230129

import "testing"

func Test_monkeyMove(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: 500000003,
			},
			want: 1000000006,
		},
		{
			args: args{
				n: 55,
			},
			want: 766762394,
		},
		{
			args: args{
				n: 4,
			},
			want: 14,
		},
		{
			args: args{
				n: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := monkeyMove(tt.args.n); got != tt.want {
				t.Errorf("monkeyMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_putMarbles(t *testing.T) {
	type args struct {
		weights []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				weights: []int{1, 3, 5, 1},
				k:       2,
			},
			want: 4,
		},
		{
			args: args{
				weights: []int{1, 3},
				k:       2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := putMarbles(tt.args.weights, tt.args.k); got != tt.want {
				t.Errorf("putMarbles() = %v, want %v", got, tt.want)
			}
		})
	}
}
