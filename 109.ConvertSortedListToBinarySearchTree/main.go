package main

import "fmt"

func main() {

	fmt.Println(*sortedListToBST(makeList([]int{0, 1, 2, 3, 4, 5})))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	list := []int{}
	for ; head != nil; head = head.Next {
		list = append(list, head.Val)
	}

	medium := len(list) / 2

	tree := &TreeNode{
		Val:   list[medium],
		Left:  nil,
		Right: nil,
	}

	tree.Left = populateTree(list[:medium])
	tree.Right = populateTree(list[medium+1:])

	return tree
}

func populateTree(list []int) *TreeNode {
	if len(list) == 0 {
		return nil
	}
	medium := len(list) / 2

	tree := &TreeNode{
		Val:   list[medium],
		Left:  nil,
		Right: nil,
	}

	tree.Left = populateTree(list[:medium])
	tree.Right = populateTree(list[medium+1:])

	return tree
}

func populateTreeRight(node *TreeNode, list []int) {
	if len(list) == 0 {
		return
	}
	node.Val = list[0]

	fmt.Println(node.Val, list, list[:len(list)-1])
	if len(list) == 1 {
		return
	}
	node.Right = &TreeNode{}
	populateTreeRight(node.Right, list[1:])
}

func populateTreeLeft(node *TreeNode, list []int) {
	if len(list) == 0 {
		return
	}
	node.Val = list[len(list)-1]

	fmt.Println(node.Val, list, list[:len(list)-1])
	if len(list) == 1 {
		return
	}
	node.Left = &TreeNode{}
	populateTreeLeft(node.Left, list[:len(list)-1])
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
