package main

import "fmt"

func main() {
	//obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	obstacleGrid := [][]int{{0, 1}}

	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(obstacleGrid[0])+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if obstacleGrid[i-1][j-1] == 1 {
				dp[i][j] = 0
			} else if i == 1 && j == 1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	fmt.Println(dp)

	return dp[len(dp)-1][len(dp[0])-1]
}
