package dynamic

import (
	"math"
	"sort"
)

//给定一个非空二叉树，返回其最大路径和。
//
//本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
//
//示例 1:
//
//输入: [1,2,3]
//
//       1
//      / \
//     2   3
//
//输出: 6
//示例 2:
//
//输入: [-10,9,20,null,null,15,7]
//
//   -10
//   / \
//  9  20
//    /  \
//   15   7
//
//输出: 42

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := math.MinInt64
	dsf(root, &max)
	return max
}

func dsf(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	l, r := dsf(root.Left, max), dsf(root.Right, max)
	var m int
	if root.Val < 0 {
		if root.Left != nil && root.Right != nil {
			m = int(math.Max(float64(l), float64(r)))
		} else if root.Left == nil && root.Right == nil {
			m = root.Val
		} else if root.Left == nil {
			m = int(math.Max(float64(root.Val), float64(r)))
		} else {
			m = int(math.Max(float64(l), float64(root.Val)))
		}
	} else {
		res := []int{root.Val, root.Val + l + r, root.Val + r, root.Val + l}
		sort.Ints(res)
		m = res[len(res)-1]
	}
	if m > *max {
		*max = m
	}
	return m
}
