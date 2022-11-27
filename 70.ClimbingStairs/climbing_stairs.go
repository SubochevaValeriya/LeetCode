package main

import "fmt"

func main() {
	n := 5
	fmt.Println(climbStairs(n))
}

func climbStairs(n int) int {
	one, two := 1, 1

	for i := 1; i < n; i++ {
		temp := one
		one = one + two
		two = temp
	}

	return one
	//
	//dp := make([]int, n+1, n+1)
	//
	//dp[1] = n1
	//dp[2] = n2
	//
	//for i := 3; i <= n; i++ {
	//	dp[i] = dp[i-1] + 1 + dp[i-2] + 1
	//}
	//return dp[n]
}
