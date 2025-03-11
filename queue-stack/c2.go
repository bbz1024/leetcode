package queue_stack

type MyStack struct {
	in, out []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyStack) Pop() int {
	for i := 0; i < len(this.in); i++ {
		this.out = append(this.out, this.in[i])
	}
	this.in = this.in[:0]
	res := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	for i := 0; i < len(this.out); i++ {
		this.in = append(this.in, this.out[i])
	}
	this.out = this.out[:0]
	return res
}

func (this *MyStack) Top() int {
	return this.in[len(this.in)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.in) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
