package main

func candy(ratings []int) int {
	var candies int

	candies = len(ratings)

	for i := 0; i < len(ratings); i++ {
		if ratings[i] > ratings[i+1] {
			candies++
		}

		if ratings[i-1] > ratings[i] {
			candies++
		}

		m[]

		if i == 0 {
			if ratings[i] > ratings[i+1]
		}

		if ratings[i] > ratings[i-1] {
			candies += 2
		} else {
			candies += 1
		}
	}
	return candies + 1
}
