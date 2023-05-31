package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	values := []int{}

	listOfNodes := []*TreeNode{}

	for {

		if len(listOfNodes) <= 0 && root == nil {
			break
		}
		for root != nil {
			listOfNodes = append(listOfNodes, root)
			root = root.Left
		}
		if len(listOfNodes) > 0 {
			root = listOfNodes[len(listOfNodes)-1]
			values = append(values, root.Val)
			listOfNodes = listOfNodes[:len(listOfNodes)-1]
			root = root.Right
		}

	}
	return values
}
