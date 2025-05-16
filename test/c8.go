package test

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	pd := make([][]int, m)
	for i := range pd {
		pd[i] = make([]int, n)
	}
	// 初始化上边界和左边界
	pd[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		pd[i][0] = pd[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		pd[0][i] = pd[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			pd[i][j] = min(pd[i-1][j], pd[i][j-1]) + grid[i][j]
		}
	}
	return pd[m-1][n-1]
}
