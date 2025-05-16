package main

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var res = 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			res++
		} else {
			res = 1
		}
	}
	return res
}
