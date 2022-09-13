package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	answer := []float64{}
	if root == nil {
		return answer
	}
	sum, count := []int{}, []int{}

	sum, count = byLevels(root, 0, sum, count)

	for i := 0; i < len(sum); i++ {
		answer = append(answer, float64(sum[i])/float64(count[i]))
	}
	return answer
}

func byLevels(root *TreeNode, i int, sum, count []int) ([]int, []int) {
	if root == nil {
		return sum, count
	}
	if len(sum)-1 >= i {
		sum[i] += root.Val
		count[i] += 1
	} else {
		sum = append(sum, root.Val)
		count = append(count, 1)
	}

	sum, count = byLevels(root.Left, i+1, sum, count)
	return byLevels(root.Right, i+1, sum, count)

}
