package codetop

import (
	"slices"
)

// 滑动窗口+去重
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	var res [][]int
	slices.Sort(nums)
	// [-1,-1,0,1,2]
	l := len(nums)
	for i := 0; i < l; i++ {
		left, right := i+1, l-1
		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 去重
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
		// i 去重
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
