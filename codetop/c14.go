package codetop

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	// 深度优先搜索
	m := len(grid)
	n := len(grid[0])

	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				res++

			}
		}
	}
	return res
}

func dfs(grid [][]byte, i, j int) {
	m := len(grid)
	n := len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
	return
}
