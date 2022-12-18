package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//node3f := ListNode{3, nil}
	//node2f := ListNode{4, &node3f}
	//node1f := ListNode{2, &node2f}
	//
	//node3s := ListNode{4, nil}
	//node2s := ListNode{6, &node3s}
	//node1s := ListNode{5, &node2s}

	node7f := ListNode{9, nil}
	node6f := ListNode{9, &node7f}
	node5f := ListNode{9, &node6f}
	node4f := ListNode{9, &node5f}
	node3f := ListNode{9, &node4f}
	node2f := ListNode{9, &node3f}
	node1f := ListNode{9, &node2f}

	//	node3s := ListNode{9, nil}
	node4s := ListNode{9, nil}
	node3s := ListNode{9, &node4s}
	node2s := ListNode{9, &node3s}
	node1s := ListNode{9, &node2s}

	//	fmt.Println(AddTwoNumbers(&node1f, &node1s))
	result := AddTwoNumbers(&node1f, &node1s)

	fmt.Println(result.Val)

	//go context.WithTimeout(context.Background(), 1)

	for {
		result = result.Next
		fmt.Println(result.Val)
		if result.Next == nil {
			return
		}
	}

	//ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond*1)
	//defer cancel()
	//
	//select {
	//case <-ctx.Done():
	//	return
	//}

}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	answer, _ := AddTwoNumbersWithRemainder(l1, l2, 0)
	return answer
}

func AddTwoNumbersWithRemainder(l1 *ListNode, l2 *ListNode, remainder int) (*ListNode, int) {
	answerTail := &ListNode{
		Val:  0,
		Next: nil,
	}

	if l1.Val+l2.Val+remainder < 10 {
		answerTail.Val = l1.Val + l2.Val + remainder
		remainder = 0
	} else {
		answerTail.Val = l1.Val + l2.Val - 10 + remainder
		remainder = 1
	}

	if l1.Next != nil && l2.Next != nil {
		answerTail.Next, remainder = AddTwoNumbersWithRemainder(l1.Next, l2.Next, remainder)
	} else if l1.Next != nil {
		answerTail.Next, remainder = AddTwoNumbersWithRemainder(l1.Next, &ListNode{
			Val:  0,
			Next: nil,
		}, remainder)
	} else if l2.Next != nil {
		answerTail.Next, remainder = AddTwoNumbersWithRemainder(&ListNode{
			Val:  0,
			Next: nil,
		}, l2.Next, remainder)
	} else {
		if remainder != 0 {
			answerTail.Next = &ListNode{Val: remainder, Next: nil}
		}
	}

	return answerTail, remainder
}
