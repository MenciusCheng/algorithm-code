package offer_special

import "testing"

func Test_divide(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				a: -2147483648,
				b: -1,
			},
			want: 2147483648,
		},
		{
			args: args{
				a: -2147483648,
				b: 1,
			},
			want: -2147483648,
		},
		{
			args: args{
				a: 15,
				b: 2,
			},
			want: 7,
		},
		{
			args: args{
				a: 7,
				b: -3,
			},
			want: -2,
		},
		{
			args: args{
				a: 0,
				b: 1,
			},
			want: 0,
		},
		{
			args: args{
				a: 1,
				b: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
