package main

import "fmt"

func main() {

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
	//matrix3 := make([][]int, 3, 3)
	matrix2 := make([][]int, len(matrix[0]))
	for i := range matrix2 {
		matrix2[i] = make([]int, len(matrix))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix2[j][i] = matrix[i][j]
			fmt.Println(matrix2)
		}
		//	fmt.Println(matrix)
	}
	fmt.Println(len(matrix))
	fmt.Println(matrix2)
}

func transpose(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Println(matrix[i][j])
		}
	}
	return matrix
}
