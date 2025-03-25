package greedy

// timeout
func canCompleteCircuit(gas []int, cost []int) int {
	capacity := 0
	n := len(gas)
	for i := 0; i < n; i++ {
		ok := true

		// --------------- 优化 ---------------
		if i > 1 {
			last := gas[i-1] - cost[i-1]
			cur := gas[i] - cost[i]
			if cur <= last {
				continue
			}
		}
		// ---------------  ---------------

		for j := 0; j < n; j++ {
			start := (i + j) % n
			capacity += gas[start] - cost[start]
			if capacity < 0 {
				capacity = 0
				ok = false
				break
			}
		}
		if ok {
			return i
		}
	}
	return -1
}

// 3
func canCompleteCircuit2(gas []int, cost []int) int {
	n := len(gas)
	profit := make([]int, n)
	for i := 0; i < n; i++ {
		profit[i] = gas[i] - cost[i]
	}
	prof := 0
	start := 0
	stop := n
	for {
		// -1，-1,1
		start = start % n
		prof += profit[start]
		if prof < 0 {
			stop = start
		}
		start++
		if start == stop {
			return start
		}
	}
}
