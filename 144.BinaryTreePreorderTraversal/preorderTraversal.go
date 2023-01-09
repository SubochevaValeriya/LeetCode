package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {

	answer := []int{}

	for root != nil {
		answer = append(answer, root.Val)

		answer = append(answer, preorderTraversal(root.Left)...)
		answer = append(answer, preorderTraversal(root.Right)...)
	}

	return answer
}
