package backtracking

import "testing"

func TestSolveNQueens(t *testing.T) {
	for i, strings := range solveNQueens(4) {
		t.Logf("第%d组结果：%v", i, strings)
	}
}
func TestCheckAround(t *testing.T) {
	c := [][]string{
		{".", "Q", ".", "."},
		{".", ".", ".", "."},
		{".", ".", ".", "."},
		{".", ".", ".", "."},
	}
	t.Log(checkAround(c, 1, 1))
}
