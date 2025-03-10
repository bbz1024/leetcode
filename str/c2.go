package str

func reverseStr(s string, k int) string {
	l := len(s)
	cnt := l / (2 * k)
	b := []byte(s)
	/*
		12345678
		0
	*/
	for i := 0; i < cnt; i++ {
		start := i * 2 * k
		reverseString(b[start : start+k])
	}
	mod := l % (2 * k)
	if mod < k {
		reverseString(b[l-mod:])
	} else if mod >= k {
		reverseString(b[l-mod : l-mod+k])
	}

	return string(b)
}
