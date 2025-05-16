package codetop

import "sort"

func merge2(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	l := len(intervals)
	var end, start int

	for i := 1; i < l; i++ {
		if intervals[i][0] <= intervals[start][1] {
			intervals[start][1] = max(intervals[i][1], intervals[start][1])
		} else {
			intervals[end+1] = intervals[i]
			start = i
			end++
		}
	}
	return intervals[:end+1]
}
