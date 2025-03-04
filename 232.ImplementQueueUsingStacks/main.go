package main

import "fmt"

func main() {
	s := Constructor()

	s.Push(1)
	fmt.Println(s)

	s.Push(2)
	fmt.Println(s)

	s.Push(3)
	fmt.Println(s)
}

type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{stack1: []int{}, stack2: []int{}}
}

func (this *MyQueue) Push(x int) {
	for i := 0; i <= len(this.stack1)-1; i++ {
		this.stack2 = append(this.stack2, this.stack1[i])
	}
	this.stack1 = []int{}
	this.stack1 = append(this.stack1, x)
	for i := 0; i <= len(this.stack2)-1; i++ {
		this.stack1 = append(this.stack1, this.stack2[i])
	}
	this.stack2 = []int{}
}

func (this *MyQueue) Pop() int {
	value := this.stack1[len(this.stack1)-1]
	this.stack1 = this.stack1[:len(this.stack1)-1]

	return value
}

func (this *MyQueue) Peek() int {
	return this.stack1[len(this.stack1)-1]
}

func (this *MyQueue) Empty() bool {
	if len(this.stack1) == 0 {
		return true
	}

	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
