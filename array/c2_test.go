package main

import (
	"fmt"
	"testing"
)

// 3,2,2,3
// 0,
func TestRemoveElement(t *testing.T) {
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	removeElement(arr, 2)
	fmt.Println(arr)
}
