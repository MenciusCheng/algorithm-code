package w20230108

import "testing"

func Test_maxKelements(t *testing.T) {
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
				nums: []int{756902131, 995414896, 95906472, 149914376, 387433380, 848985151},
				k:    6,
			},
			want: 3603535575,
		},
		{
			args: args{
				nums: []int{10, 10, 10, 10, 10},
				k:    5,
			},
			want: 50,
		},
		{
			args: args{
				nums: []int{1, 10, 3, 3, 3},
				k:    3,
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxKelements(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxKelements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findCrossingTime(t *testing.T) {
	type args struct {
		n    int
		k    int
		time [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n:    1,
				k:    3,
				time: [][]int{{1, 1, 2, 1}, {1, 1, 3, 1}, {1, 1, 4, 1}},
			},
			want: 6,
		},
		{
			args: args{
				n:    3,
				k:    2,
				time: [][]int{{1, 9, 1, 8}, {10, 10, 10, 10}},
			},
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCrossingTime(tt.args.n, tt.args.k, tt.args.time); got != tt.want {
				t.Errorf("findCrossingTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
