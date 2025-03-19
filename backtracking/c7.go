package backtracking

func subsets(nums []int) [][]int {
	var res [][]int
	var backtrack func(path []int, start int)
	backtrack = func(path []int, start int) {
		p := make([]int, len(path))
		copy(p, path)
		res = append(res, p)
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(path, i+1)
			path = path[:len(path)-1]
		}
	}
	backtrack([]int{}, 0)
	return res
}
