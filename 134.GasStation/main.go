package main

import (
	"fmt"
)

func main() {
	//gas := []int{1, 2, 3, 4, 5}
	//cost := []int{3, 4, 5, 1, 2}

	//gas := []int{3, 1, 1}
	//cost := []int{1, 2, 2}

	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}

	fmt.Println(canCompleteCircuit(gas, cost))
}

func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	considered := map[int]struct{}{0: {}}
	gasTank := 0
	for i := 0; i < len(gas); i++ {
		gasTank += gas[i]
		if gasTank-cost[i] < 0 {
			start = i + 1
			if _, ok := considered[start]; ok {
				return -1
			}
			considered[start] = struct{}{}
			gasTank = 0
			//i = start - 1

			continue
		}

		gasTank -= cost[i]
		if i == start-1 {
			return start
		}

		if i == len(gas)-1 {
			if start == 0 {
				return start
			}

			i = -1
		}
	}

	return -1
}
