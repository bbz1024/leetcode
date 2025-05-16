package codetop

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	l := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[l] = nums1[i]
			i--
		} else {
			nums1[l] = nums2[j]
			j--
		}
		l--
	}
	if i >= 0 {
		for i >= 0 {
			nums1[l] = nums1[i]
			i--
			l--
		}
	}
	if j >= 0 {
		for j >= 0 {
			nums1[l] = nums2[j]
			j--
			l--
		}
	}

}
