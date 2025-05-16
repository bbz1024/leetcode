package sort

func QuickSort(arr []int, l, r int) {
	if l < r {
		mid := partition(arr, l, r)
		QuickSort(arr, l, mid-1)
		QuickSort(arr, mid+1, r)
	}
}
func partition(arr []int, l, r int) int {
	left, right := l, r
	for left < right {
		// 右边找第一个比基准小的
		for left < right && arr[l] <= arr[right] {
			right--
		}
		// 左边找第一个比基准大的
		for left < right && arr[l] >= arr[left] {
			left++
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	arr[l], arr[left] = arr[left], arr[l]
	return left
}
