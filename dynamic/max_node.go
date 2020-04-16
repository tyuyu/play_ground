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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int

func maxPathSum(root *TreeNode) (sum int) {
	if root == nil {
		return 0
	}
	res = math.MinInt64 //用来判断 所有节点都是负数时，应返回哪一个。
	dfs(root)
	return res
}

func dfs(root *TreeNode) (sum int) {
	if root == nil {
		return
	}
	lMax := max(0, dfs(root.Left))     //左分支 最大贡献
	rMax := max(0, dfs(root.Right))    //右分支 最大贡献
	res = max(res, root.Val+lMax+rMax) //与当前结果对比
	return root.Val + max(lMax, rMax)  //返回单分支最大值
}

func max(nums ...int) int {
	sort.Ints(nums)
	return nums[len(nums)-1]
}
