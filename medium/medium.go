package medium

import "math"

// 322. 零钱兑换
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 0; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

// 931. 下降路径最小和
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	copy(dp[0], matrix[0])

	// dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			minVal := dp[i-1][j]
			if j > 0 {
				minVal = min(minVal, dp[i-1][j-1])
			}
			if j+1 < n {
				minVal = min(minVal, dp[i-1][j+1])
			}

			dp[i][j] = minVal + matrix[i][j]
		}
	}

	minRes := dp[n-1][0]

	for i := 1; i < n; i++ {
		minRes = min(minRes, dp[n-1][i])
	}

	return minRes
}

// 72. 编辑距离
func minDistance(word1 string, word2 string) int {
	//m, n := len(word1), len(word2)
	//
	//var dp func(word1 string, m int, word2 string, n int) int
	//
	//dp = func(word1 string, m int, word2 string, n int) int {
	//	if m == -1 {
	//		return n + 1
	//	}
	//
	//	if n == -1 {
	//		return m + 1
	//	}
	//
	//	if word1[m] == word2[n] {
	//		return dp(word1, m-1, word2, n-1)
	//	}
	//
	//	return min(dp(word1, m-1, word2, n-1), dp(word1, m-1, word2, n), dp(word1, m, word2, n-1)) + 1
	//
	//}
	//
	//return dp(word1, m-1, word2, n-1)

	m, n := len(word1), len(word2)

	dp := make([][]int, m+1)

	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 0; i < n+1; i++ {
		dp[0][i] = i
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[m][n]
}
