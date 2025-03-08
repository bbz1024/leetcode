package main

import (
	"fmt"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	matrix := generateMatrix(4)
	fmt.Println(matrix)
}
