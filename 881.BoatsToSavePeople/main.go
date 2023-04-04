package main

import "sort"

func main() {

}

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	left, right := 0, len(people)-1
	boats := 0

	for left <= right {
		if people[right] == limit {
			boats++
			right--
			continue
		}
		if people[right]+people[left] <= limit {
			boats++
			right--
			left++
		} else {
			boats++
			right--
		}
	}

	return boats
}
