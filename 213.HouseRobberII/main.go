package main

import "fmt"

func main() {
	nums := []int{2, 3, 2}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))

	if nums[len(nums)-1] >= nums[0] {
		dp[0] = 0
	} else {
		dp[0] = nums[0]
	}

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return max(dp[len(dp)-1], dp[len(dp)-2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}