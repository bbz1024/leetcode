package test

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	var res int
	left := 0
	right := len(height) - 1
	for left < right {
		h := min(height[left], height[right])
		res = max(res, h*(right-left))

		// 移动长度小的边
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}
