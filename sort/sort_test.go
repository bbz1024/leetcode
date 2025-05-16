package sort

import (
	"math/rand"
	"sort"
	"testing"
)

// quickly

func TestSortTest(t *testing.T) {
	var sample [][]int
	for i := 0; i < 5; i++ {
		arr := make([]int, 20)
		for j := 0; j < 20; j++ {
			arr[j] = rand.Intn(100)
		}
		sample = append(sample, arr)
	}
	t.Run("quickly sort", func(t *testing.T) {
		for _, arr := range sample {
			cpArr := make([]int, len(arr))
			copy(cpArr, arr)
			QuickSort(cpArr, 0, len(cpArr)-1)
			if !sort.IsSorted(sort.IntSlice(cpArr)) {
				t.Error("quickly sort fail")
			}
		}
		t.Log("quickly sort pass")
	})
	t.Run("heap sort", func(t *testing.T) {
		for _, arr := range sample {
			cpArr := make([]int, len(arr))
			copy(cpArr, arr)
			HeapSort(cpArr)
			if !sort.IsSorted(sort.IntSlice(cpArr)) {
				t.Error("heap sort fail")
			}
		}
	})
	t.Run("merge sort", func(t *testing.T) {
		for _, arr := range sample {
			cpArr := make([]int, len(arr))
			copy(cpArr, arr)
			MergeSort(cpArr, 0, len(cpArr)-1)
			if !sort.IsSorted(sort.IntSlice(cpArr)) {
				t.Error("merge sort fail")
			}
		}
	})
}
