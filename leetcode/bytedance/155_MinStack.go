package main

import "math"

type MinStack struct {
	value []Value
	min   int
}

type Value struct {
	min int
	x   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{min: math.MaxInt32}
}

func (this *MinStack) Push(x int) {
	if x < this.min {
		this.min = x
	}
	this.value = append(this.value, Value{min: this.min, x: x})
}

func (this *MinStack) Pop() {
	if this.Len() > 0 {
		this.value = this.value[:this.Len()-1]
		if this.Len() > 0 {
			this.min = this.value[this.Len()-1].min
		} else {
			this.min = math.MaxInt32
		}
	}
}

func (this *MinStack) Top() int {
	if this.Len() > 0 {
		return this.value[this.Len()-1].x
	}
	return -1

}

func (this *MinStack) GetMin() int {
	return this.min
}

func (this *MinStack) Len() int {
	return len(this.value)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
