package greedy

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	candies := make([]int, len(ratings))
	for i := 0; i < len(candies); i++ {
		candies[i] = 1
	}
	for i := 1; i < len(candies); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
			continue
		}
		// 大于需要进行向上调整
		// 1,4 3 2   1,2,1,
		for j := i; j >= 1; j-- {
			if candies[j-1] > candies[j] {
				break
			}
			if ratings[j-1] > ratings[j] {
				candies[j-1] = candies[j] + 1
			}
		}
	}
	sum := 0
	for _, v := range candies {
		sum += v
	}
	return sum
}

func candy2(ratings []int) (ans int) {
	n := len(ratings)
	left, right := make([]int, n), make([]int, n)

	for i := range left {
		left[i], right[i] = 1, 1
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i - 1] {
			left[i] = left[i - 1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i + 1] {
			right[i] = right[i + 1] + 1
		}
	}

	for i := range left {
		ans += max(left[i], right[i])
	}
	return
}