package test

import "sort"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	var res int
	end := len(nums) - 1
	start := 0
	for start <= end {
		if nums[start] == val {
			nums[start], nums[end] = nums[end], nums[start]
			end--
		} else {
			start++
			res++
		}

	}
	return res
}
