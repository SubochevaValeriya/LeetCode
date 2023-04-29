package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 4, 6, 0, 0, 3}
	fmt.Println(jump(nums))
}

func jump(nums []int) int {
	dp := make([]int, len(nums))
	dp[len(nums)-1] = 0

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] == 0 {
			dp[i] = -1
			continue
		}
		if i+nums[i] >= len(nums)-1 {
			dp[i] = 1
			continue
		}
		minJump := 100000
		for j := i + nums[i]; j > i; j-- {
			if dp[j] == -1 {
				continue
			}
			if dp[j] < minJump {
				minJump = dp[j]
			}
		}
		if minJump == -1 {
			dp[i] = -1
		} else {
			dp[i] = minJump + 1
		}
	}

	fmt.Println(dp)
	return dp[0]
}
