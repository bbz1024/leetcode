package dynamic

import "testing"

func TestNumTrees(t *testing.T) {
	t.Run("test numTrees", func(t *testing.T) {
		n := 3
		res := numTrees(n)
		if res != 5 {
			t.Errorf("numTrees(%d) = %d, want %d", n, res, 5)
		}
	})
}
