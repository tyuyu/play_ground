package tree

import "math"

//给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
//
//示例:
//
//输入: [1,2,3,null,5,null,4]
//输出: [1, 3, 4]
//解释:
//
//   1            <---
// /   \
//2     3         <---
// \     \
//  5     4       <---
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-right-side-view
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func rightSideView(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	level := make([][]int, 0)
	deph(root, &level, 0)
	res := make([]int, 0)
	for _, l := range level {
		res = append(res, l[0])
	}
	return res
}

func deph(root *TreeNode, levels *[][]int, deep int) {
	if len(*levels) < deep+1 {
		*levels = append(*levels, []int{})
	}
	if root.Right != nil {
		deph(root.Right, levels, deep+1)
	}
	if root.Left != nil {
		deph(root.Left, levels, deep+1)
	}
	(*levels)[deep] = append((*levels)[deep], root.Val)
}

//执行结果：
//通过
//显示详情
//执行用时 :
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗 :
//2.3 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户

func reverse(x int) int {

	res := 0
	for x != 0 {
		c := x % 10
		res = res*10 + c
		x = x / 10
	}
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return res
}

//执行结果：
//通过
//显示详情
//执行用时 :
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗 :
//2.2 MB
//, 在所有 Go 提交中击败了
//42.86%
//的用户
