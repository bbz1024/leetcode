package str

import (
	"fmt"
	"testing"
)

func TestRepeatedSubstringPattern(t *testing.T) {
	pattern := repeatedSubstringPattern("babbabbabbabbab")
	fmt.Println(pattern)
}
