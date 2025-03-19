package backtracking

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var backtrack func(nums []int, start int)
	backtrack = func(nums []int, start int) {
		if len(nums) == k {
			if sumArr(nums) == n {
				arr := make([]int, len(nums))
				copy(arr, nums)
				res = append(res, arr)
			}
			return
		}
		for i := start; i <= 9; i++ {
			nums = append(nums, i)
			backtrack(nums, i+1)
			nums = nums[:len(nums)-1]
		}
	}
	backtrack([]int{}, 1)
	return res
}
func sumArr(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}
