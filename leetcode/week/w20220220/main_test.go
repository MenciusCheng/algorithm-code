package w20220220

import (
	"reflect"
	"testing"
)

func Test_countEven(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				num: 4,
			},
			want: 2,
		},
		{
			args: args{
				num: 30,
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countEven(tt.args.num); got != tt.want {
				t.Errorf("countEven() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeNodes(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				head: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val: 5,
										Next: &ListNode{
											Val: 2,
											Next: &ListNode{
												Val:  0,
												Next: nil,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repeatLimitedString(t *testing.T) {
	type args struct {
		s           string
		repeatLimit int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s:           "cczazcc",
				repeatLimit: 3,
			},
			want: "zzcccac",
		},
		{
			args: args{
				s:           "aababab",
				repeatLimit: 2,
			},
			want: "bbabaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatLimitedString(tt.args.s, tt.args.repeatLimit); got != tt.want {
				t.Errorf("repeatLimitedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
