package main

import "fmt"

func main() {

	//b := TreeNode{
	//	Val:   2,
	//	Left:  nil,
	//	Right: nil,
	//}
	//
	//c := TreeNode{
	//	Val:   3,
	//	Left:  nil,
	//	Right: nil,
	//}
	//
	//a := TreeNode{
	//	Val:   1,
	//	Left:  &b,
	//	Right: &c,
	//}

	c := TreeNode{
		Val:   -3,
		Left:  nil,
		Right: nil,
	}

	a := TreeNode{
		Val:   -2,
		Left:  nil,
		Right: &c,
	}

	fmt.Println(hasPathSum(&a, -5))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum = targetSum - root.Val

	if root.Left == nil && root.Right == nil && targetSum == 0 {
		return true
	}

	if root.Left != nil {
		if hasPathSum(root.Left, targetSum) == true {
			return true
		}
	}

	if root.Right != nil {
		if hasPathSum(root.Right, targetSum) == true {
			return true
		}
	}

	return false
}
