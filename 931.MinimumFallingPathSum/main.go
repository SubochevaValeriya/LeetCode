package main

func main() {

}

func minFallingPathSum(matrix [][]int) int {
	var minFallingPath int
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := 0; j < len(matrix[i]); j++ {
			if i != len(matrix)-1 {
				var minimum int
				if j == 0 {
					minimum = min(matrix[i+1][j], matrix[i+1][j+1])
				} else if j == len(matrix[i])-1 {
					minimum = min(matrix[i+1][j], matrix[i+1][j-1])
				} else {
					minimum = min(min(matrix[i+1][j], matrix[i+1][j+1]), matrix[i+1][j-1])
				}
				matrix[i][j] += minimum
			}
			if i == 0 {
				if j == 0 {
					minFallingPath = matrix[i][j]
				}
				if matrix[i][j] < minFallingPath {
					minFallingPath = matrix[i][j]
				}
			}
		}
	}

	return minFallingPath
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
