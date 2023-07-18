package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {

	if root.Left == nil && root.Right == nil{
		return 1
	}

	if root.Left == nil {
		return dfs(root.Right, 1)
	}

	if root.Right == nil {
		return dfs(root.Left, 1)
	}

	return int(math.Min(float64(dfs(root.Left, 1)), float64(dfs(root.Right, 1))))
}

func dfs(root *TreeNode, level int) int {
	if root.Left == nil && root.Right == nil{
		return level+1
	}

	if root.Left == nil {
		return dfs(root.Right, level+1)
	}
	if root.Right == nil{
		return dfs(root.Right, level+1)
	}

	return int(math.Min(float64(dfs(root.Left, level+1)), float64(dfs(root.Right, level+1)))
}
