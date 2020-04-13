package tree

import (
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "case1",
			args: args{
				root: ArrayToTree("3,5,1,6,2,0,8,null,null,7,4"),
				p: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				q: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
			want: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		}, {
			name: "case2",
			args: args{
				root: ArrayToTree("3,5,1,6,2,0,8,null,null,7,4"),
				p: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				q: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			want: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); tt.want.Val != got.Val {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
