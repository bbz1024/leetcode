package codetop

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	// 从第一墙开始
	var res int
	var stack []int
	for i := 0; i < len(height); i++ {

		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			h := min(height[left], height[i]) - height[top]
			w := i - left - 1
			res += h * w
		}
		stack = append(stack, i)

	}
	return res
}
