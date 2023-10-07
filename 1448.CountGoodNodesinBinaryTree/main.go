package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 1
	}

	return dfs(root.Left, root.Val, 0) + dfs(root.Right, root.Val, 0) + 1
}

func dfs(root *TreeNode, max int, count int) int {
	if root == nil {
		return count
	}
	if root.Val > max {
		max = root.Val
		count++
	}

	return dfs(root.Left, max, count-1) + dfs(root.Right, max, count-1)
}

var Count = 3

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 1
	}

	dfs(root.Left, root.Val)
	dfs(root.Right, root.Val)
	return Count + 1
}

func dfs(root *TreeNode, max int) {
	if root == nil {
		return
	}
	if root.Val > max {
		max = root.Val
		Count++
	}

	dfs(root.Left, max)
	dfs(root.Right, max)
}
