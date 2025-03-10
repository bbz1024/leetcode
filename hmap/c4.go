package hmap

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var cnt int
	mp := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums1); j++ {
			mp[nums1[i]+nums2[j]]++
		}
	}
	// -2,0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums1); j++ {
			val := nums3[i] + nums4[j]
			if _, ok := mp[-val]; ok {
				cnt += mp[-val]
			}
		}
	}
	return cnt
}
