package queue_stack

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	var window []int
	//window = maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	window = maxSlidingWindow([]int{1, -3, -1, -3, 5, 3, 6, 7}, 3)
	fmt.Println(window)
}
