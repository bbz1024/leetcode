package main

func sortedSquares(nums []int) []int {
	l := len(nums)
	for first := 0; first < l; first++ {
		nums[first] = nums[first] * nums[first]
	}
	newarr := make([]int, l)
	left := 0
	right := l - 1
	end := l - 1
	for left <= right {
		if nums[left] <= nums[right] {
			newarr[end] = nums[right]
			right--
		} else {
			newarr[end] = nums[left]
			left++
		}
		end--
	}
	return newarr
}
