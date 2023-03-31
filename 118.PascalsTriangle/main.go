package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {

	dp := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		dp[i] = make([]int, i+1)
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			fmt.Println(i, j)
			if j == 0 || j == i {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			}
		}
	}

	return dp
}
