package codetop

import "math/rand"

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := rand.Intn(right-left+1) + left // 随机化中间值
	nums[mid], nums[left] = nums[left], nums[mid]
	pivot := nums[left] // 确定基准值

	lt, rt := left, right

	// 三指针分区法（处理重复元素优化）
	for i := left + 1; i <= rt; {
		if nums[i] < pivot {
			nums[i], nums[lt] = nums[lt], nums[i]
			lt++
			i++
		} else if nums[i] > pivot {
			nums[i], nums[rt] = nums[rt], nums[i]
			rt--
		} else {
			// 去掉相等的元素
			i++
		}
	}
	quickSort(nums, left, lt-1)
	quickSort(nums, rt+1, right)
}
