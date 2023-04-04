package main

import (
	"fmt"
	"sort"
)

func main() {
	reward1 := []int{1, 1, 3, 4}
	reward2 := []int{4, 4, 1, 1}
	k := 2
	miceAndCheese(reward1, reward2, k)
}

type rewards struct {
	reward1 []int
	reward2 []int
}

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	sort.Slice(reward1, func(i, j int) bool {
		return reward1[0] > reward2[0]
	})
	fmt.Println(reward1)
	fmt.Println(reward2)
	return 0
}
