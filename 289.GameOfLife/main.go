package main

func main() {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board)
}

func gameOfLife(board [][]int) {
	newState := make([][]int, len(board))
	for i := 0; i < len(newState); i++ {
		newState[i] = make([]int, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			var countLive int
			if i != 0 {
				countLive += board[i-1][j]
				if j != 0 {
					countLive += board[i-1][j-1]
				}
				if j != len(board[0])-1 {
					countLive += board[i-1][j+1]
				}
			}

			if j != 0 {
				countLive += board[i][j-1]
			}

			if i != len(board)-1 {
				countLive += board[i+1][j]
				if j != 0 {
					countLive += board[i+1][j-1]
				}
				if j != len(board[0])-1 {
					countLive += board[i+1][j+1]
				}
			}

			if j != len(board[0])-1 {
				countLive += board[i][j+1]
			}

			switch board[i][j] {
			case 1:
				if countLive == 2 || countLive == 3 {
					newState[i][j] = 1
				}
			case 0:
				if countLive == 3 {
					newState[i][j] = 1
				}
			}
		}
	}
	copy(board, newState)
}
