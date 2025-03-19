package backtracking

func partition(s string) [][]string {
	var res [][]string
	var backtrack func(start int, path []string)
	backtrack = func(start int, path []string) {
		if start == len(s) {
			p := make([]string, len(path))
			copy(p, path)
			res = append(res, p)
			return
		}
		for i := start; i < len(s); i++ {
			if isPalindrome(s[start : i+1]) {
				path = append(path, s[start:i+1])
				backtrack(i+1, path)
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(0, []string{})
	return res
}
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
