package random

import (
	"math"
	"reflect"
	"testing"
)

func TestWeightedRandomIdx(t *testing.T) {
	type args struct {
		weights []int
		n       int
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "10% 20% 30% 40%",
			args: args{
				weights: []int{1, 2, 3, 4},
				n:       100000,
			},
			want: map[int]int{0: 10, 1: 20, 2: 30, 3: 40},
		},
		{
			name: "80% 20%",
			args: args{
				weights: []int{8, 2},
				n:       100000,
			},
			want: map[int]int{0: 80, 1: 20},
		},
		{
			name: "100%",
			args: args{
				weights: []int{1},
				n:       100000,
			},
			want: map[int]int{0: 100},
		},
		{
			name: "empty",
			args: args{
				weights: []int{},
				n:       100000,
			},
			want: map[int]int{-1: 100},
		},
		{
			name: "5% 0% 95%",
			args: args{
				weights: []int{5, 0, 95},
				n:       100000,
			},
			want: map[int]int{0: 5, 2: 95},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			countMap := make(map[int]int)
			for i := 0; i < tt.args.n; i++ {
				gotIdx := WeightedRandomIdx(tt.args.weights)
				countMap[gotIdx]++
			}
			res := make(map[int]int)
			for k, count := range countMap {
				percentage := int(math.Round(float64(count) / float64(tt.args.n) * 100))
				res[k] = percentage
			}
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("WeightedRandomIdx(), res = %v, want = %v, countMap = %v, run %dtimes", res, tt.want, countMap, tt.args.n)
			} else {
				t.Logf("WeightedRandomIdx(), res = %v, countMap = %v, run %d times", res, countMap, tt.args.n)
			}
		})
	}
}
