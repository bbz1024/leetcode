package codetop

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	minVal := prices[0]
	var res int
	for i := 1; i < len(prices); i++ {
		res = max(res, prices[i]-minVal)
		minVal = min(minVal, prices[i])
	}
	return res
}
