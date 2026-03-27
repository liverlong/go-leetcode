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

// 汉明距离
func hammingDistance(x int, y int) int {
	n := x ^ y
	count := 0
	for n != 0 {
		n &= n - 1 // clear lowest set bit
		count++
	}
	return count
}
