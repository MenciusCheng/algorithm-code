package w20220319

import "testing"

func Test_divideArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				nums: []int{3, 2, 3, 2, 2, 2},
			},
			want: true,
		},
		{
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divideArray(tt.args.nums); got != tt.want {
				t.Errorf("divideArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumSubsequenceCount(t *testing.T) {
	type args struct {
		text    string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				text:    "abdcdbc",
				pattern: "ac",
			},
			want: 4,
		},
		{
			args: args{
				text:    "aabb",
				pattern: "ab",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSubsequenceCount(tt.args.text, tt.args.pattern); got != tt.want {
				t.Errorf("maximumSubsequenceCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
