package easy

// 反转二叉树
func invertTree(root *TreeNode) *TreeNode {
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

// 1 两数之和
func twoSum(nums []int, target int) []int {
	memo := make(map[int]int)
	r := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		if res, ok := memo[target-nums[i]]; ok {
			r[0], r[1] = res, i
			break
		} else {
			memo[nums[i]] = i
		}
	}

	return r
}

// 94. 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)

	}

	dfs(root)

	return res
}

// 136. 只出现一次的数字
func singleNumber(nums []int) int {
	res := 0

	for _, num := range nums {
		res ^= num
	}

	return res
}

// 104. 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1

}

// 70. 爬楼梯
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	//
	//return climbStairs(n-1) + climbStairs(n-2)

	first, second := 0, 1

	res := 0
	for i := 1; i < n+1; i++ {
		res = first + second
		first, second = second, res
	}

	return res
}

// 617. 合并二叉树
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil || root2 == nil {
		if root1 == nil {
			return root2
		}
		return root1
	}
	root1.Val += root2.Val
	left := mergeTrees(root1.Left, root2.Left)
	right := mergeTrees(root1.Right, root2.Right)
	root1.Left, root1.Right = left, right

	return root1
}

// 20 有效的括号
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0)

	for _, val := range s {
		if val == '{' || val == '[' || val == '(' {
			stack = append(stack, val)
		} else {
			if len(stack) == 0 {
				return false
			}

			c := stack[len(stack)-1]

			if val == '}' && c != '{' {
				return false
			}

			if val == ']' && c != '[' {
				return false
			}

			if val == ')' && c != '(' {
				return false
			}

			stack = stack[:len(stack)-1]

		}
	}

	return len(stack) == 0
}

// 206. 反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	var pre *ListNode

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

// 338. 比特位计数
func countBits(n int) []int {
	res := make([]int, n+1)
	res[0] = 0

	for i := 0; i <= n; i++ {
		if i%2 == 0 { // 偶数
			res[i] = res[i%2]
		} else {
			res[i] = res[i-1] + 1
		}
	}

	return res
}

// 101. 对称二叉树（深度优先搜索，清晰图解）
func isSymmetric(root *TreeNode) bool {
	var recur func(left, right *TreeNode) bool

	recur = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}

		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		return recur(left.Left, right.Right) && recur(left.Right, right.Left)
	}

	return root == nil || recur(root.Left, root.Right)
}

// 121. 买卖股票的最佳时机
func maxProfit(prices []int) int {
	//n := len(prices)
	//dp := make([][]int, n)
	//for i := 0; i < n; i++ {
	//	dp[i] = make([]int, 2)
	//}
	//for i := 0; i < n; i++ {
	//	if i == 0 {
	//		dp[0][0] = 0
	//		dp[0][1] = -prices[i]
	//	} else {
	//		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
	//		dp[i][1] = max(dp[i-1][1], -prices[i]) // 只买一次
	//	}
	//}
	//
	//return dp[n-1][0]

	if len(prices) < 1 {
		return 0
	}

	dp_i_0, dp_i_1 := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, -prices[i])
	}

	return dp_i_0
}

// 448. 找到所有数组中消失的数字
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		nums[(nums[i]-1)%n] += n
	}

	res := make([]int, 0)
	for i := 0; i < n; i++ {
		if nums[i] <= n {
			res = append(res, i+1)
		}
	}

	return res
}

// 283. 移动零
func moveZeroes(nums []int) {
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[slow] = nums[i]
			slow++
		}
	}
	for ; slow < len(nums); slow++ {
		nums[slow] = 0
	}
}

// 169. 多数元素
func majorityElement(nums []int) int {
	candidate, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
