package str

import "strings"

func repeatedSubstringPattern(s string) bool {
	return strings.Contains(s[1:]+s[:len(s)-1], s)
}
