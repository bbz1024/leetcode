package dynamic

/*

 */

func findTargetSumWays(nums []int, target int) int {
	/*
		left+right=sum
		right=sum-left
		left-right=target
		left-sum+left=target
		left=(target+sum)/2
	*/
	sum := 0
	for _, num := range nums {
		sum += num
	}
	left := (target + sum) / 2
	if abs(target) > sum {
		return 0
	}
	if (target+sum)%2 == 1 {
		return 0
	} // 此时没有方案

	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, left+1)
	}
	// dp[i][j]：表示nums[0:i]中，和为j的方案数
	// 初始化
	if nums[0] <= left {
		dp[0][nums[0]] = 1
	}

	if nums[0] == 0 {
		dp[0][0] = 2
	} else {
		dp[0][0] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < left+1; j++ {
			if j >= nums[i] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)-1][left] // 这里是通过left来进行记录最终装满的方案
}

func abs(target int) int {
	if target < 0 {
		return -target
	}
	return target
}
func findTargetSumWays2(nums []int, target int) int {
	/*
		left+right=sum
		right=sum-left
		left-right=target
		left-sum+left=target
		left=(target+sum)/2
	*/
	sum := 0
	for _, num := range nums {
		sum += num
	}
	left := (target + sum) / 2
	if abs(target) > sum {
		return 0
	}
	if (target+sum)%2 == 1 {
		return 0
	} // 此时没有方案

	dp := make([]int, left+1)
	if nums[0] <= left {
		dp[nums[0]] = 1
	}
	if nums[0] == 0 {
		dp[0] = 2
	} else {
		dp[0] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := left; j >= 0; j-- {
			if j >= nums[i] {
				dp[j] = dp[j] + dp[j-nums[i]]
			}
		}
	}
	return dp[left]
}
