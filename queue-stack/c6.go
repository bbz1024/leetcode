package queue_stack

/*
	func maxSlidingWindow(nums []int, k int) []int {
		var res []int
		// 暴力解法
		l := 0
		r := k
		for r <= len(nums) {
			res = append(res, getMax(nums[l:r]))
			l++
			r++
		}
		return res
	}

	func getMax(nums []int) int {
		m := nums[0]
		for i := 1; i < len(nums); i++ {
			m = max(m, nums[i])
		}
		return m
	}
*/
type IncrQueue struct {
	queue []int
}

func (i *IncrQueue) IsEmpty() bool {
	return len(i.queue) == 0
}
func (i *IncrQueue) GetMax() int {
	if i.IsEmpty() {
		return -1
	}
	return i.queue[0]
}
func (i *IncrQueue) Push(v int) {
	// 动态维护单调队列
	for !i.IsEmpty() && i.queue[len(i.queue)-1] < v {
		i.queue = i.queue[:len(i.queue)-1] // 删除最后一个元素
	}
	i.queue = append(i.queue, v)
}
func (i *IncrQueue) Pop(v int) {
	if !i.IsEmpty() && i.queue[0] == v {
		i.queue = i.queue[1:]
	}
	return
}
func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	if len(nums) < k {
		return res
	}
	queue := &IncrQueue{}
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	res = append(res, queue.GetMax())

	for i := k; i < len(nums); i++ {
		queue.Push(nums[i])
		queue.Pop(nums[i-k])
		res = append(res, queue.GetMax())
	}
	return res
}
