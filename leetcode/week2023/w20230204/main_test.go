package w20230204

import "testing"

func Test_maximizeWin(t *testing.T) {
	type args struct {
		prizePositions []int
		k              int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				prizePositions: []int{1, 1, 2, 2, 3, 3, 5},
				k:              2,
			},
			want: 7,
		},
		{
			args: args{
				prizePositions: []int{1, 2, 3, 4},
				k:              0,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximizeWin(tt.args.prizePositions, tt.args.k); got != tt.want {
				t.Errorf("maximizeWin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPossibleToCutPath(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				grid: [][]int{{1, 1, 1}, {1, 0, 0}, {1, 1, 1}},
			},
			want: true,
		},
		{
			args: args{
				grid: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPossibleToCutPath(tt.args.grid); got != tt.want {
				t.Errorf("isPossibleToCutPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
