package queue_stack

/*
type MyQueue struct {
	in, out []int
}

func Constructor() MyQueue {
	return MyQueue{
		in:  make([]int, 0),
		out: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	for i := len(this.in) - 1; i >= 0; i-- {
		this.out = append(this.out, this.in[i])
	}
	this.in = this.in[:0]

	res := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	for i := len(this.out) - 1; i >= 0; i-- {
		this.in = append(this.in, this.out[i])
	}
	this.out = this.out[:0]
	return res
}

func (this *MyQueue) Peek() int {
	for i := len(this.in) - 1; i >= 0; i-- {
		this.out = append(this.out, this.in[i])
	}
	this.in = this.in[:0]
	res := this.out[len(this.out)-1]
	for i := len(this.out) - 1; i >= 0; i-- {
		this.in = append(this.in, this.out[i])
	}
	this.out = this.out[:0]
	return res
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}


*/
