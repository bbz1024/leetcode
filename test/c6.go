package test

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	var res int
	var stack []int
	for i := 0; i < len(height); i++ {
		if i == 0 {
			stack = append(stack, i)
			continue
		}
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			// left 和 i 肯定是大于top的
			h := min(height[left], height[i]) - height[top]
			w := i - left - 1
			res += h * w
		}
		stack = append(stack, i)
	}
	return res
}
