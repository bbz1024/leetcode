package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		a[i] = i
	}
	runtime.GC()
	fmt.Println(a)
}
