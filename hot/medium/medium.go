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
	res := make([][]int, 0)

	used := make([]bool, len(nums))

	var backtrace func(nums []int, used []bool, tmp []int, res *[][]int)

	backtrace = func(nums []int, used []bool, tmp []int, res *[][]int) {
		if len(nums) == len(tmp) {
			t := make([]int, len(tmp))
			copy(t, tmp)
			*res = append(*res, t)
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
			backtrace(nums, used, tmp, res)

			// 撤销选择
			tmp = tmp[:len(tmp)-1]
			used[i] = false
		}
	}

	backtrace(nums, used, []int{}, &res)

	return res
}

// 22 括号生存
func generateParenthesis(n int) []string {
	res := new([]string)
	if n <= 0 {
		return *res
	}

	var dfs func(left, right int, track string, res *[]string)

	dfs = func(left, right int, track string, res *[]string) {
		if left > right {
			return
		}

		if left < 0 || right < 0 {
			return
		}

		if left == 0 && right == 0 {
			*res = append(*res, track)
			return
		}

		// 加入左括号
		track += "("
		dfs(left-1, right, track, res)
		track = track[:len(track)-1]

		// 加入右括号
		track += ")"
		dfs(left, right-1, track, res)
		track = track[:len(track)-1]
	}

	dfs(n, n, "", res)
	return *res
}
