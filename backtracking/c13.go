package backtracking

func solveSudoku(board [][]byte) {
	if len(board) == 0 {
		return
	}
	var backtrack func() bool
	backtrack = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] != '.' {
					continue
				}
				for val := '1'; val <= '9'; val++ {
					if isvalid(i, j, byte(val), board) {
						board[i][j] = byte(val)
						if backtrack() {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
		return true
	}
	backtrack()
}

// 判断填入数字是否满足要求
func isvalid(row, col int, k byte, board [][]byte) bool {
	for i := 0; i < 9; i++ { //行
		if board[row][i] == k {
			return false
		}
	}
	for i := 0; i < 9; i++ { //列
		if board[i][col] == k {
			return false
		}
	}
	//方格
	startrow := (row / 3) * 3
	startcol := (col / 3) * 3
	for i := startrow; i < startrow+3; i++ {
		for j := startcol; j < startcol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}
