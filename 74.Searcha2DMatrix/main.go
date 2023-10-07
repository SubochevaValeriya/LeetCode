package main

import "fmt"

func main() {

	//	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	matrix := [][]int{{1, 3}}

	fmt.Println(searchMatrix(matrix, 1))

}

func searchMatrix(matrix [][]int, target int) bool {
	left, right := 0, len(matrix)-1

	for left <= right {
		medium := left + (right-left)/2
		if matrix[medium][len(matrix[0])-1] == target {
			return true
		}
		if (matrix[medium][len(matrix[0])-1] > target && matrix[medium][0] < target) || (matrix[medium][len(matrix[0])-1] > target && medium == 1) {
			leftRow, rightRow := 0, len(matrix[0])-1
			for leftRow <= rightRow {
				mediumRow := leftRow + (rightRow-leftRow)/2
				if matrix[medium][mediumRow] == target {
					return true
				}
				if matrix[medium][mediumRow] < target {
					leftRow++
				} else {
					rightRow--
				}
			}
			return false
		}

		if matrix[medium][len(matrix[0])-1] < target {
			left++
		} else {
			right--
		}
	}
	return false
}
