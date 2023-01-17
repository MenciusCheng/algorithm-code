package w20230115

import "testing"

func Test_countGood(t *testing.T) {
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
				nums: []int{1, 1, 1, 1, 1},
				k:    10,
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{3, 1, 4, 3, 2, 2, 4},
				k:    2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGood(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countGood() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxOutput(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
		price []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				n:     6,
				edges: [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {3, 5}},
				price: []int{9, 8, 7, 6, 10, 5},
			},
			want: 24,
		},
		{
			args: args{
				n:     3,
				edges: [][]int{{0, 1}, {1, 2}},
				price: []int{1, 1, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxOutput(tt.args.n, tt.args.edges, tt.args.price); got != tt.want {
				t.Errorf("maxOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}
