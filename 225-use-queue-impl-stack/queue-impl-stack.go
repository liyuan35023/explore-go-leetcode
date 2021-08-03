package _25_use_queue_impl_stack

type MyStack struct {
	queue1 []int
	queue2 []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		queue1: make([]int, 0),
		queue2: make([]int, 0),
	}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.queue1 = append(this.queue1, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	for len(this.queue1) != 0 {
		if len(this.queue1) == 1 {
			ans := this.queue1[0]
			this.queue1 = this.queue1[1:]
			return ans
		}
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:]
	}
	for len(this.queue2) != 0 {
		if len(this.queue2) == 1 {
			ans := this.queue2[0]
			this.queue2 = this.queue2[1:]
			return ans
		}
		this.queue1 = append(this.queue1, this.queue2[0])
		this.queue2 = this.queue2[1:]
	}
	return -1
}


/** Get the top element. */
func (this *MyStack) Top() int {
	for len(this.queue1) != 0 {
		if len(this.queue1) == 1 {
			ans := this.queue1[0]
			return ans
		}
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:]
	}
	for len(this.queue2) != 0 {
		if len(this.queue2) == 1 {
			ans := this.queue2[0]
			this.queue1 = append(this.queue1, this.queue2[0])
			this.queue2 = this.queue2[1:]
			return ans
		}
		this.queue1 = append(this.queue1, this.queue2[0])
		this.queue2 = this.queue2[1:]
	}
	return -1
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0 && len(this.queue2) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */