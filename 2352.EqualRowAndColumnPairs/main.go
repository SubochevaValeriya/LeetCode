package main

import "fmt"

func main() {
	grid := [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}
	fmt.Println(equalPairs(grid))
}

func equalPairs(grid [][]int) int {
	var count int
	j := 0
	n := 0
	for i := 0; i < len(grid); i++ {
		for m := 0; m < len(grid); m++ {
			j = 0
			n = 0
			for grid[i][j] == grid[n][m] {
				if j == len(grid)-1 {
					count++
					break
				}
				j++
				n++
			}
		}
	}

	return count
}
