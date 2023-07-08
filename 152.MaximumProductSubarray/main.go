package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{0, -2, 0}))
	fmt.Println(maxProduct([]int{2, -5, -2, -4, 3}))
	fmt.Println(maxProduct([]int{6, 3, -10, 0, 2}))
	fmt.Println(maxProduct([]int{0, -2, -3}))
	fmt.Println(maxProduct([]int{-3, -1, -1}))
	fmt.Println(maxProduct([]int{-1, 0, -1, 2, -3, 1, 2, 3, -2}))

	//fmt.Println(maxProduct([]int{0, -2, 0}))
}

func maxProduct(nums []int) int {

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	prod, prodWtFirstNegative := nums[0], nums[0]
	negative, firstNegative := 0, -1
	if nums[0] < 0 {
		negative++
		firstNegative = 0
		prodWtFirstNegative = 1
	}
	for i := 1; i < len(nums); i++ {
		prod *= nums[i]
		if nums[i-1] == 0 {
			prod = nums[i]
		}
		dp[i] = max(prod, dp[i-1])
		dp[i] = max(nums[i], dp[i])

		if nums[i] < 0 {
			negative++
			if negative == 1 {
				prodWtFirstNegative = 1
				firstNegative = i
			}
		}

		if firstNegative != i {
			prodWtFirstNegative *= nums[i]
			dp[i] = max(dp[i], prodWtFirstNegative)
		}

		if nums[i] == 0 && i != len(nums)-1 {
			maxNextSubarray := maxProduct(nums[i+1:])
			return max(dp[i], maxNextSubarray)
		}
	}

	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
