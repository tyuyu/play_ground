package tree

var pre *TreeNode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Right)
	flatten(root.Left)
	root.Right = pre
	root.Left = nil
	pre = root
	return
}
