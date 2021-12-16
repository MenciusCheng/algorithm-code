package exmath

import "testing"

func TestTanExample(t *testing.T) {
	type args struct {
		degree float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{degree: 45},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TanExample(tt.args.degree)
		})
	}
}
