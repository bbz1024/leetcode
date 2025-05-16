package test

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var res int
	for i := 0; i < len(nums); i++ {
		// 遇到重复的数字，跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		nums[res] = nums[i]
		res++
	}
	return res
}
