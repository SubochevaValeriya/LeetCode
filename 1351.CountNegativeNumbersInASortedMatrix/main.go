package main

import "fmt"

func main() {

	grid := [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}
	fmt.Println(countNegatives(grid))
}

func countNegatives(grid [][]int) int {

	count := 0
	for i := len(grid) - 1; i >= 0; i-- {
		right := len(grid[i]) - 1
		if grid[i][right] > 0 {
			return count
		}
		for left := 0; left < len(grid[i]); left++ {
			if grid[i][left] < 0 {
				count += right - left + 1
				break
			}
		}
	}

	return count
}
