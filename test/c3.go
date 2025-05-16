package test

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	var res [][]int
	sort.Ints(nums)
	i := 0
	for i < len(nums)-2 {
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				// 去重
				// j 去重
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				// k 去重
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
		for i < len(nums)-2 && nums[i] == nums[i+1] {
			i++
		}
		i++
	}
	return res
}
