package str

import "testing"

func TestGetNext(t *testing.T) {
	//	a，a，b，a，a，f
	println(strStr("abac", "ac"))
}

// a，a，b，a，a，b，a，a，f
// 0,1,0,1,2,0
// a，a，b，a，a，f
