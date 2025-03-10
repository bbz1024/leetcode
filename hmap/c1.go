package hmap

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mp := [26]int{}
	for _, v := range s {
		mp[v-'a'] += 1
	}
	for _, v := range t {
		mp[v-'a'] -= 1
		if mp[v-'a'] < 0 {
			return false
		}
	}
	return true
}
