package w20220410

import "testing"

func Test_minimizeResult(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{expression: "12+34"},
			want: "1(2+3)4",
		},
		{
			args: args{expression: "1+1"},
			want: "(1+1)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimizeResult(tt.args.expression); got != tt.want {
				t.Errorf("minimizeResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumProduct(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{0, 4},
				k:    5,
			},
			want: 20,
		},
		{
			args: args{
				nums: []int{6, 3, 3, 2},
				k:    2,
			},
			want: 216,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProduct(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
