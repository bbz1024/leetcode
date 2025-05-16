package dynamic

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	pd := make([][]int, m)
	for i := range pd {
		pd[i] = make([]int, n)
	}
	// 初始化pd
	// 左边和上边，如果有障碍物，则pd[i][j] = 0
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		pd[i][0] = 1
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		pd[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			pd[i][j] = pd[i-1][j] + pd[i][j-1]
		}
	}
	return pd[m-1][n-1]
}
