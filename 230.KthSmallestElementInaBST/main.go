package main

import (
	"fmt"
)

func main() {
	n := generateTreeNode([]int{5, 3, 6, 2, 4, 0, 0, 1})
	fmt.Println(kthSmallest(n, 3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	arr := []int{}
	arr = inorder(root, arr)
	return arr[k-1]
	//min := []int{root.Val}
	//
	//min = add(root.Left, min)
	//if len(min) >= k {
	//	sort.Ints(min)
	//	return min[k-1]
	//}
	//
	//min = add(root.Right, min)
	//
	//sort.Ints(min)
	//return min[k-1]
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

func add(root *TreeNode, min []int) []int {
	min = append(min, root.Val)
	if root.Left != nil {
		add(root.Left, min)
	}
	if root.Right != nil {
		add(root.Right, min)
	}
	return min
}

func inorder(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}
	arr = inorder(root.Left, arr)
	arr = append(arr, root.Val)
	arr = inorder(root.Right, arr)
	return arr
}
