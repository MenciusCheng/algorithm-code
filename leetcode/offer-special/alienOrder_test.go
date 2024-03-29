package offer_special

import "testing"

func Test_alienOrder(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				words: []string{"abc", "ab"},
			},
			want: "",
		},
		{
			args: args{
				words: []string{"wrt", "wrf", "er", "ett", "rftt"},
			},
			want: "wertf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alienOrder(tt.args.words); got != tt.want {
				t.Errorf("alienOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
