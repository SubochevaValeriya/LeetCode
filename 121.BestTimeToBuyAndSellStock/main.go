package main

import "fmt"

func main() {
	prices := []int{
		7, 1, 5, 3, 6, 4,
		//2,1,2,1,0,1,2
	}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {

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

//func maxProfit(prices []int) int {
//	type p struct {
//		a, b, profit int
//	}
//	var maxProfit, profit int
//
//	var dict p
//
//	i := 0
//	j := len(prices) - 1
//
//	for i != len(prices)-1 {
//
//		if j == i {
//			i++
//			if i == len(prices)-1 {
//				break
//			}
//			j = len(prices) - 1
//		}
//
//		if prices[j] < prices[i] {
//			j--
//			continue
//		}
//		val1 := prices[i]
//		val2 := prices[j]
//
//		if dict.a == val1 && dict.b == val2 {
//			profit = dict.profit
//		} else {
//			profit = val2 - val1
//		}
//
//		if profit > maxProfit {
//			maxProfit = profit
//		}
//		j--
//	}
//	return maxProfit
//}

//func maxProfit(prices []int) int {
//	var profit int
//
//	dict := map[int]int{}
//
//	for i, val1 := range prices {
//		for _, val2 := range prices[i+1:] {
//			if val2-val1 > profit {
//				profit = val2 - val1
//			}
//		}
//	}
//	return profit
//}

//func maxProfit(prices []int) int {
//	type p struct {
//		a, b, profit int
//	}
//	var maxProfit, profit int
//
//	dict := make([]p, 1, len(prices))
//
//	for i, val1 := range prices {
//		for _, val2 := range prices[i+1:] {
//			if val2 <= val1 {
//				continue
//			}
//			for i := range dict {
//				if dict[i].a == val1 && dict[i].b == val2 {
//					profit = dict[i].profit
//				} else {
//					profit = val2 - val1
//					item := p{val1, val2, profit}
//					dict = append(dict, item)
//				}
//			}
//
//			if profit > maxProfit {
//				maxProfit = profit
//			}
//		}
//	}
//	return maxProfit
//}
