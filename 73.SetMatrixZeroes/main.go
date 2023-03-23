package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{{0, 0, 0, 5}, {4, 3, 1, 4}, {0, 1, 1, 4}, {1, 2, 1, 3}, {0, 0, 1, 1}}
	for _, ints := range matrix {
		fmt.Println(ints)
	}
	fmt.Println()

	setZeroes(matrix)
}

func setZeroes(matrix [][]int) {

	mSettedRow := map[int]struct{}{}
	mSettedColumn := map[int]struct{}{}
	mNotZeroBefore := map[[2]int]struct{}{}

	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			if matrix[i][j] == 0 {
				_, ok := mNotZeroBefore[[2]int{i, j}]
				if ok {
					continue
				}
				_, settedRow := mSettedRow[i]
				_, settedColumn := mSettedColumn[j]
				if !settedRow {
					matrix, mNotZeroBefore = setZeroesRow(matrix, i, mNotZeroBefore)
					mSettedRow[i] = struct{}{}
				}

				if !settedColumn {
					matrix, mNotZeroBefore = setZeroesColumn(matrix, j, mNotZeroBefore)
					mSettedColumn[j] = struct{}{}
				}
			}
		}
	}

	for _, ints := range matrix {
		fmt.Println(ints)
	}
}

func setZeroesRow(matrix [][]int, row int, notZeroBefore map[[2]int]struct{}) ([][]int, map[[2]int]struct{}) {
	for i, _ := range matrix[row] {
		if matrix[row][i] != 0 {
			notZeroBefore[[2]int{row, i}] = struct{}{}
		}
		matrix[row][i] = 0
	}
	return matrix, notZeroBefore
}

func setZeroesColumn(matrix [][]int, column int, notZeroBefore map[[2]int]struct{}) ([][]int, map[[2]int]struct{}) {
	for i, _ := range matrix {
		if matrix[i][column] != 0 {
			notZeroBefore[[2]int{i, column}] = struct{}{}
		}
		matrix[i][column] = 0
	}
	return matrix, notZeroBefore
}
