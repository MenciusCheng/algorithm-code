package sword

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{[]int{1, 2, 3, 1}},
			want: 4,
		},
		{
			args: args{[]int{2, 7, 9, 3, 1}},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{[]int{1, 5, 11, 5}},
			want: true,
		},
		{
			args: args{[]int{1, 2, 3, 5}},
			want: false,
		},
		{
			args: args{[]int{1, 2, 5}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition(tt.args.nums); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	args: args{
		//		nums:   []int{1, 1, 1, 1, 1},
		//		target: 3,
		//	},
		//	want: 5,
		//},
		//{
		//	args: args{
		//		nums:   []int{1},
		//		target: 1,
		//	},
		//	want: 1,
		//},
		//{
		//	args: args{
		//		nums:   []int{1},
		//		target: 2,
		//	},
		//	want: 0,
		//},
		{
			args: args{
				nums:   []int{0, 0, 0, 0, 0, 0, 0, 0, 1},
				target: 1,
			},
			want: 256,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
