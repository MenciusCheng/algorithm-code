package w20211219

import "testing"

func Test_firstPalindrome(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{words: []string{"abc", "car", "ada", "racecar", "cool"}},
			want: "ada",
		},
		{
			args: args{words: []string{"def", "ghi"}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstPalindrome(tt.args.words); got != tt.want {
				t.Errorf("firstPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addSpaces(t *testing.T) {
	type args struct {
		s      string
		spaces []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s:      "LeetcodeHelpsMeLearn",
				spaces: []int{8, 13, 15},
			},
			want: "Leetcode Helps Me Learn",
		},
		{
			args: args{
				s:      "icodeinpython",
				spaces: []int{1, 5, 7, 9},
			},
			want: "i code in py thon",
		},
		{
			args: args{
				s:      "spacing",
				spaces: []int{0, 1, 2, 3, 4, 5, 6},
			},
			want: " s p a c i n g",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addSpaces(tt.args.s, tt.args.spaces); got != tt.want {
				t.Errorf("addSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDescentPeriods(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				prices: []int{3, 2, 1, 4},
			},
			want: 7,
		},
		{
			args: args{
				prices: []int{8, 6, 7, 7},
			},
			want: 4,
		},
		{
			args: args{
				prices: []int{1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDescentPeriods(tt.args.prices); got != tt.want {
				t.Errorf("getDescentPeriods() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kIncreasing(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	args: args{
		//		arr: []int{5, 4, 3, 2, 1},
		//		k:   1,
		//	},
		//	want: 4,
		//},
		//{
		//	args: args{
		//		arr: []int{4, 1, 5, 2, 6, 2},
		//		k:   2,
		//	},
		//	want: 0,
		//},
		//{
		//	args: args{
		//		arr: []int{4, 1, 5, 2, 6, 2},
		//		k:   3,
		//	},
		//	want: 2,
		//},
		{
			args: args{
				arr: []int{12, 6, 12, 6, 14, 2, 13, 17, 3, 8, 11, 7, 4, 11, 18, 8, 8, 3},
				k:   1,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kIncreasing(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("kIncreasing() = %v, want %v", got, tt.want)
			}
		})
	}
}
