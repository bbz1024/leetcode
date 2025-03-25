package greedy

import "sort"

func largestSumAfterKNegations(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			nums[i] = -nums[i]
			k--
			continue
		}
		break
	}
	sort.Ints(nums)
	if k > 0 {
		if k%2 == 1 {
			nums[0] = -nums[0]
		}
	}
	return sum(nums)
}
func sum(nums []int) int {
	sumVal := 0
	for i := 0; i < len(nums); i++ {
		sumVal += nums[i]
	}
	return sumVal
}
