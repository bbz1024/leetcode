package greedy

import "testing"

func TestFindMinArrowShots(t *testing.T) {
	t.Log(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	//[[1,2],[3,4],[5,6],[7,8]]
	t.Log(findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
}
