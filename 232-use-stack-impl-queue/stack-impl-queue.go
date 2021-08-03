package _32_use_stack_impl_queue


type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}
/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.stack1 = append(this.stack1, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	for len(this.stack1) != 1 {
		top := this.stack1[len(this.stack1)-1]
		this.stack2 = append(this.stack2, top)
		this.stack1 = this.stack1[:len(this.stack1)-1]
	}
	ans := this.stack1[0]
	this.stack1 = this.stack1[:0]
	for len(this.stack2) != 0 {
		this.stack1 = append(this.stack1, this.stack2[len(this.stack2)-1])
		this.stack2 = this.stack2[:len(this.stack2)-1]
	}
	return ans
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	for len(this.stack1) != 1 {
		top := this.stack1[len(this.stack1)-1]
		this.stack2 = append(this.stack2, top)
		this.stack1 = this.stack1[:len(this.stack1)-1]
	}
	ans := this.stack1[0]
	for len(this.stack2) != 0 {
		this.stack1 = append(this.stack1, this.stack2[len(this.stack2)-1])
		this.stack2 = this.stack2[:len(this.stack2)-1]
	}
	return ans
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0
}
