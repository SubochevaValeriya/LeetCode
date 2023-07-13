package main

import (
	"fmt"
	"math"
)

func main() {

	f := &TreeNode{
		Val:   -10,
		Left:  nil,
		Right: nil,
	}

	d := &TreeNode{
		Val:   -5,
		Left:  nil,
		Right: nil,
	}

	e := &TreeNode{
		Val:   -20,
		Left:  nil,
		Right: nil,
	}

	b := &TreeNode{
		Val:   -300,
		Left:  d,
		Right: e,
	}

	c := &TreeNode{
		Val:   -200,
		Left:  f,
		Right: nil,
	}

	a := &TreeNode{
		Val:   -100,
		Left:  b,
		Right: c,
	}

	fmt.Println(maxLevelSum(a))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {

	sumsByLevel := make([]int, 1)
	sumsByLevel, lastLevel := dfs(root, sumsByLevel, 0)

	maxSum := math.Inf(-1)
	maxLevel := -1
	for i, sum := range sumsByLevel[:lastLevel+1] {
		if float64(sum) > maxSum {
			maxSum = float64(sum)
			maxLevel = i + 1
		}
	}
	fmt.Println(sumsByLevel)
	return maxLevel
}

func dfs(root *TreeNode, sumsByLevel []int, level int) ([]int, int) {
	var lastLevelLeft, lastLevelRight, lastLevel int
	if root == nil {
		return sumsByLevel, level - 1
	}
	sumsByLevel[level] += root.Val

	if root.Left != nil || root.Right != nil {
		fmt.Println(root.Left, root.Right, root.Val, sumsByLevel)
		sumsByLevel = append(sumsByLevel, 0)
	}
	level += 1

	sumsByLevel, lastLevelLeft = dfs(root.Left, sumsByLevel, level)
	sumsByLevel, lastLevelRight = dfs(root.Right, sumsByLevel, level)

	if lastLevelLeft > lastLevelRight {
		lastLevel = lastLevelLeft
	} else {
		lastLevel = lastLevelRight
	}

	return sumsByLevel, lastLevel
}
