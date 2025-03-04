package main

func main() {

}

func exist(board [][]byte, word string) bool {
	boardUsed := make([][]bool, len(board))
	for i := range boardUsed {
		boardUsed[i] = make([]bool, len(board[0]))
	}

	for i, bytes := range board {
		for j, letter := range bytes {
			if letter == word[0] {
				boardUsed[i][j] = true
				if isExist(word[1:], i, j, board, boardUsed) {
					return true
				}
				boardUsed[i][j] = false
			}
		}
	}

	return false
}

func isExist(word string, lastI, lastJ int, board [][]byte, boardUsed [][]bool) bool {
	if len(word) == 0 {
		return true
	}

	for i := lastI - 1; i <= lastI+1; i++ {
		if i < 0 || i >= len(board) {
			continue
		}
		if board[i][lastJ] == word[0] && !boardUsed[i][lastJ] {
			boardUsed[i][lastJ] = true
			if isExist(word[1:], i, lastJ, board, boardUsed) {
				return true
			}
			boardUsed[i][lastJ] = false
		}
	}

	for j := lastJ - 1; j <= lastJ+1; j++ {
		if j < 0 || j >= len(board[0]) {
			continue
		}
		if board[lastI][j] == word[0] && !boardUsed[lastI][j] {
			boardUsed[lastI][j] = true
			if isExist(word[1:], lastI, j, board, boardUsed) {
				return true
			}
			boardUsed[lastI][j] = false
		}
	}

	return false
}

//func exist1(board [][]string, word string) bool {
//	m := map[string][][]int{}
//	for i := 0; i < len(board); i++ {
//		for j := 0; j < len(board[i]); j++ {
//			sl := []int{i, j}
//			m[board[i][j]] = append(m[board[i][j]], sl)
//		}
//	}
//
//	matrixUse := [][]bool{}
//	var currRow int
//	var currColumn int
//
//	for i := 0; i < len(word); i++ {
//		if _, ok := m[string(word[i])]; !ok {
//			return false
//		}
//		variants := m[string(word[i])]
//		for j := range variants {
//			if i == 0 || abs(currRow-variants[j][0]) == 1 && matrixUse[variants[j][0]][variants[j][1]] != true || abs(currRow-variants[j][0]) == 1 {
//				currRow, currColumn := variants[j][0], variants[j][1]
//			} else {
//
//			}
//		}
//	}
//
//	return word
//}
//
//func abs(a int) int {
//	if a >= 0 {
//		return a
//	}
//	return -a
//}
