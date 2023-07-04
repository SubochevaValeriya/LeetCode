package main

import "fmt"

func main() {

	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	fmt.Println(maxProfit(prices, fee))

}

func maxProfit(prices []int, fee int) int {
	hold := make([]int, len(prices))
	free := make([]int, len(prices))

	hold[0] = -prices[0]
	free[0] = 0
	for i := 1; i < len(prices); i++ {
		hold[i] = max(hold[i-1], free[i-1]-prices[i])
		free[i] = max(free[i-1], hold[i-1]+prices[i]-fee)
	}
	return free[len(free)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
