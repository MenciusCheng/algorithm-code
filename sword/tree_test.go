package sword

import (
	"reflect"
	"testing"
)

var treeRoot1 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	},
	Right: &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	},
}

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{root: treeRoot1},
			want: []int{4, 2, 5, 1, 6, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inorderTraversalFor(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{root: treeRoot1},
			want: []int{4, 2, 5, 1, 6, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversalFor(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversalFor() = %v, want %v", got, tt.want)
			}
		})
	}
}
