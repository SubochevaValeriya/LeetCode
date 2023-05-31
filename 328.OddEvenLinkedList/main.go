package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {

	headOdd, headEven, lastHeadOdd := &ListNode{}, &ListNode{}, &ListNode{}

	odd := true

	for i := head; i != nil; i = head.Next {
		switch {
		case odd:
			headOdd.Val = i.Val
			headOdd = headOdd.Next
			odd = false
			lastHeadOdd = head
		case !odd:
			headEven.Val = i.Val
			headEven = headEven.Next
			odd = true
		}
	}

	lastHeadOdd.Next = headEven.Next
	return headOdd
}
