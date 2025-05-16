package test

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var sum int
	var res int
	left := 0
	for right := 0; right < len(nums); {
		for right < len(nums) && sum < target {
			sum += nums[right]
			right++
		}
		if sum >= target {
			for left < right && sum >= target {
				sum -= nums[left]
				left++
			}
			if res == 0 || right-left < res {
				res = right - left
			}
		}
	}
	return res
}
