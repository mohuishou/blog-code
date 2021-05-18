package stack

type CQueue struct {
	stack []int
	back  []int
}

func CQueueConstructor() CQueue {
	return CQueue{stack: []int{}, back: []int{}}
}

func (this *CQueue) AppendTail(value int) {
	for len(this.back) > 0 {
		// 出栈
		val := this.back[len(this.back)-1]
		this.back = this.back[:len(this.back)-1]

		this.stack = append(this.stack, val)
	}
	this.stack = append(this.stack, value)
}

func (this *CQueue) DeleteHead() int {
	for len(this.stack) > 0 {
		val := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		this.back = append(this.back, val)
	}
	if len(this.back) == 0 {
		return -1
	}
	val := this.back[len(this.back)-1]
	this.back = this.back[:len(this.back)-1]
	return val
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
