package backtracking

func permute(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}
	used := make(map[int]bool)
	var backtrack func(path []int)
	backtrack = func(path []int) {
		if len(path) == len(nums) {
			p := make([]int, len(path))
			copy(p, path)
			res = append(res, p)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			used[nums[i]] = true
			path = append(path, nums[i])
			backtrack(path)
			path = path[:len(path)-1]
			used[nums[i]] = false
		}
	}
	backtrack([]int{})
	return res
}

/*
 123


*/
