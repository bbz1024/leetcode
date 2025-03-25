package greedy

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}
	maxCover := 0
	for i := 0; i <= maxCover; i++ {
		maxCover = max(maxCover, i+nums[i])
		if maxCover >= len(nums)-1 {
			return true
		}
	}
	return false
}

// [1, 2, 2, 0, 0, 4]
// [3,2,1,0,4]
