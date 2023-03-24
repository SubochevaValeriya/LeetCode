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
	headOdd, headEven, tailOdd, tailEven := new(ListNode), new(ListNode), new(ListNode), new(ListNode)
	headOdd = tailOdd
	headEven = tailEven
	odd := true

	for i := head; i != nil; i = i.Next {
		tmp := &ListNode{
			Val:  i.Val,
			Next: nil,
		}
		switch {
		case odd:
			tailOdd.Next = tmp
			tailOdd = tmp
		case !odd:
			tailEven.Next = tmp
			tailEven = tmp
		}
		odd = !odd
	}

	tailOdd.Next = headEven.Next
	return headOdd.Next
}
