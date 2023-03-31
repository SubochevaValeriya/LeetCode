package main

import (
	"fmt"
)

func main() {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*1))
	//defer cancel()

	sl := []int{9, -42, -42, 0, 76, 76, 0, 0, 13, 0, 13}
	//go func() {
	//	fmt.Println(treeNodeToInt(generateTreeNode(sl), len(sl)))
	//}()

	fmt.Println(isSymmetric(generateTreeNode(sl)))
	//select {
	//case <-ctx.Done():
	//	fmt.Printf("Context cancelled: %v\n", ctx.Err())
	//}

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left == nil || root.Right == nil {
		return false
	}

	stackLeft := []*TreeNode{root.Left}
	stackRight := []*TreeNode{root.Right}

	for len(stackLeft) != 0 && len(stackRight) != 0 {

		left := stackLeft[0]
		right := stackRight[0]

		if (left != nil && right == nil) || (left == nil && right != nil) {
			return false
		}

		if left != nil && right != nil {
			if left.Val != right.Val {
				return false
			}
		}

		if left != nil {
			stackLeft = append(stackLeft, left.Left, left.Right)
		}
		if stackRight[0] != nil {
			stackRight = append(stackRight, right.Right, right.Left)
		}
		stackLeft = stackLeft[1:]
		stackRight = stackRight[1:]
	}

	return true
}

func generateTreeNode(nums []int) *TreeNode {
	if len(nums) == 0 {
		return &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		}
	}

	if len(nums) == 1 {
		return &TreeNode{
			Val:   nums[0],
			Left:  nil,
			Right: nil,
		}
	}

	treeNodes := make([]*TreeNode, len(nums))

	ind := -1
	leftEmpty := true

	for i, num := range nums {
		a := &TreeNode{
			Val:   num,
			Left:  nil,
			Right: nil,
		}
		treeNodes[i] = a
		if ind != -1 {
			switch {
			case leftEmpty:
				treeNodes[ind].Left = a
				leftEmpty = !leftEmpty
			case !leftEmpty:
				treeNodes[ind].Right = a
				leftEmpty = !leftEmpty
				ind++
			}
		} else {
			ind++
		}

	}

	return treeNodes[0]
}

func treeNodeToInt(root *TreeNode, n int) []int {
	if root == nil {
		return []int{}
	}

	slice := []int{}
	stack := []*TreeNode{root}

	for len(stack) != 0 && len(slice) < n {
		fmt.Println(slice)
		if stack[0] != nil {
			slice = append(slice, stack[0].Val)
			stack = append(stack, stack[0].Left, stack[0].Right)
		} else {
			slice = append(slice, 0)
		}
		stack = stack[1:]
	}

	return slice
}
