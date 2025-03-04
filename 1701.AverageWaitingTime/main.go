package main

import "fmt"

func main() {
	customers := [][]int{{1, 2}, {2, 5}, {4, 3}}
	fmt.Println(averageWaitingTime(customers))

}

func averageWaitingTime(customers [][]int) float64 {
	var timeWhenChefWillBeIdle, totalWaitingTime int
	for _, customer := range customers {
		arrival := customer[0]
		timeToPrepareOrder := customer[1]
		if arrival > timeWhenChefWillBeIdle {
			timeWhenChefWillBeIdle = arrival + timeToPrepareOrder
		} else {
			timeWhenChefWillBeIdle += timeToPrepareOrder
		}
		totalWaitingTime = totalWaitingTime + timeWhenChefWillBeIdle - arrival
	}

	return float64(totalWaitingTime) / float64(len(customers))
}
