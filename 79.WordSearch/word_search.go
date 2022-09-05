package main

func main() {

}

func exist(board [][]string, word string) bool {
	m := map[string][][]int{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			sl := []int{i, j}
			m[board[i][j]] = append(m[board[i][j]], sl)
		}
	}

	matrixUse := [][]bool{}
	var currRow int
	var currColumn int

	for i := 0; i < len(word); i++ {
		if _, ok := m[string(word[i])]; !ok {
			return false
		}
		variants := m[string(word[i])]
		for j := range variants {
			if i == 0 || abs(currRow-variants[j][0]) == 1 && matrixUse[variants[j][0]][variants[j][1]] != true || abs(currRow-variants[j][0]) == 1 {
				currRow, currColumn := variants[j][0], variants[j][1]
			} else {

			}
		}
	}

	return word
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
