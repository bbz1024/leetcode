package test

func hammingWeight(n int) int {
	var cnt int
	for n != 0 {
		if n&1 == 1 {
			cnt++
		}
		n >>= 1
	}
	return cnt
}
