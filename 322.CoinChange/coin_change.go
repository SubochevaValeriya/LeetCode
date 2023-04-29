package main

import (
	"fmt"
	"sort"
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := range dp {
		dp[i] = amount + 1
	}

	for _, coin := range coins {
		if coin > amount {
			continue
		}
		dp[coin] = 1
		for i := coin + 1; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

//func coinChange(coins []int, amount int) int {
//	dp := make([]int, amount+1, amount+1)
//
//	dp[0] = 0
//
//	for i := 1; i < len(dp); i++ {
//		dp[i] = amount + 1
//
//		for j := 0; j < len(coins); j++ {
//			if i-coins[j] >= 0 && dp[i-coins[j]] != -1 {
//				dp[i] = min(dp[i], dp[i-coins[j]]+1)
//			}
//			if j == len(coins)-1 && dp[i] == amount+1 {
//				dp[i] = -1
//			}
//		}
//	}
//
//	return dp[amount]
//}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1, amount+1)

	dp[0] = 0

	for i := 1; i < len(dp); i++ {
		currMin := amount + 1
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && dp[i-coins[j]] != -1 {
				if dp[i-coins[j]]+1 < currMin {
					dp[i] = dp[i-coins[j]] + 1
					currMin = dp[i]
				}
			}
			if j == len(coins)-1 && currMin == amount+1 {
				dp[i] = -1
			}
		}
	}

	return dp[amount]
}

func coinChangeGreedy(coins []int, amount int) int {
	sort.Ints(coins)
	q := 0
	leftAmount := amount
	for i := len(coins) - 1; i >= 0; i-- {
		fmt.Println(i, coins, leftAmount, q)
		if leftAmount >= coins[i] {
			q += leftAmount / coins[i]
			leftAmount -= (leftAmount / coins[i]) * coins[i]
		}

		if leftAmount == 0 {
			return q
		}
	}

	return -1
}
