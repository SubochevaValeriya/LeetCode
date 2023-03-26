package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
}

func rotate(matrix [][]int) {

	n := len(matrix)
	m := len(matrix[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m/2; j++ {
			matrix[i][j], matrix[i][m-j-1], matrix[n-1-i][m-j-1], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j], matrix[i][m-j-1], matrix[n-1-i][m-j-1]
		}
	}

	fmt.Println(matrix)
}
