package main

func main() {

}

func tribonacci(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	if 1 <= n {
		dp[1] = 1
	}
	if 2 <= n {
		dp[2] = 1
	}

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	return dp[n]
}
