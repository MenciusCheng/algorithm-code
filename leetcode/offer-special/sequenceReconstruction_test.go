package offer_special

import "testing"

func Test_sequenceReconstruction(t *testing.T) {
	type args struct {
		nums      []int
		sequences [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				nums:      []int{4, 1, 5, 2, 6, 3},
				sequences: [][]int{{5, 2, 6, 3}, {4, 1, 5, 2}},
			},
			want: true,
		},
		{
			args: args{
				nums:      []int{1, 2, 3},
				sequences: [][]int{{1, 2}, {1, 3}, {2, 3}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sequenceReconstruction(tt.args.nums, tt.args.sequences); got != tt.want {
				t.Errorf("sequenceReconstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}
