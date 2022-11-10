package main

import "fmt"

func main() {

	matrix := [][]int{{1, 2, 3}, {3, 1, 2}, {2, 3, 1}}
	fmt.Println(checkValid(matrix))

}

func checkValid(matrix [][]int) bool {
	for i := 1; i <= len(matrix); i++ {

	}

	columnMaps := map[int]map[int]int{}

	for i := 0; i < len(matrix); i++ {
		rowMap := map[int]int{}
		for j := 0; j < len(matrix[i]); j++ {
			integer := matrix[i][j]

			rowMap[integer]++
			if rowMap[integer] > 1 {
				return false
			}
			if columnMaps[j] == nil {
				columnMaps[j] = map[int]int{}
			}
			columnMaps[j][integer]++
			fmt.Println(columnMaps)
			if columnMaps[j][integer] > 1 {
				return false
			}
		}
	}

	return true
}
