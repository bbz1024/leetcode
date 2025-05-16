package test

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	var res string
	tmap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tmap[t[i]]++
	}
	smap := make(map[byte]int)
	left, right := 0, 0
	for right < len(s) {
		if right < len(s) && tmap[s[right]] > 0 {
			smap[s[right]]++
		}
		for IsContains(smap, tmap) && left <= right {
			if len(res) == 0 || right-left+1 < len(res) {
				res = s[left : right+1]
			}
			if smap[s[left]] > 1 {
				smap[s[left]]--
			} else {
				delete(smap, s[left])
			}
			left++
		}
		right++
	}

	return res
}
func IsContains(smap, tmap map[byte]int) bool {
	for k, v := range tmap {
		if smap[k] < v {
			return false
		}
	}
	return true
}
