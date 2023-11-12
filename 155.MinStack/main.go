package main

import (
	"sort"
)

func main() {

}

type MinStack struct {
	stack       []int
	min         int
	stackSorted []int
}

func Constructor() MinStack {
	return MinStack{stack: []int{}, min: 0, stackSorted: []int{}}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if val < this.min || len(this.stack) == 1 {
		this.min = val
	}
}

func (this *MinStack) Pop() {
	if this.min == this.Top() {
		this.stackSorted = make([]int, len(this.stack))
		copy(this.stackSorted, this.stack)
		sort.Ints(this.stackSorted)
		if len(this.stackSorted) > 1 {
			this.min = this.stackSorted[1]
		}
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
