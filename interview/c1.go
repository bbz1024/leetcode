package main

import (
	"fmt"
)

// 计算小明能分到的最大苹果数
func maxApplesForXiaoming(m, n, k int) int {
	// 二分查找的左右边界
	left, right := 1, m
	result := 0

	// 二分查找
	for left <= right {
		mid := left + (right-left)/2
		// 计算当前假设小明分到 mid 个苹果时，总共需要的苹果数
		total := totalApples(n, k, mid)
		if total <= m {
			// 如果所需苹果数不超过总苹果数，更新结果并尝试更大的值
			result = mid
			left = mid + 1
		} else {
			// 所需苹果数超过总苹果数，尝试更小的值
			right = mid - 1
		}
	}
	return result
}

// 计算在小明分到 x 个苹果时，总共需要的苹果数
func totalApples(n, k, x int) int {
	total := 0
	// 计算小明左边的小孩所需苹果数
	for i := k - 1; i >= 1; i-- {
		if x-(k-i) < 1 {
			total += 1
		} else {
			total += x - (k - i)
		}
	}
	// 计算小明自己的苹果数
	total += x
	// 计算小明右边的小孩所需苹果数
	for i := k + 1; i <= n; i++ {
		if x-(i-k) < 1 {
			total += 1
		} else {
			total += x - (i - k)
		}
	}
	return total
}

func main() {
	// 示例输入
	m := 20
	n := 5
	k := 3
	// 计算小明能分到的最大苹果数
	result := maxApplesForXiaoming(m, n, k)
	fmt.Printf("小明最多能分到 %d 个苹果\n", result)
}
