package codetop

import "testing"

func TestMerge(t *testing.T) {
	nums1 := []int{1}
	nums2 := []int{}
	merge(nums1, 1, nums2, 0)
	t.Log(nums1)
}
