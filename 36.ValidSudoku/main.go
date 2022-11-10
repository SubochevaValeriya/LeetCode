package main

import (
	"fmt"
)

func main() {
	matrix := [][]string{
		{"5","3",".",".","7",".",".",".","."}
		,{"6",".",".","1","9","5",".",".","."}
		,{".","9","8",".",".",".",".","6","."}
		,{"8",".",".",".","6",".",".",".","3"}
		,{"4",".",".","8",".","3",".",".","1"}
		,{"7",".",".",".","2",".",".",".","6"}
		,{".","6",".",".",".",".","2","8","."}
		,{".",".",".","4","1","9",".",".","5"}
		,{".",".",".",".","8",".",".","7","9"}}
	fmt.Println(isValidSudoku(matrix))
}

func isValidSudoku(board [][]byte) bool {
	columnMaps := map[int]map[int]int{}
	boxesMap := map[int]map[int]map[int]int{}
	var integer, boxI, boxJ int
	for i := 0; i < len(board); i++ {
		rowMap := map[int]int{}
		switch {
		case i <= 2:
			boxI = 1
		case i > 5:
			boxI = 3
		default:
			boxI = 2
		}
		for j := 0; j < len(board[i]); j++ {
			if string(board[i][j]) == "." {
				continue
			} else {
				integer = int(board[i][j])
			}

			switch {
			case j <= 2:
				boxJ = 1
			case j > 5:
				boxJ = 3
			default:
				boxJ = 2
			}

			// Each row must contain the digits 1-9 without repetition.
			rowMap[integer]++
			if rowMap[integer] > 1 {
				return false
			}

			// Each column must contain the digits 1-9 without repetition.
			if columnMaps[j] == nil {
				columnMaps[j] = map[int]int{}
			}
			columnMaps[j][integer]++
			if columnMaps[j][integer] > 1 {
				return false
			}

			// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

			if boxesMap[boxI] == nil {
				boxesMap[boxI] = map[int]map[int]int{}
			}

			if boxesMap[boxI][boxJ] == nil {
				boxesMap[boxI][boxJ] = map[int]int{}
			}

			boxesMap[boxI][boxJ][integer]++
			if boxesMap[boxI][boxJ][integer] > 1 {
				return false
			}
		}
	}

	return true
}