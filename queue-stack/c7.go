package queue_stack

import "container/heap"

type Node struct {
	val   int
	count int
}
type CounterHeap struct {
	counter []*Node
}

func (c *CounterHeap) Len() int {
	return len(c.counter)
}

func (c *CounterHeap) Less(i, j int) bool {
	return c.counter[i].count > c.counter[j].count // 改为最小堆
}

func (c *CounterHeap) Swap(i, j int) {
	c.counter[i], c.counter[j] = c.counter[j], c.counter[i]
}

func (c *CounterHeap) Push(x any) {
	c.counter = append(c.counter, x.(*Node))
}

func (c *CounterHeap) Pop() any {
	old := c.counter
	n := len(old)
	item := old[0]
	c.counter = old[1:n]
	return item
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	cheap := CounterHeap{
		counter: make([]*Node, 0),
	}
	heap.Init(&cheap)
	res := make([]int, k)
	for num, count := range m {
		heap.Push(&cheap, &Node{num, count})
	}

	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(&cheap).(*Node).val
	}

	return res
}
