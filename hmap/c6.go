package hmap

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	var res [][]int
	for i := 0; i < l; i++ {

		for j := i + 1; j < l; j++ {

			left := j + 1
			right := l - 1

			for left < right {
				if nums[i]+nums[j]+nums[left]+nums[right] > target {
					right--
				} else if nums[i]+nums[j]+nums[left]+nums[right] < target {
					left++
				} else {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
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
			for j < l-1 && nums[j] == nums[j+1] {
				j++
			}
		}
		// 去重
		for i < l-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

/*

-2,-1,0,0,1,2
*/
