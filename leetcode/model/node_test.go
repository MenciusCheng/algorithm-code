package model

import (
	"reflect"
	"testing"
)

func TestArrToListNode(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				arr: []int{1, 2, 3},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrToListNode(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrToListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNodeToArr(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListNodeToArr(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListNodeToArr() = %v, want %v", got, tt.want)
			}
		})
	}
}
