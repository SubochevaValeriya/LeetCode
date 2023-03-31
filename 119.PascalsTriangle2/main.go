package main

import "fmt"

func main() {
	fmt.Println(getRow(3))
}

func getRow(rowIndex int) []int {
	dp := make([][]int, rowIndex+1)

	for i := 0; i < rowIndex+1; i++ {
		dp[i] = make([]int, i+1)
	}

	for i := 0; i < rowIndex+1; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			}
		}
	}

	return dp[rowIndex]
}
