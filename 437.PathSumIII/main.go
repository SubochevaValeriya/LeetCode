package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	return count(root, targetSum, 0, 0) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func count(root *TreeNode, targetSum, sum, paths int) int {
	if root == nil {
		return paths
	}

	if sum+root.Val == targetSum {
		paths++
	}

	paths = count(root.Left, targetSum, sum+root.Val, paths)
	paths = count(root.Left, targetSum, sum+root.Val, paths)

	return paths
}

func dfs(root *TreeNode, targetSum, paths int) int {
	paths = count(root, targetSum, 0, paths)
	dfs(root.Left, targetSum, paths)
}
