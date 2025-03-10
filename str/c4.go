package str

func strStr(haystack string, needle string) int {
	next := getNext(needle)
	l := len(haystack)
	for i, j := 0, 0; i < l; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
			if j == len(needle) {
				return i - j + 1
			}
		}
	}
	return -1

}

// a，a，b，a，a，f
/*
	i:2，j：1
*/
func getNext(p string) []int {
	next := make([]int, len(p))
	next[0] = 0
	for i := 1; i < len(p); i++ {
		j := next[i-1]
		for p[j] != p[i] {
			if j == 0 {
				break
			}
			j = next[j-1]
		}
		if p[j] == p[i] {
			next[i] = j + 1
		} else {
			next[i] = 0
		}
	}
	return next
}

// a，a，b，a，a，b，a，a，f
// a，a，b，a，a，f
