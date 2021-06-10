package stack

import "math"

type MinStack struct {
	data []int
	min  int
}

/** initialize your data structure here. */
func MinStackConstructor() MinStack {
	return MinStack{min: math.MaxInt32}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	if val < this.min {
		this.min = val
	}
}

func (this *MinStack) Pop() {
	val := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	if val == this.min {
		this.min = math.MaxInt32
		for _, v := range this.data {
			if this.min > v {
				this.min = v
			}
		}
	}
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
