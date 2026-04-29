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
