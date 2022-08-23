package main

import (
	"fmt"
	"sort"
)

func coinChange(coins []int, amount int) int {
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
