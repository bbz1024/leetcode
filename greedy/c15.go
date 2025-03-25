package greedy

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return false
	})

	var idx = 0
	start := intervals[idx][1]
	for i := 0; i < len(intervals); {
		temp := i
		for temp < len(intervals) && start >= intervals[temp][0] {
			start = max(start, intervals[temp][1])
			intervals[idx][1] = start
			temp++
		}
		if temp == i {
			intervals[idx] = intervals[i]
			i++
			idx++
			continue
		}
		i = temp
		idx++
	}
	return intervals[:idx]
}
