package test

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestTrap(t *testing.T) {
	trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
}
func TestScan(t *testing.T) {
	reader := bufio.NewReader(os.Stdin)
	var m, L, R int
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &L, &R) // 逐行快速读取测试用例[2](@ref)
	}
}
