package main

import "fmt"

func main() {

	node1 := ListNode{2, nil}
	node2 := ListNode{1, &node1}
	node1 = ListNode{2, &node2}

	fmt.Println(hasCycle(&node1))

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// Task:
	// if cycle in any place - then true
	// time complexity - O(n)
	// memory usage - O(1)
	// corner cases: 1 node - false, link to yourself - true, head == nil

	if head == nil || head.Next == nil {
		return false
	}

	tortoise, hare := head, head

	for hare != nil && hare.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next
		if hare == tortoise {
			return true
		}
	}

	return false
}

func hasCycleBase(head *ListNode) bool {
	// Task:
	// if cycle in any place - then true
	// time complexity - O(n)
	// memory usage - O(1)
	// corner cases: 1 node - false, link to yourself - true, head == nil

	if head == nil || head.Next == nil {
		return false
	}

	m := map[*ListNode]*ListNode{}

	for ; head.Next != nil; head = head.Next {

		if _, ok := m[head.Next]; ok {
			return true
		}
		m[head] = m[head.Next]
	}

	return false
}
