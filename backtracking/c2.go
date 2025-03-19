package backtracking

import "strings"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	mp := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var res []string
	var backtrack func(start int, path []string)
	backtrack = func(start int, path []string) {
		if len(path) == len(digits) {
			res = append(res, strings.Join(path, ""))
			return
		}
		for i := start; i < len(digits); i++ {
			for _, v := range mp[string(digits[i])] {
				path = append(path, string(v))
				backtrack(i+1, path)
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(0, []string{})
	return res
}
