package dynamic

func findMaxForm(strs []string, m int, n int) int {
	// 尽可能的装满两个最大的背包，
	// m，n 表示两个背包的容量，strs 表示物品列表
	// strs[i]: 物品的重量
	if len(strs) == 0 {
		return 0
	}
	// dp[i][m][n]int 表示第i个物品，第m个背包，第n个背包，装了多少个物品
	dp := make([][][]int, len(strs))
	for i := 0; i < len(strs); i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
			for k := 0; k <= n; k++ {
				dp[i][j][k] = 0
			}
		}
	}
	str := strs[0]
	zeroCnt := count0(str)
	for i := 0; i <= m; i++ {
		if zeroCnt > i {
			continue
		}
		for j := 0; j <= n; j++ {
			if len(str)-zeroCnt > j {
				continue
			}
			dp[0][i][j] = 1
		}
	}
	for i := 1; i < len(strs); i++ {
		cnt0 := count0(strs[i])
		cnt1 := len(strs[i]) - cnt0
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				// 容量不足
				if j-cnt0 < 0 || k-cnt1 < 0 {
					dp[i][j][k] = dp[i-1][j][k]
				} else {
					// 尝试去装 ，判断装了之后+之前剩余计算的最大数
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-cnt0][k-cnt1]+1)
				}
			}
		}
	}
	return dp[len(strs)-1][m][n]
}
func findMaxForm2(strs []string, m int, n int) int {

	if len(strs) == 0 {
		return 0
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// 初始化
	str := strs[0]
	cnt0 := count0(str)
	for i := 0; i <= m; i++ {
		if cnt0 > i {
			continue
		}
		for j := 0; j <= n; j++ {
			if len(str)-cnt0 > j {
				continue
			}
			dp[i][j] = 1
		}
	}

	for i := 1; i < len(strs); i++ {
		cnt0 := count0(strs[i])
		cnt1 := len(strs[i]) - cnt0
		for j := m; j >= cnt0; j-- {
			for k := n; k >= cnt1; k-- {
				// 容量不足
				if j-cnt0 < 0 || k-cnt1 < 0 {
					continue
				}
				// 尝试去装 ，判断装了之后+之前剩余计算的最大数
				dp[j][k] = max(dp[j][k], dp[j-cnt0][k-cnt1]+1)
			}
		}
	}
	return dp[m][n]
}
func count0(str string) (count0 int) {
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			count0++
		}
	}
	return
}
