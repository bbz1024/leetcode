package main

import "math"

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	start := 0
	m := math.MaxInt
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target {
			m = min(m, i-start+1)
			sum -= nums[start]
			start++
		}

	}
	if m == math.MaxInt {
		return 0
	}
	return m
}
