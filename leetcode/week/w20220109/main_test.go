package w20220109

import "testing"

func Test_checkValid(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{matrix: [][]int{{1, 2, 3}, {3, 1, 2}, {2, 3, 1}}},
			want: true,
		},
		{
			args: args{matrix: [][]int{{1, 1, 1}, {1, 2, 3}, {1, 2, 3}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValid(tt.args.matrix); got != tt.want {
				t.Errorf("checkValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wordCount(t *testing.T) {
	type args struct {
		startWords  []string
		targetWords []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				startWords:  []string{"ant", "act", "tack"},
				targetWords: []string{"tack", "act", "acti"},
			},
			want: 2,
		},
		{
			args: args{
				startWords:  []string{"ab", "a"},
				targetWords: []string{"abc", "abcd"},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordCount(tt.args.startWords, tt.args.targetWords); got != tt.want {
				t.Errorf("wordCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
