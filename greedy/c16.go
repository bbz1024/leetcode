package greedy

import "math"

func monotoneIncreasingDigits(n int) int {
	if n < 10 {
		return n
	}
	// 只要个位小于更大一位就需要个位变为9，更大位置的数字减1
	var res int
	var cnt int
	for {
		if isMonotoneIncreasingDigits(n) {
			n *= int(math.Pow10(cnt))
			return res + n
		}
		res += 9 * int(math.Pow10(cnt))
		n /= 10
		n -= 1
		cnt++
	}

}
func isMonotoneIncreasingDigits(n int) bool {
	pre := n % 10
	for n > 0 {
		cur := n % 10
		if cur > pre {
			return false
		}
		pre = cur
		n /= 10
	}
	return true
}
