package codetop

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	var res []int
	left := 0
	right := n - 1
	up := 0
	down := m - 1
	for left <= right && up <= down {
		// 左到右不进行处理边界
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		// 从上到下
		for i := up + 1; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		if left < right && up < down {
			// 右到左不进行处理边界
			for i := right - 1; i > left; i-- {
				res = append(res, matrix[down][i])
			}
			for i := down; i > up; i-- {
				res = append(res, matrix[i][left])
			}
		}
		left++
		right--
		up++
		down--
	}
	return res
}
