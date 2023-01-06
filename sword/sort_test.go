package sword

import (
	"reflect"
	"testing"
)

func Test_sortInts(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{4, 1, 5, 3, 6, 2, 7, 8},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortInts(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("sortInts() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
