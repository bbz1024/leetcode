package test

func singleNumber(nums []int) int {
	// 异运算，相同的数异为0，0与任何数异为任何数
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}
