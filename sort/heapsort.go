package sort

func HeapSort(arr []int) {
	// 构建堆，从最后一个非叶子节点开始调整
	leaf := len(arr)/2 - 1
	for start := leaf; start >= 0; start-- {
		adjust(arr, start, len(arr))
	}
	for i := 0; i < len(arr); i++ {
		arr[0], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[0]
		adjust(arr, 0, len(arr)-i-1)
	}
}

func adjust(arr []int, parent, length int) {
	left := 2*parent + 1
	right := 2*parent + 2
	temp := parent
	if left < length && arr[left] > arr[temp] {
		temp = left
	}
	if right < length && arr[right] > arr[temp] {
		temp = right
	}
	if temp != parent {
		arr[temp], arr[parent] = arr[parent], arr[temp]
		adjust(arr, temp, length)
	}
}
