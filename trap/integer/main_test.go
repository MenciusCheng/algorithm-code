package main

import "testing"

func TestPrefixZero(t *testing.T) {
	type args struct {
		d int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				d: 0,
			},
			want: "00",
		},
		{
			args: args{
				d: 1,
			},
			want: "01",
		},
		{
			args: args{
				d: 12,
			},
			want: "12",
		},
		{
			args: args{
				d: 123,
			},
			want: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrefixZero(tt.args.d); got != tt.want {
				t.Errorf("PrefixZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
