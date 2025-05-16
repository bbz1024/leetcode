package dynamic

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	size := sum / 2
	dp := make([][]int, len(stones))
	for i := 0; i < len(stones); i++ {
		dp[i] = make([]int, size+1)
	}
	// 初始化
	for i := stones[0]; i <= size; i++ {
		dp[0][i] = stones[0]
	}
	for i := 1; i < len(stones); i++ {
		for j := 1; j <= size; j++ {
			if stones[i] <= j {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-stones[i]]+stones[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return sum - 2*dp[len(stones)-1][size]
}
func lastStoneWeightII2(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	size := sum / 2
	dp := make([]int, size+1)
	for i := stones[0]; i <= size; i++ {
		dp[i] = stones[0]
	}
	for i := 1; i < len(stones); i++ {
		for j := size; j >= 0; j-- {
			if stones[i] <= j {
				dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
			} else {
				dp[j] = dp[j]
			}
		}
	}
	return sum - 2*dp[size]
}
