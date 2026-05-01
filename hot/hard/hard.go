package hard

// 10. 正则表达式匹配
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	// 初始化空字符串匹配 a*, a*b*, a*b*c* 等模式
	for j := 2; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				// * 匹配零个前面的元素
				dp[i][j] = dp[i][j-2]
				// * 匹配一个或多个前面的元素
				if j >= 2 && (s[i-1] == p[j-2] || p[j-2] == '.') {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}

// 85. 最大矩形
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	heights := make([]int, cols)
	maxArea := 0

	for i := 0; i < rows; i++ {
		// 更新当前行的高度（柱状图）
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}

		// 单调栈计算当前柱状图的最大矩形面积
		stack := make([]int, 0, cols)
		for j := 0; j <= cols; j++ {
			var h int
			if j < cols {
				h = heights[j]
			}
			// 当遇到更矮的柱子时，以栈顶高度为高的矩形结束
			for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
				height := heights[stack[len(stack)-1]]
				stack = stack[:len(stack)-1]
				left := -1
				if len(stack) > 0 {
					left = stack[len(stack)-1]
				}
				width := j - left - 1
				area := height * width
				if area > maxArea {
					maxArea = area
				}
			}
			stack = append(stack, j)
		}
	}

	return maxArea
}
