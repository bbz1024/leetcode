package codetop

// 前缀和
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var res int
	res = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+nums[i] > nums[i] {
			nums[i] = nums[i-1] + nums[i]
		}
		res = max(res, nums[i])
	}
	return res
}
