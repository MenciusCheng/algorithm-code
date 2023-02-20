package w20230219

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{n: 74415},
			want: 3,
		},
		//{
		//	args: args{n: 39},
		//	want: 3,
		//},
		//{
		//	args: args{n: 54},
		//	want: 3,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.n); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
