package greedy

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	var res int
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] < intervals[j][1] {
			return true
		} else if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] < intervals[j][0]
		}
		return false
	})
	for i := 0; i < len(intervals)-1; i++ {
		//
		if intervals[i][1] > intervals[i+1][0] {
			res++
			intervals[i+1][0] = intervals[i][0]
			intervals[i+1][1] = intervals[i][1]
		}
	}
	return res
}

/*
[[1,2],[2,3],[3,4],[1,3]]

[[1,2],[1,3],[2,3],[3,4]]


*/
