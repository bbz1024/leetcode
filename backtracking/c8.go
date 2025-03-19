package backtracking

import "sort"

func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums)
	var res [][]int
	var backtrack func(path []int, start int)
	used := make([]bool, len(nums))
	backtrack = func(path []int, start int) {
		p := make([]int, len(path))
		copy(p, path)
		res = append(res, p)
		for i := start; i < len(nums); i++ {
			// 去重，used[i-1] 表示当前位置的前一个位置是否被使用过，如果被使用过，则跳过当前位置
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			backtrack(path, i+1)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack([]int{}, 0)
	return res
}
