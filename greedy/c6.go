package greedy

func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	var res int
	var cur int
	var next int
	for i := 0; i < n; i++ {
		next = max(next, i+nums[i])
		if cur == i {
			//说明超出了当前的覆盖范围，需要进行跳到下一个最大的覆盖范围
			res++
			cur = next
			if cur >= n-1 {
				return res
			}
		}
	}
	return -1
}
