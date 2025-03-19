package backtracking

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)
	var backtrack func(nums []int, start int)
	backtrack = func(nums []int, start int) {
		if sumArr(nums) > target {
			return
		}
		if sumArr(nums) == target {
			arr := make([]int, len(nums))
			copy(arr, nums)
			res = append(res, arr)
			return
		}
		for i := start; i < len(candidates); i++ {
			nums = append(nums, candidates[i])
			backtrack(nums, i)
			nums = nums[:len(nums)-1]
		}
	}
	backtrack([]int{}, 0)
	return res
}
