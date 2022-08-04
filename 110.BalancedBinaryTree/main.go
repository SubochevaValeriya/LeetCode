package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//left3 := TreeNode{8, nil, nil}
	//left2 := TreeNode{6, nil, nil}
	//right1 := TreeNode{5, nil, nil}
	//left1 := TreeNode{4, &left3, nil}
	//left := TreeNode{2, &left1, &right1}
	//right := TreeNode{3, &left2, nil}
	//root := TreeNode{1, &left, &right}

	right3 := TreeNode{8, nil, nil}
	left2 := TreeNode{7, nil, nil}
	right2 := TreeNode{6, nil, &right3}
	right1 := TreeNode{5, nil, nil}
	left1 := TreeNode{4, &left2, nil}
	left := TreeNode{2, &left1, &right1}
	right := TreeNode{3, nil, &right2}
	root := TreeNode{1, &left, &right}

	//left2 := TreeNode{4, nil, nil}
	//right2 := TreeNode{4, nil, nil}
	//left1 := TreeNode{3, &left2, &right2}
	//right1 := TreeNode{3, nil, nil}
	//left := TreeNode{2, &left1, &right1}
	//right := TreeNode{2, nil, nil}
	//root := TreeNode{1, &left, &right}

	//left1 := TreeNode{3, nil, nil}
	//right1 := TreeNode{3, nil, nil}
	//left := TreeNode{2, &left1, &right1}
	//right := TreeNode{2, nil, nil}
	//root := TreeNode{1, &left, &right}

	//	fmt.Println(BinaryTreeToSlice(&root))
	fmt.Print(isBalanced(&root))
}

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}

	left := 0
	right := 0

	if root.Left != nil {
		left = 1
		leftTree := root.Left
		left += howDeep(leftTree)
	}

	if root.Right != nil {
		right = 1
		rightTree := root.Right
		right += howDeep(rightTree)
	}

	if math.Abs(float64(left-right)) > 1 {
		return false
	}

	if isBalanced(root.Left) == false {
		return false
	} else {
		return isBalanced(root.Right)
	}

	return true
}

func howDeep(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 0
	}

	l := 0
	r := 0
	if root.Right != nil {
		r = howDeep(root.Right)
		r += 1
	}

	if root.Left != nil {
		l = howDeep(root.Left)
		l += 1
	}

	if l > r {
		return l
	}

	return r
}

func BinaryTreeToSlice(root *TreeNode) []int {
	ourRoot := root
	ourRootSlice := []int{}

	ourRootSlice = append(ourRootSlice, ourRoot.Val)
	ourRootSlice = append(ourRootSlice, ourRoot.Left.Val)
	ourRootSlice = append(ourRootSlice, ourRoot.Right.Val)

	ourRoot = ourRoot.Left
	for ourRoot.Left != nil {
		ourRootSlice = append(ourRootSlice, ourRoot.Left.Val)
		ourRootSlice = append(ourRootSlice, ourRoot.Right.Val)
		ourRoot = ourRoot.Left
	}

	ourRoot = root.Right

	for ourRoot.Right != nil {
		ourRootSlice = append(ourRootSlice, ourRoot.Left.Val)
		ourRootSlice = append(ourRootSlice, ourRoot.Right.Val)
		ourRoot = ourRoot.Right

	}

	return ourRootSlice
}

func PrintingBinaryTree(root *TreeNode) {
	ourRoot := root

	fmt.Print(ourRoot.Val)
	fmt.Print(ourRoot.Left.Val)
	fmt.Print(ourRoot.Right.Val)
	ourRoot = ourRoot.Left
	for ourRoot.Left != nil {
		fmt.Print(ourRoot.Left.Val)
		fmt.Print(ourRoot.Right.Val)
		ourRoot = ourRoot.Left
	}

	ourRoot = root.Right

	for ourRoot.Right != nil {
		fmt.Print(ourRoot.Left.Val)
		fmt.Print(ourRoot.Right.Val)
		ourRoot = ourRoot.Right

	}
}
