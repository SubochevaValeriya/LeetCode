package main

import "fmt"

func main() {
	nums := []int{13, 13, 20, 0, 8, 9, 9}
	fmt.Println(minimizeArrayValue(nums))
}

func minimizeArrayValue(nums []int) int {

	dp := make([]int, len(nums))

	dp[0] = nums[0]
	prefixSum := nums[0]
	prefixN := 1
	for i := 1; i < len(nums); i++ {
		prefixSum += nums[i]
		prefixN++
		newMax := prefixSum / prefixN
		if prefixSum%prefixN != 0 {
			newMax++
		}
		dp[i] = max(newMax, dp[i-1])
	}
	fmt.Println(dp)
	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
