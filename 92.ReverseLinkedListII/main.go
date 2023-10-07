package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var index int
	toReverse := make([]*ListNode, right-left+1)
	reversedList := head
	var endOfFirstPart *ListNode
	for ; head != nil; head = head.Next {
		index++
		if index == left-1 {
			endOfFirstPart = head
			if left == 1 {
				reversedList = endOfFirstPart
			}
		}
		if index < left {
			continue
		}

		if index == left {
			toReverse[index-left] = &ListNode{
				Val:  head.Val,
				Next: nil,
			}
		} else {
			toReverse[index-left] = &ListNode{
				Val:  head.Val,
				Next: toReverse[index-left-1],
			}
		}

		if index == right {
			if endOfFirstPart == nil {
				endOfFirstPart = toReverse[len(toReverse)-1]
			} else {
				endOfFirstPart.Next = toReverse[len(toReverse)-1]
			}
			toReverse[0].Next = head.Next
			break
		}
	}

	return reversedList
}
