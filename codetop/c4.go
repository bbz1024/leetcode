package codetop

func findKthLargest(nums []int, k int) int {
	l := len(nums)
	// 构建堆 （大顶堆） 从最后一个非叶子节点开始调整
	for i := l/2 - 1; i >= 0; i-- {
		adjust(nums, i, l)
	}
	for i := 0; i < k-1; i++ {
		// 堆顶和最后一个元素交换位置，然后调整堆，每次取数可以确保末尾的数是最大值
		nums[0], nums[l-1-i] = nums[l-1-i], nums[0]
		adjust(nums, 0, l-1-i)
	}
	return nums[0]
}

// 构建堆
func adjust(nums []int, start, l int) {
	for {
		left := start*2 + 1
		right := start*2 + 2
		p := start
		if left < l && nums[left] > nums[p] {
			p = left
		}
		if right < l && nums[right] > nums[p] {
			p = right
		}
		if start == p {
			break
		}
		nums[start], nums[p] = nums[p], nums[start]
		start = p
	}
}
