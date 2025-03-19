package backtracking

import "strings"

func solveNQueens(n int) [][]string {
	var res [][]string
	mp := make([][]string, n)
	for i := 0; i < n; i++ {
		mp[i] = make([]string, n)
		for j := 0; j < n; j++ {
			mp[i][j] = "."
		}
	}
	var backtrack func(level int, path []string)
	var strBuild strings.Builder
	backtrack = func(level int, path []string) {

		if level == n {
			tmp := make([]string, n)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for y := 0; y < n; y++ {
			// 上米线
			if checkAround(mp, level, y) {
				continue
			}
			mp[level][y] = "Q"
			// 收集path

			for i := 0; i < n; i++ {
				strBuild.WriteString(mp[level][i])
			}
			path = append(path, strBuild.String())
			strBuild.Reset()
			backtrack(level+1, path)
			mp[level][y] = "."
			path = path[:len(path)-1]
		}
	}

	backtrack(0, []string{})
	return res
}
func checkAround(mp [][]string, x, y int) bool {
	// 同一列，同斜线
	for i := 0; i < x; i++ {
		if mp[i][y] == "Q" {
			return true
		}
	}
	left := x - 1
	up := y - 1
	for left >= 0 && up >= 0 {
		if mp[left][up] == "Q" {
			return true
		}
		left--
		up--
	}
	right := y + 1
	up = x - 1
	for right < len(mp) && up >= 0 {
		if mp[up][right] == "Q" {
			return true
		}
		right++
		up--
	}
	return false
}
