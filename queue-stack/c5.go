package queue_stack

import "strconv"

func evalRPN(tokens []string) int {
	mp := map[string]struct{}{
		"+": {},
		"-": {},
		"*": {},
		"/": {},
	}
	var stack []int
	for _, val := range tokens {
		if _, ok := mp[val]; ok {
			// val1
			val1 := stack[len(stack)-1]
			val2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch val {
			case "+":
				stack = append(stack, val2+val1)
			case "-":
				stack = append(stack, val2-val1)
			case "*":
				stack = append(stack, val2*val1)
			case "/":
				stack = append(stack, val2/val1)
			}
			continue
		}
		i, _ := strconv.ParseInt(val, 10, 64)

		stack = append(stack, int(i))
	}
	return stack[0]
}
