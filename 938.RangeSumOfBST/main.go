package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	var sum int

	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}

	if root.Val > low {
		sum += rangeSumBST(root.Left, low, high)
	}

	if root.Val < high {
		sum += rangeSumBST(root.Right, low, high)
	}

	return sum
}
