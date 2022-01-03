package main

import "testing"

func Test_checkString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{s: "aaabbb"},
			want: true,
		},
		{
			args: args{s: "abab"},
			want: false,
		},
		{
			args: args{s: "bbb"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkString(tt.args.s); got != tt.want {
				t.Errorf("checkString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberOfBeams(t *testing.T) {
	type args struct {
		bank []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{bank: []string{"011001", "000000", "010100", "001000"}},
			want: 8,
		},
		{
			args: args{bank: []string{"000", "111", "000"}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfBeams(tt.args.bank); got != tt.want {
				t.Errorf("numberOfBeams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_asteroidsDestroyed(t *testing.T) {
	type args struct {
		mass      int
		asteroids []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				mass:      10,
				asteroids: []int{3, 9, 19, 5, 21},
			},
			want: true,
		},
		{
			args: args{
				mass:      5,
				asteroids: []int{4, 9, 23, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := asteroidsDestroyed(tt.args.mass, tt.args.asteroids); got != tt.want {
				t.Errorf("asteroidsDestroyed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumInvitations(t *testing.T) {
	type args struct {
		favorite []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{[]int{2, 2, 1, 2}},
			want: 3,
		},
		{
			args: args{[]int{1, 2, 0}},
			want: 3,
		},
		{
			args: args{[]int{3, 0, 1, 4, 1}},
			want: 4,
		},
		{
			args: args{[]int{1, 0, 0, 2, 1, 4, 7, 8, 9, 6, 7, 10, 8}},
			want: 6,
		},
		{
			args: args{[]int{1, 0, 3, 2, 5, 6, 7, 4, 9, 8, 11, 10, 11, 12, 10}},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumInvitations(tt.args.favorite); got != tt.want {
				t.Errorf("maximumInvitations() = %v, want %v", got, tt.want)
			}
		})
	}
}
