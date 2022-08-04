package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	left1 := TreeNode{1, nil, nil}
	right1 := TreeNode{3, nil, nil}
	left2 := TreeNode{6, nil, nil}
	right2 := TreeNode{9, nil, nil}
	left := TreeNode{2, &left1, &right1}
	right := TreeNode{7, &left2, &right2}
	root := TreeNode{4, &left, &right}

	fmt.Println(BinaryTreeToSlice(&root))
	ourRoot := invertTree(&root)

	fmt.Println(BinaryTreeToSlice(ourRoot))

}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		invertTree(root.Left)
		invertTree(root.Right)
		return root
	}
	return root
}

//if root.Left == nil {
//	return root
//}
//currRoot := root
//
//currRoot.Left.Val, currRoot.Right.Val = currRoot.Right.Val, currRoot.Left.Val
//
//currRootLeft := currRoot.Left
//currRootRight := currRoot.Right
//
//for currRootLeft.Left != nil {
//	currRootLeft.Left.Val, currRootRight.Right.Val = currRootRight.Right.Val, currRootLeft.Left.Val
//	currRootLeft = currRootLeft.Left
//	currRootRight = currRootRight.Right
//}
//
//currRootLeft = currRoot.Left
//currRootRight = currRoot.Right
//
//for currRootLeft.Right != nil {
//	currRootLeft.Right.Val, currRootRight.Left.Val = currRootRight.Left.Val, currRootLeft.Right.Val
//	currRootLeft = currRootLeft.Right
//	currRootRight = currRootRight.Left
//}

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
