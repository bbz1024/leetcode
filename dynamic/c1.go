package dynamic

func fib(n int) int {
	// 0,1,1,2,3
	if n <= 1 {
		return n
	}
	pd := make([]int, n+1)
	pd[0] = 0
	pd[1] = 1
	// formula：pd[i]=pd[i-1]+pd[i-2]
	for i := 2; i <= n; i++ {
		pd[i] = pd[i-1] + pd[i-2]
	}
	return pd[n]
}
