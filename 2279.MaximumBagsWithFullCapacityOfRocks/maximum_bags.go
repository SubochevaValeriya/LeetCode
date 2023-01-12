package main

import (
	"fmt"
	"sort"
)

func main() {
	capacity := []int{54, 18, 91, 49, 51, 45, 58, 54, 47, 91, 90, 20, 85, 20, 90, 49, 10, 84, 59, 29, 40, 9, 100, 1, 64, 71, 30, 46, 91}
	rocks := []int{14, 13, 16, 44, 8, 20, 51, 15, 46, 76, 51, 20, 77, 13, 14, 35, 6, 34, 34, 13, 3, 8, 1, 1, 61, 5, 2, 15, 18}
	additionalRocks := 77

	//capacity := []int{10, 1000000000}
	//rocks := []int{10, 0}
	//additionalRocks := 1000000000

	fmt.Println(maximumBags(capacity, rocks, additionalRocks))

}

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {

	differences := make([]int, len(capacity))
	for i := range capacity {
		needRocks := capacity[i] - rocks[i]
		differences[i] = needRocks
	}

	sort.Ints(differences)

	maxBags := 0

	fmt.Println(differences)
	for _, needRocks := range differences {

		if additionalRocks-needRocks >= 0 {
			additionalRocks -= needRocks
			maxBags += 1
		}

		if additionalRocks <= 0 {
			break
		}

	}

	return maxBags
}
