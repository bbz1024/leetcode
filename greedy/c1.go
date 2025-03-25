package greedy

import "sort"

func findContentChildren(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}

	sort.Ints(g)
	sort.Ints(s)
	var res int
	child := 0
	candidate := 0
	for child < len(g) && candidate < len(s) {
		// 可以满足
		if g[child] <= s[candidate] {
			res++
			child++
		}
		candidate++
	}
	return res
}
