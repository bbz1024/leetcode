package backtracking

func findSubsequences(nums []int) [][]int {
	var res [][]int

	if len(nums) == 0 {
		return res
	}
	var backtrack func(path []int, start int)
	backtrack = func(path []int, start int) {

		if len(path) > 1 {
			p := make([]int, len(path))
			copy(p, path)
			res = append(res, p)
		}
		used := make(map[int]struct{}, len(nums))

		for i := start; i < len(nums); i++ {
			if _, ok := used[nums[i]]; ok { // 去重
				continue
			}
			if len(path) == 0 || nums[i] >= path[len(path)-1] {
				used[nums[i]] = struct{}{}
				path = append(path, nums[i])
				backtrack(path, i+1)
				path = path[:len(path)-1]
			}

		}

	}
	backtrack([]int{}, 0)
	return res

}
