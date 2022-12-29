package sword

import (
	"reflect"
	"testing"
)

func Test_dfsInOrder(t *testing.T) {
	type args struct {
		root *TreeNode
		f    func(v int)
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				root: treeRoot1,
			},
			want: []int{4, 2, 5, 1, 6, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, 0)
			dfsInOrder(tt.args.root, func(v int) {
				got = append(got, v)
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dfsInOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dfsInOrderFor(t *testing.T) {
	type args struct {
		root *TreeNode
		f    func(v int)
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				root: treeRoot1,
			},
			want: []int{4, 2, 5, 1, 6, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, 0)
			dfsInOrderFor(tt.args.root, func(v int) {
				got = append(got, v)
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dfsInOrderFor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dfsPreOrder(t *testing.T) {
	type args struct {
		root *TreeNode
		f    func(v int)
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				root: treeRoot1,
			},
			want: []int{1, 2, 4, 5, 3, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, 0)
			dfsPreOrder(tt.args.root, func(v int) {
				got = append(got, v)
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dfsPreOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dfsPreOrderFor(t *testing.T) {
	type args struct {
		root *TreeNode
		f    func(v int)
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				root: treeRoot1,
			},
			want: []int{1, 2, 4, 5, 3, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, 0)
			dfsPreOrderFor(tt.args.root, func(v int) {
				got = append(got, v)
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dfsPreOrderFor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dfsPostOrder(t *testing.T) {
	type args struct {
		root *TreeNode
		f    func(v int)
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				root: treeRoot1,
			},
			want: []int{4, 5, 2, 6, 7, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, 0)
			dfsPostOrder(tt.args.root, func(v int) {
				got = append(got, v)
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dfsPostOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dfsPostOrderFor(t *testing.T) {
	type args struct {
		root *TreeNode
		f    func(v int)
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				root: treeRoot1,
			},
			want: []int{4, 5, 2, 6, 7, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, 0)
			dfsPostOrderFor(tt.args.root, func(v int) {
				got = append(got, v)
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dfsPostOrderFor() = %v, want %v", got, tt.want)
			}
		})
	}
}
