package main

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var res int
	left := 0
	sum := 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		if sum >= target {
			for left <= right {
				if res == 0 || right-left+1 < res {
					res = right - left + 1
				}
				sum -= nums[left]
				left++
				if sum < target {
					break
				}
			}

		}
	}
	return res
}
