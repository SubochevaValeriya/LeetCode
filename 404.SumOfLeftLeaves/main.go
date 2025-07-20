package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	var sum int

	if root == nil {
		return sum
	}

	if root.Left != nil {
		sum += sumOfLeftLeaves(root.Left)
		if isLeaf(root.Left) {
			sum += root.Left.Val
		}
	}

	if root.Right != nil {
		sum += sumOfLeftLeaves(root.Right)
	}

	return sum
}

func isLeaf(r *TreeNode) bool {
	if r.Left == nil && r.Right == nil {
		return true
	}

	return false
}
