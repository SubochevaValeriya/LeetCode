package main

import "fmt"

func main() {

	b := &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}

	c := &TreeNode{
		Val:   20,
		Left:  nil,
		Right: nil,
	}
	a := &TreeNode{
		Val:   3,
		Left:  b,
		Right: c,
	}

	fmt.Println(zigzagLevelOrder(a))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeWithLevel struct {
	TreeNode *TreeNode
	Level    int
}

func zigzagLevelOrder(root *TreeNode) [][]int {

	answer := [][]int{}

	if root == nil {
		return answer
	}

	stack := []TreeNodeWithLevel{{
		TreeNode: root,
		Level:    1,
	}}

	for current := stack[0]; ; current = stack[0] {
		if len(answer) >= current.Level {
			if current.Level%2 == 1 {
				answer[current.Level-1] = append(answer[current.Level-1], current.TreeNode.Val)
			} else {
				newArr := append([]int{current.TreeNode.Val}, answer[current.Level-1]...)
				answer[current.Level-1] = newArr
			}
		} else {
			answer = append(answer, []int{current.TreeNode.Val})
		}

		leftNode := TreeNodeWithLevel{
			TreeNode: current.TreeNode.Left,
			Level:    current.Level + 1,
		}

		rightNode := TreeNodeWithLevel{
			TreeNode: current.TreeNode.Right,
			Level:    current.Level + 1,
		}

		if current.TreeNode.Left != nil {
			stack = append(stack, leftNode)
		}
		if current.TreeNode.Right != nil {
			stack = append(stack, rightNode)
		}

		stack = stack[1:]
		if len(stack) == 0 {
			break
		}
	}

	return answer
}

func zigzagLevelOrderOld(root *TreeNode) [][]int {

	answer := [][]int{}

	if root == nil {
		return answer
	}
	toRight := false

	stack := []TreeNodeWithLevel{{
		TreeNode: root,
		Level:    1,
	}}

	for current := stack[0]; ; current = stack[0] {
		if len(answer) >= current.Level {
			answer[current.Level-1] = append(answer[current.Level-1], current.TreeNode.Val)
		} else {
			answer = append(answer, []int{current.TreeNode.Val})
		}

		leftNode := TreeNodeWithLevel{
			TreeNode: current.TreeNode.Left,
			Level:    current.Level + 1,
		}

		rightNode := TreeNodeWithLevel{
			TreeNode: current.TreeNode.Right,
			Level:    current.Level + 1,
		}

		switch toRight {
		case true:
			if current.TreeNode.Left != nil {
				stack = append(stack, leftNode)
			}
			if current.TreeNode.Right != nil {
				stack = append(stack, rightNode)
			}
		case false:
			if current.TreeNode.Right != nil {
				stack = append(stack, rightNode)
			}
			if current.TreeNode.Left != nil {
				stack = append(stack, leftNode)
			}
		}
		toRight = !toRight
		stack = stack[1:]
		if len(stack) == 0 {
			break
		}
	}

	return answer
}
