package exmath

import "testing"

func TestIncome(t *testing.T) {
	type args struct {
		startAmount   int
		monthlyAmount int
		year          int
		yearlyPer     float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				startAmount:   0,
				monthlyAmount: 5000,
				year:          20,
				yearlyPer:     3.0,
			},
		},
		{
			args: args{
				startAmount:   50000,
				monthlyAmount: 0,
				year:          20,
				yearlyPer:     3.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Income(tt.args.startAmount, tt.args.monthlyAmount, tt.args.year, tt.args.yearlyPer)
		})
	}
}
