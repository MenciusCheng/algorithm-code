package w20230101

import "testing"

func Test_minimumPartition(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s: "1829727651",
				k: 10,
			},
			want: 10,
		},
		{
			args: args{
				s: "165462",
				k: 60,
			},
			want: 4,
		},
		{
			args: args{
				s: "238182",
				k: 5,
			},
			want: -1,
		},
		{
			args: args{
				s: "1",
				k: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumPartition(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("minimumPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
