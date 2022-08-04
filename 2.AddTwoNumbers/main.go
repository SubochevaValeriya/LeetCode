package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node3f := ListNode{3, nil}
	node2f := ListNode{4, &node3f}
	node1f := ListNode{2, &node2f}

	node3s := ListNode{4, nil}
	node2s := ListNode{6, &node3s}
	node1s := ListNode{5, &node2s}

	//fmt.Println(AddTwoNumbers(&node1f, &node1s))
	fmt.Println(AddTwoNumbers(&node1f, &node1s).Val)

}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	answer, _ := AddTwoNumbersWithRemainder(l1, l2, 0)
	return answer
}

func AddTwoNumbersWithRemainder(l1 *ListNode, l2 *ListNode, remainder int) (*ListNode, int) {
	var answer *ListNode
	answer = l1
	if l1.Val+l2.Val < 10 {
		answer.Val = l1.Val + l2.Val + remainder
	} else {
		answer.Val = l1.Val + l2.Val - 10 + remainder
		remainder += 1
	}

	fmt.Println(answer.Val, remainder)

	if l1.Next != nil {
		answer.Next, remainder = AddTwoNumbersWithRemainder(l1.Next, l2.Next, remainder)
	} else {
		answer.Val = l1.Val + l2.Val + remainder
	}

	return answer, remainder
}
