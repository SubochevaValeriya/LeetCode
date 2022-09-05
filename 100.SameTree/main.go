package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSameTreeFirst(p *TreeNode, q *TreeNode) bool {
	if p.Val != q.Val {
		return false
	}

	if p.Left != nil && q.Left != nil {
		isSameTree(p.Left, q.Left)
	} else if (p.Left == nil && q.Left != nil) || (q.Left == nil && p.Left != nil) {
		return false
	}

	if p.Right != nil && q.Right != nil {
		isSameTree(p.Right, q.Right)
	} else if (p.Right == nil && q.Right != nil) || (q.Right == nil && p.Right != nil) {
		return false
	}

	return true
}
