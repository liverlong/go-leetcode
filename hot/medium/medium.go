package medium

import "math"

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

// 98 验证二叉搜索树
func isValidBST(root *TreeNode) bool {

	// 中序遍历

	pre := math.MinInt

	var dfs func(root *TreeNode) bool

	dfs = func(root *TreeNode) bool {
		if root == nil {
			return true
		}

		if !dfs(root.Left) {
			return false
		}

		if root.Val < pre {
			return false
		}

		pre = root.Val

		return dfs(root.Right)
	}

	return dfs(root)
}

// 46 全排列
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := make([][]int, 0)

	// 状态变量
	used := make([]bool, len(nums))
	tmp := make([]int, 0, len(nums))

	var backtrace func()

	backtrace = func() {
		if len(nums) == len(tmp) {
			t := make([]int, len(tmp))
			copy(t, tmp)
			res = append(res, t)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			// 做选择
			tmp = append(tmp, nums[i])
			used[i] = true

			// 回溯
			backtrace()

			// 撤销选择
			tmp = tmp[:len(tmp)-1]
			used[i] = false
		}
	}

	backtrace()

	return res
}

// 22 括号生存
func generateParenthesis(n int) []string {
	var res []string
	if n <= 0 {
		return res
	}

	track := make([]byte, 0, 2*n)

	var dfs func(left, right int)

	dfs = func(left, right int) {
		// 终止条件
		if left == 0 && right == 0 {
			res = append(res, string(track))
			return
		}

		// 剪枝
		if left > right {
			return
		}

		// 尝试添加左括号
		if left > 0 {
			track = append(track, '(')
			dfs(left-1, right)
			track = track[:len(track)-1]
		}

		// 尝试添加右括号
		if right > 0 {
			track = append(track, ')')
			dfs(left, right-1)
			track = track[:len(track)-1]
		}

	}

	dfs(n, n)
	return res
}

// 78 子集
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	var tmp []int

	var backtrace func(start int)

	backtrace = func(start int) {
		// 当前路径加入到结果集
		t := make([]int, len(tmp))
		copy(t, tmp)
		res = append(res, t)

		// 从start遍历，避免回头选重复的数据
		for i := start; i < len(nums); i++ {

			// 做选择
			tmp = append(tmp, nums[i])
			// 下一层不选当前的数
			backtrace(i + 1)
			// 撤销选择（回溯）
			tmp = tmp[:len(tmp)-1]
		}
	}

	backtrace(0)

	return res
}

// 48 旋转图像
func rotate(matrix [][]int) {
	n := len(matrix)

	// 1. 转置矩阵
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// 交换对角线两侧的元素
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 2. 镜像翻转每一行
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			// 交换行内的首尾元素
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}
