package backtracking

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var sum int
	used := make([]bool, len(candidates))
	var res [][]int
	var backtrack func(nums []int, start int)
	backtrack = func(nums []int, start int) {
		if sum > target {
			return
		}
		if sum == target {

			arr := make([]int, len(nums))
			copy(arr, nums)
			res = append(res, arr)
			return
		}
		for i := start; i < len(candidates); i++ {
			if i > 0 && candidates[i] == candidates[i-1] && !used[i-1] {
				continue
			}
			sum += candidates[i]
			nums = append(nums, candidates[i])
			used[i] = true
			backtrack(nums, i+1)
			nums = nums[:len(nums)-1]
			used[i] = false
			sum -= candidates[i]
		}

	}
	backtrack([]int{}, 0)
	return res
}
