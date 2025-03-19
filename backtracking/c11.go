package backtracking

import "sort"

func permuteUnique(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}
	var backtrack func(path []int)
	sort.Ints(nums)
	used := make([]bool, len(nums))
	backtrack = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
				continue
			}
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack(path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack([]int{})
	return res
}
