package main

import (
	"fmt"
)

func main() {

	node1 := ListNode{2, nil}
	node2 := ListNode{1, &node1}
	node1 = ListNode{2, &node2}

	fmt.Println(detectCycle(&node1))

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	// Task:
	// if cycle in any place - then true
	// time complexity - O(n)
	// memory usage - O(1)
	// corner cases: 1 node - false, link to yourself - true, head == nil

	if head == nil || head.Next == nil {
		return nil
	}

	intersect := intersection(head)
	if intersect == nil {
		return nil
	}

	start := head

	for start != intersect {
		start = start.Next
		intersect = intersect.Next
	}

	return start
}

func intersection(head *ListNode) *ListNode {
	// Task:
	// if cycle in any place - then true
	// time complexity - O(n)
	// memory usage - O(1)
	// corner cases: 1 node - false, link to yourself - true, head == nil

	if head == nil || head.Next == nil {
		return nil
	}

	tortoise, hare := head, head

	for hare != nil && hare.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next
		if hare == tortoise {
			return tortoise
		}
	}
	return nil
}
