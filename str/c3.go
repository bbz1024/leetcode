package str

import (
	"strings"
)

func reverseWords(s string) string {
	res := strings.Builder{}
	l := len(s)
	fast := l - 1
	slow := l - 1
	for fast >= 0 {
		for fast >= 0 && s[fast] != ' ' {
			fast--
		}
		res.WriteString(s[fast+1 : slow+1])
		if fast >= 0 {
			res.WriteByte(' ')
		}
		for fast >= 0 && s[fast] == ' ' {
			fast--
		}
		slow = fast
	}
	return res.String()
}
