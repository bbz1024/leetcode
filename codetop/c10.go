package codetop

func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindrome(s[i : j+1]) {
				if len(s[i:j+1]) > len(res) {
					res = s[i : j+1]
				}
			}
		}
	}
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
func longestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}
	var res string
	strLen := len(s)
	for i := 0; i < strLen; i++ {
		l := 1
		left, right := i-1, i+1
		for left >= 0 && s[left] == s[i] {
			l++
			left--
		}
		for right < strLen && s[right] == s[i] {
			l++
			right++
		}
		for left >= 0 && right < strLen && s[left] == s[right] {
			l += 2
			left--
			right++
		}
		if l > len(res) {
			res = s[left+1 : right]
		}

	}
	return res
}
