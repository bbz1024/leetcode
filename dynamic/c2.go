package dynamic

func climbStairs(n int) int {
	/*
		1
		2
		前一层+前两层
		formula : pd[i] = pd[i-1] + pd[i-2]
	*/
	if n <= 1 {
		return n
	}
	pd := make([]int, n)
	pd[0] = 1 // 第一层
	pd[1] = 2 // 第二层
	for i := 2; i < len(pd); i++ {
		pd[i] = pd[i-1] + pd[i-2]
	}
	return pd[n-1]
}
