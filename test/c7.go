package test

func canJump(nums []int) bool {
	l := len(nums)

	if l == 0 {
		return false
	}
	var maxCover int
	for i := 0; i < l; i++ {
		maxCover = max(maxCover, i+nums[i])
		if maxCover >= l-1 {
			return true
		}
		if maxCover <= i {
			break
		}
	}
	return false
}

/*
[2,3,1,1,4]

[3,2,1,0,4]


*/
