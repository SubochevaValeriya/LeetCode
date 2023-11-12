package main

import "fmt"

func main() {
	n := 5
	leftChild := []int{0, -1, 3, 1, 3}
	rightChild := []int{4, 3, 0, 1, -1}
	fmt.Println(validateBinaryTreeNodes(n, leftChild, rightChild))
}

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	parents := map[int]int{}
	var valid bool
	for i := 0; i < n; i++ {
		parents, valid = addParent(leftChild[i], i, parents)
		if !valid {
			return false
		}
		parents, valid = addParent(rightChild[i], i, parents)
		if !valid {
			return false
		}
	}

	rootCount := 0
	for i := 0; i < n; i++ {
		if _, ok := parents[i]; !ok {
			rootCount++
			if rootCount > 1 {
				return false
			}
		}
	}

	return true
}

func addParent(child int, parent int, parents map[int]int) (map[int]int, bool) {
	if child == -1 {
		return parents, true
	}

	if child == parent {
		return parents, false
	}

	if _, ok := parents[child]; ok {
		return parents, false
	}

	if childIsParent(child, parent, parents) {
		return parents, false
	}

	parents[child] = parent

	return parents, true
}

func childIsParent(child int, parent int, parents map[int]int) bool {
	fmt.Println(child, parent, parents)
	if _, ok := parents[parent]; !ok {
		return false
	}

	if child == parents[parent] {
		return true
	}

	return childIsParent(child, parents[parent], parents)
}
