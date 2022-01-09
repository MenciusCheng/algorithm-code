package w20220108

import "testing"

func Test_capitalizeTitle(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{title: "capiTalIze tHe titLe"},
			want: "Capitalize The Title",
		},
		{
			args: args{title: "First leTTeR of EACH Word"},
			want: "First Letter of Each Word",
		},
		{
			args: args{title: "i lOve leetcode"},
			want: "i Love Leetcode",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := capitalizeTitle(tt.args.title); got != tt.want {
				t.Errorf("capitalizeTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{words: []string{"lc", "cl", "gg"}},
			want: 6,
		},
		{
			args: args{words: []string{"ab", "ty", "yt", "lc", "cl", "ab"}},
			want: 8,
		},
		{
			args: args{words: []string{"cc", "ll", "xx"}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.words); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_possibleToStamp(t *testing.T) {
	type args struct {
		grid        [][]int
		stampHeight int
		stampWidth  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				grid:        [][]int{{1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}},
				stampHeight: 4,
				stampWidth:  3,
			},
			want: true,
		},
		{
			args: args{
				grid:        [][]int{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}},
				stampHeight: 2,
				stampWidth:  2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleToStamp(tt.args.grid, tt.args.stampHeight, tt.args.stampWidth); got != tt.want {
				t.Errorf("possibleToStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
