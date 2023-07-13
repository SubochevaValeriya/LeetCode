package main

import (
	"fmt"
	"sort"
)

func main() {

	d := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	e := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	b := &TreeNode{
		Val:   2,
		Left:  d,
		Right: e,
	}

	c := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}

	a := &TreeNode{
		Val:   4,
		Left:  b,
		Right: c,
	}

	fmt.Println(getMinimumDifference(a))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {

	values := dfs(root, []int{})

	sort.Ints(values)

	minDifference := 10000000
	for i := 1; i < len(values); i++ {
		if values[i]-values[i-1] < minDifference {
			minDifference = values[i] - values[i-1]
		}
	}
	fmt.Println(values)

	return minDifference
}

func dfs(root *TreeNode, values []int) []int {
	if root == nil {
		return values
	}
	values = append(values, root.Val)

	valuesLeft := dfs(root.Left, []int{})
	valuesRight := dfs(root.Right, []int{})

	values = append(values, valuesLeft...)
	values = append(values, valuesRight...)
	return values
}
