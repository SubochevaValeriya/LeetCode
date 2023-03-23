package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(listToSlice(oddEvenList(makeList(slice))))
	fmt.Println(slice)
	fmt.Println(listToSlice(makeList(slice)))
	//fmt.Println(listToSlice(makeList(slice)))

}

func makeList(slice []int) *ListNode {

	list := &ListNode{
		Val:  0,
		Next: nil,
	}

	var last = list
	for _, num := range slice {
		last.Next = &ListNode{
			Val:  num,
			Next: nil,
		}

		last = last.Next
	}

	return list.Next
}

func listToSlice(list *ListNode) []int {

	slice := []int{}

	for i := list; i != nil; i = i.Next {
		slice = append(slice, i.Val)
	}

	return slice
}

func oddEvenList(head *ListNode) *ListNode {
	fistHeadOdd, firstHeadEven, headOdd, headEven := new(ListNode), new(ListNode), new(ListNode), new(ListNode)
	fistHeadOdd = headOdd
	firstHeadEven = headEven
	odd := true

	for i := head; i != nil; i = i.Next {
		switch {
		case odd:
			headOdd.Next = &ListNode{
				Val:  i.Val,
				Next: nil,
			}
			headOdd = headOdd.Next
			odd = false
		case !odd:
			headEven.Next = &ListNode{
				Val:  i.Val,
				Next: nil,
			}
			headEven = headEven.Next
			odd = true
		}
	}

	headOdd.Next = firstHeadEven.Next
	return fistHeadOdd.Next
}
