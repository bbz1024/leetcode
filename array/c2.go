package main

// [0,1,2,2,3,0,4,2]
// 2
func removeElement(nums []int, val int) int {
	fast := 0
	slow := 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		for fast < len(nums) && nums[fast] == val {
			fast++
		}
	}
	return slow
}
