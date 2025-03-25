package greedy

import "strings"

func partitionLabels(s string) []int {
	if len(s) == 0 {
		return nil
	}
	var res []int
	for i := 0; i < len(s); i++ {
		idx := strings.LastIndex(s, string(s[i]))
		if idx == i {
			res = append(res, 1)
			continue
		}
		for j := i + 1; j < idx; j++ {
			if s[j] != s[idx] {
				idx = max(idx, strings.LastIndex(s, string(s[j])))
			}
		}
		res = append(res, idx-i+1)
		i = idx
	}
	return res

}

func partitionLabels2(s string) []int {
	if len(s) == 0 {
		return nil
	}
	var res []int
	mp := make([]int, 26)
	// 每次记录字符串中每个字符出现的最后位置
	for i := 0; i < len(s); i++ {
		mp[s[i]-'a'] = i
	}
	start := 0
	end := 0
	for i, v := range s {
		end = max(end, mp[v-'a'])
		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}
	return res

}
