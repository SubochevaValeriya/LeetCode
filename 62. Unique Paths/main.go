package main

import "fmt"

func main() {
	m, n := 3, 2
	fmt.Println(uniquePaths(m, n))
	fmt.Println(currStep(m, n))
}

func uniquePaths(m int, n int) int {

	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				grid[i][j] = 1
			} else {
				grid[i][j] = grid[i-1][j] + grid[i][j-1]
			}
		}
	}

	fmt.Println(grid)
	return grid[m-1][n-1]
}

func currStep(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	return currStep(m-1, n) + currStep(m, n-1)
}
