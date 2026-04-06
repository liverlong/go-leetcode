package medium

// 除了自身以外数组的乘积
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	tmp := 1

	for i := len(res) - 1; i >= 0; i-- {
		// i处左边乘积位res[i] 右边乘积为 tmp
		res[i] = res[i] * tmp

		// 计算下一个的时候更新tmp的值
		tmp *= nums[i]
	}
	return res
}

// 102 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int(nil)
	}

	q := make([]*TreeNode, 0)

	q = append(q, root)

	res := make([][]int, 0)

	for len(q) != 0 {
		sz := len(q)
		tmp := make([]int, 0)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			tmp = append(tmp, cur.Val)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		res = append(res, tmp)
	}

	return res

}
