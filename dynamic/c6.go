package dynamic

func integerBreak(n int) int {
	if n <= 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j <= i; j++ {
			dp[i] = max(dp[i-j]*j, (i-j)*j, dp[i])
		}
	}
	return dp[n]
}
