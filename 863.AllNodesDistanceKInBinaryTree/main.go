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

func dfs(neighbours map[*TreeNode][]*TreeNode, root *TreeNode, parent *TreeNode, target *TreeNode, rootFound bool, k int, nodesAtDistance []int, checked map[*TreeNode]bool) ([]int, map[*TreeNode][]*TreeNode, map[*TreeNode]bool) {
	if root == nil {
		return nodesAtDistance, neighbours, checked
	}

	_, ok := neighbours[root]
	if !ok {
		neighbours[root] = []*TreeNode{parent, root.Left, root.Right}
	}

	if !rootFound {
		if root == target {
			rootFound = true
		} else {
			nodesAtDistance, neighbours, checked = dfs(neighbours, root.Left, root, target, rootFound, k, nodesAtDistance, checked)
			nodesAtDistance, neighbours, checked = dfs(neighbours, root.Right, root, target, rootFound, k, nodesAtDistance, checked)
		}
	}

	if rootFound {
		checked[root] = true
		if k == 0 {
			nodesAtDistance = append(nodesAtDistance, root.Val)
			return nodesAtDistance, neighbours, checked
		}
		for _, neighbour := range neighbours[root] {
			if !checked[root] {
				nodesAtDistance, neighbours, checked = dfs(neighbours, neighbour, root, target, rootFound, k-1, nodesAtDistance, checked)
			}
		}
		return nodesAtDistance, neighbours, checked
	}
	return nodesAtDistance, neighbours, checked
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	nodesAtDistance, _, _ := dfs(map[*TreeNode][]*TreeNode{}, root, nil, target, false, k, []int{}, map[*TreeNode]bool{})
	return nodesAtDistance
}
