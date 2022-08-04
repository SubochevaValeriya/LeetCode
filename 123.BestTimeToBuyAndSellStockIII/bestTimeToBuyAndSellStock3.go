package main

import "fmt"

func maxProfit(prices []int) int {

	var bestBuyingPrice, profit, maxProfit int
	buyingDay1 := len(prices) - 1

	bestBuyingPrice = prices[0]

	for i, val := range prices {
		profit = 0
		if val < bestBuyingPrice {
			bestBuyingPrice = val
			buyingDay1 = i
		}
		profit = val - bestBuyingPrice

		if buyingDay1 < i && i != len(prices)-1 {
			profit = profit + maxProfitR(prices[i+1:])
			fmt.Println(profit, prices)
		}
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

func maxProfitR(prices []int) int {

	var bestBuyingPrice, profit, maxProfit int

	bestBuyingPrice = prices[0]

	for _, val := range prices {

		if val < bestBuyingPrice {
			bestBuyingPrice = val
		}

		profit = val - bestBuyingPrice

		if profit > maxProfit {
			maxProfit = profit
		}

	}

	return maxProfit
}

//func maxProfit2(prices []int) int {
//
//	var bestBuyingPrice, bestBuyingPrice2, profit1, profit2, maxProfit int
//
//	bestBuyingPrice = prices[0]
//
//	for i, val := range prices {
//
//		if val < bestBuyingPrice {
//			bestBuyingPrice = val
//		}
//
//		profit1 = val - bestBuyingPrice
//
//		if i != len(prices)-1 {
//			bestBuyingPrice2 = prices[i+1]
//			for _, val2 := range prices[i+1:] {
//				if val2 < bestBuyingPrice2 {
//					bestBuyingPrice2 = val2
//				}
//				fmt.Println(val2, prices, val)
//				profit2 = val2 - bestBuyingPrice2
//				if profit1+profit2 > maxProfit {
//					maxProfit = profit1 + profit2
//				}
//			}
//		} else {
//			if profit1 > maxProfit {
//				maxProfit = profit1
//			}
//		}
//	}
//
//	return maxProfit
//}
