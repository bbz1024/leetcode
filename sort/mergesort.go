package sort

func MergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	MergeSort(nums, l, mid)
	MergeSort(nums, mid+1, r)
	merge(nums, l, mid, r)
}

func merge(nums []int, l int, mid int, r int) {
	// 此时 [l:mid]左边已经排好序了
	// [mid+1:r]右边已经排好序了
	// 进行合并
	temp := make([]int, r-l+1)
	i, j, k := l, mid+1, 0
	for i <= mid && j <= r {
		if nums[i] < nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++
		}
		k++
	}
	// 剩余的数据
	for i <= mid {
		temp[k] = nums[i]
		i++
		k++
	}
	for j <= r {
		temp[k] = nums[j]
		j++
		k++
	}
	for i := l; i <= r; i++ {
		nums[i] = temp[i-l]
	}
}
