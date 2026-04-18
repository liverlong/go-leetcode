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
