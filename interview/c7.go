package main

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	//	冒泡排序
	for i := 0; i < len(nums); i++ {
		isSwap := false
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] == 0 && nums[j+1] != 0 {
				isSwap = true
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		if !isSwap {
			break
		}
	}
}
