package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	dp := make([]bool, len(nums))

	dp[len(nums)-1] = true

	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= len(nums)-1 {
			dp[i] = true
			continue
		}
		for j := i + nums[i]; j > i; j-- {
			if dp[j] == true {
				dp[i] = true
				break
			}
		}
	}
	return dp[0]
}
