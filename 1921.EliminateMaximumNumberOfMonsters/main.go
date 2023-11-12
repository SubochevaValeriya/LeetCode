package main

import (
	"fmt"
	"slices"
)

func main() {

	dist := []int{3, 5, 7, 4, 5}
	speed := []int{2, 3, 6, 3, 2}

	fmt.Println(eliminateMaximum(dist, speed))
}

func eliminateMaximum(dist []int, speed []int) int {
	timesToArrive := make([]float64, len(dist))

	for i := 0; i < len(timesToArrive); i++ {
		timesToArrive[i] = float64(dist[i]) / float64(speed[i])
	}

	slices.Sort(timesToArrive)

	for i, time := range timesToArrive {
		if i != 0 && time <= float64(i) {
			return i
		}
	}

	return len(dist)
}
