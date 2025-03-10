package hmap

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	mp := make(map[int]struct{})
	if len(nums1) > len(nums2) {
		for _, v := range nums1 {
			mp[v] = struct{}{}
		}
		for _, v := range nums2 {
			if _, ok := mp[v]; ok {
				res = append(res, v)
				delete(mp, v)
			}
		}
	} else {
		for _, v := range nums2 {
			mp[v] = struct{}{}
		}
		for _, v := range nums1 {
			if _, ok := mp[v]; ok {
				res = append(res, v)
				delete(mp, v)
			}
		}

	}
	return res
}
