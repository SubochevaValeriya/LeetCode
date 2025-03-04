package main

import "fmt"

func main() {
	rowSum := []int{3, 8}
	colSum := []int{4, 7}

	fmt.Println(restoreMatrix(rowSum, colSum))
}

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	matrix := make([][]int, len(rowSum))
	for i := range matrix {
		matrix[i] = make([]int, len(colSum))
	}
	for i, ints := range matrix {
		for j := range ints {
			matrix[i][j] = min(rowSum[i], colSum[j])
			rowSum[i] -= matrix[i][j]
			colSum[j] -= matrix[i][j]
		}
	}
	fmt.Println(rowSum, colSum)
	return matrix
}
