package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type linkedList struct {
	head   *ListNode
	lenght int
}

func main() {
	node3f := ListNode{4, nil}
	node2f := ListNode{2, &node3f}
	node1f := ListNode{1, &node2f}

	node3s := ListNode{4, nil}
	node2s := ListNode{3, &node3s}
	node1s := ListNode{1, &node2s}

	head := MergeTwoLists(&node1f, &node1s)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

	fmt.Println(MergeTwoLists(&node1f, &node1s))

}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head, tail *ListNode

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		head = list1
		tail = list1
		list1 = list1.Next
	} else {
		head = list2
		tail = list2
		list2 = list2.Next
	}

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			tail = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			tail = list2
			list2 = list2.Next
		}
	}

	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	fmt.Println(tail)

	return head
}
