package main

import "fmt"

func main() {

	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}

	fmt.Println(minCostClimbingStairs(cost))

}

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)

	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < len(dp); i++ {
		if i == len(cost) {
			dp[i] = min(dp[i-1], dp[i-2])
			continue
		}
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	fmt.Println(dp)

	return dp[len(dp)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
