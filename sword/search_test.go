package sword

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		ws  []int
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Target in the middle",
			args: args{
				ws:  []int{1, 3, 5, 7, 9, 11, 13, 15},
				num: 9,
			},
			want: 4,
		},
		{
			name: "Target at the end",
			args: args{
				ws:  []int{1, 3, 5, 7, 9, 11, 13, 15},
				num: 15,
			},
			want: 7,
		},
		{
			name: "Target smaller than all elements",
			args: args{
				ws:  []int{1, 3, 5, 7, 9, 11, 13, 15},
				num: 0,
			},
			want: -1,
		},
		{
			name: "Target larger than all elements",
			args: args{
				ws:  []int{1, 3, 5, 7, 9, 11, 13, 15},
				num: 20,
			},
			want: 7,
		},
		{
			name: "Target at the beginning",
			args: args{
				ws:  []int{1, 3, 5, 7, 9, 11, 13, 15},
				num: 1,
			},
			want: 0,
		},
		{
			name: "Target in between",
			args: args{
				ws:  []int{1, 3, 5, 7, 9, 11, 13, 15},
				num: 14,
			},
			want: 6,
		},
		{
			name: "Target at the second element",
			args: args{
				ws:  []int{1, 3, 5, 7, 9, 11, 13, 15},
				num: 3,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.ws, tt.args.num); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
