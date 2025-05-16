package test

func removeDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	slow := 2
	fast := 2
	for fast < len(nums) {

		if nums[slow-2] == nums[fast] {
			fast++
		} else {
			nums[slow] = nums[fast]
			slow++
			fast++
		}
	}
	return slow
}
