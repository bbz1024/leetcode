package main

func sortColors(nums []int) {
	// 统计白
	// 统计蓝
	var arr [2]int
	for _, v := range nums {
		if v == 1 {
			arr[0]++
		}
		if v == 2 {
			arr[1]++
		}
	}
	red := len(nums) - arr[0] - arr[1]
	for i := 0; i < red; i++ {
		nums[i] = 0
	}
	for i := red; i < red+arr[0]; i++ {
		nums[i] = 1
	}
	for i := red + arr[0]; i < len(nums); i++ {
		nums[i] = 2
	}
}
