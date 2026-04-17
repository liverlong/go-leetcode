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
