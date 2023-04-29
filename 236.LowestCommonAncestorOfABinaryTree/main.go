package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == p || root == q {
		return root
	}

	ancestor, count := dfs(root.Left, p, q, 0)
	if count == 2 {
		return ancestor
	}
	ancestor, count = dfs(root.Right, p, q, 0)
	if count == 2 {
		return ancestor
	}
	return root
}

func dfs(root, p, q *TreeNode, count int) (*TreeNode, int) {
	if root == nil {
		return nil, count
	}
	if root == p || root == q {
		count++
		if count == 2 {
			return root, count
		}
	}

	anc, countLeft := dfs(root.Left, p, q, count)
	if count == 2 {
		if root == p || root == q {
			return root, countLeft
		} else {
			return anc, countLeft
		}
	}
	anc, countRight := dfs(root.Right, p, q, count)
	if countRight == 2 {
		if root == p || root == q {
			return root, countRight
		} else {
			return anc, countRight
		}
	}

	if countLeft-count == 1 && countRight-count == 1 {
		return root, 2
	}

	if countLeft == 1 || countRight == 1 {
		return root, 1
	}
	return root, 0
}
