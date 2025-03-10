package hmap

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := mp[target-nums[i]]
		if ok {
			return []int{mp[target-nums[i]], i}
		} else {
			mp[nums[i]] = i
		}
	}
	return []int{}
}
