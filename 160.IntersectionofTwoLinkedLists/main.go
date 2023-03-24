package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	for i := headA; i != nil; i = i.Next {
		for j := headB; j != nil; j = j.Next {
			if i == j {
				return i
			}
		}
	}
	return nil
}
