package main

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])

	var res int
	var bfs func(i, j int, curRes *int)
	bfs = func(i, j int, curRes *int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		*curRes++
		bfs(i-1, j, curRes)
		bfs(i+1, j, curRes)
		bfs(i, j-1, curRes)
		bfs(i, j+1, curRes)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			curRes := 0
			if grid[i][j] == 1 {
				bfs(i, j, &curRes)
			}
			if curRes > res {
				res = curRes
			}
		}
	}
	return res
}
