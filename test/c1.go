package test

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	var res string
	l := 1
	// 中间扩散
	for i := 0; i < len(s); i++ {
		// babad
		// aabaab
		left := i - 1
		right := i + 1
		for left >= 0 && s[left] == s[i] {
			left--
			l++
		}
		for right < len(s) && s[right] == s[i] {
			right++
			l++
		}
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
			l += 2
		}
		if l > len(res) {
			// left 是最后一个 不匹配的左位置
			// right也是 最后一个 不匹配的右位置 且s[left:right]的right属于左闭右开
			res = s[left+1 : right]
		}
		l = 1
	}

	return res
}
