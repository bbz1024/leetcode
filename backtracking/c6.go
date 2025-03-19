package backtracking

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var res []string
	var backtrack func(start int, path []string)
	backtrack = func(start int, path []string) {
		if start == len(s) && len(path) == 4 { // 叶子节点
			res = append(res, strings.Join(path, "."))
			return
		}
		for i := start; i < len(s); i++ {
			Str := s[start : i+1]
			toNum, err := strconv.Atoi(Str)
			if err != nil {
				continue
			}
			if toNum > 255 || len(Str) > 1 && Str[0] == '0' {
				continue
			}
			path = append(path, Str)
			backtrack(i+1, path)
			path = path[:len(path)-1]
		}

	}
	backtrack(0, []string{})
	return res
}
