package main

import (
	"fmt"
	"sort"
)

func main() {
	reward1 := []int{1, 1, 3, 4}
	reward2 := []int{4, 4, 1, 1}
	k := 2
	fmt.Println(miceAndCheese(reward1, reward2, k))
}

type rewardsStruct struct {
	diffs   int
	reward1 int
	reward2 int
}

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	rewards := make([]rewardsStruct, len(reward1))
	for i := range reward1 {
		rewards[i].diffs = reward1[i] - reward2[i]
		rewards[i].reward1 = reward1[i]
		rewards[i].reward2 = reward2[i]
	}

	sort.Slice(rewards, func(i, j int) bool {
		if rewards[i].diffs >= rewards[j].diffs {
			return true
		}
		return false
	})

	points := 0

	for _, reward := range rewards {
		if k != 0 {
			points += reward.reward1
			k--
		} else {
			points += reward.reward2
		}
	}

	return points
}
