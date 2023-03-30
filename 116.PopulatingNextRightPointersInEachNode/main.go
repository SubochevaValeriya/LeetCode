package main

import "fmt"

func main() {

	b := Node{
		Val:   2,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}

	c := Node{
		Val:   3,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}

	a := Node{
		Val:   1,
		Left:  &b,
		Right: &c,
		Next:  nil,
	}

	fmt.Println(connect(&a).Left.Next)

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	stack := []*Node{root}

	currentLevel, nodesLast := 1, 1

	for len(stack) != 0 {
		nodesLast--
		if len(stack) != 1 && nodesLast != 0 {
			if stack[0].Left != stack[1] {
				stack[0].Next = stack[1]
			}
		}

		if stack[0].Left != nil {
			stack = append(stack, stack[0].Left, stack[0].Right)
		}

		stack = stack[1:]
		if nodesLast == 0 {
			currentLevel = currentLevel * 2
			nodesLast = currentLevel
		}

	}

	return root
}

//stack := []*Node{root}
//
//for len(stack) != 0 {
//if len(stack) != 1 {
//if stack[0].Left != stack[1] {
//stack[0].Right = stack[1]
//}
//}
//
//if stack[0].Left != nil {
//stack = append(stack, stack[0].Left, stack[0].Right)
//}
//
//stack = stack[1:]
//}
