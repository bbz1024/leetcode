package main

func generateMatrix(n int) [][]int {
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}
	c := n / 2
	// 奇数
	if n%2 == 1 {
		c++
	}
	right := n
	end := n
	cnt := 1
	for i := 0; i < c; i++ {
		for j := i; j < right-1; j++ {
			arr[i][j] = cnt
			cnt++
		}
		for j := i; j < end-1; j++ {
			arr[j][right-1] = cnt
			cnt++
		}
		for j := right - 1; j > i; j-- {
			arr[end-1][j] = cnt
			cnt++
		}
		for j := end - 1; j > i; j-- {
			arr[j][i] = cnt
			cnt++
		}
		right--
		end--
	}
	if n%2 == 1 {
		arr[n/2][n/2] = cnt
	}
	return arr
}
