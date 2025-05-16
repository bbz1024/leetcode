package dynamic

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	size := sum / 2
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, size+1)
	}
	for i := nums[0]; i <= size; i++ {
		if nums[0] <= i {
			dp[0][i] = nums[0]
		}
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j <= size; j++ {
			if nums[i] <= j {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(nums)-1][size] == size
}
func canPartition2(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	size := sum / 2
	dp := make([]int, size+1)
	// 初始化
	for i := 0; i <= size; i++ {
		if nums[0] <= i {
			dp[i] = nums[0]
		}
	}
	for i := 1; i < len(nums); i++ {
		for j := size; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[size] == size
}
