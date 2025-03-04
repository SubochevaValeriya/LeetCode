package main

import "fmt"

func candy(ratings []int) int {
	var sum int
	candies := make([]int, len(ratings))

	for i := 0; i < len(ratings); i++ {
		candies[i] = 1
		if i == 0 {
			continue
		}

		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := len(ratings) - 1; i >= 0; i-- {
		if i == len(ratings)-1 {
			sum += candies[i]
			continue
		}

		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}

		sum += candies[i]
	}
	fmt.Println(candies)

	return sum
}

//
//for i := 0; i < len(ratings); i++ {
//	if ratings[i] > ratings[i+1] {
//		candies++
//	}
//
//	if ratings[i-1] > ratings[i] {
//		candies++
//	}
//
//	m[]
//
//	if i == 0 {
//		if ratings[i] > ratings[i+1]
//	}
//
//	if ratings[i] > ratings[i-1] {
//		candies += 2
//	} else {
//		candies += 1
//	}
//}
//return candies + 1
//}
