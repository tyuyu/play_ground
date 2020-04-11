package tree

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//3,5,1,6,2,0,8,null,null,7,4
func ArrayToTree(a string) *TreeNode {

	strs := strings.Split(a, ",")

	if len(strs) == 0 {
		return nil
	}

	tree := make([]*TreeNode, len(strs))
	for i, str := range strs {
		if str == "null" {
			tree[i] = nil
		} else {
			val, _ := strconv.Atoi(str)
			tree[i] = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
		}
		if i == 0 {
			continue
		}
		p := (i - 1) / 2
		if i%2 == 1 {
			tree[p].Left = tree[i]
		} else {
			tree[p].Right = tree[i]
		}
	}

	return tree[0]
}
