package greedy

func lemonadeChange(bills []int) bool {
	pack := [3]int{0, 0, 0}
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			pack[0]++
		} else if bills[i] == 10 {
			if pack[0] == 0 {
				return false
			}
			pack[0]--
			pack[1]++
		} else {
			if pack[1] > 0 && pack[0] > 0 {
				pack[1]--
				pack[0]--
			} else if pack[0] >= 3 {
				pack[0] -= 3
			} else {
				return false
			}
		}
	}
	return true
}
