package greedy

import "testing"

func TestMerge(t *testing.T) {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	res := merge(intervals)
	t.Log(res)
}
