package greedy

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var res int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] > prices[i+1] {
			continue
		}
		res += prices[i+1] - prices[i]
	}
	return res
}
