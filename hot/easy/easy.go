package easy

import "go-lc/common"

// 反转二叉树
func invertTree(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return nil
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)

	root.Left = right
	root.Right = left
	return root
}
