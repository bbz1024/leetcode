package codetop

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	mp := make(map[byte]int, l)
	start, end := 0, 0
	var res int
	for end < l {
		v, ok := mp[s[end]]
		if ok {
			start = max(start, v)
		}
		mp[s[end]] = end + 1
		res = max(res, end-start+1)
		end++
	}
	return res
}
