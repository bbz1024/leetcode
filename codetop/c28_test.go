package codetop

import "testing"

func TestMerge2(t *testing.T) {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	res := merge2(intervals)
	t.Log(res)
}
