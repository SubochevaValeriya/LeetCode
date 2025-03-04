package main

import "sort"

func main() {

}

func findWinners(matches [][]int) [][]int {
	winners := map[int]struct{}{}
	losers := map[int]int{}

	for _, match := range matches {
		winners[match[0]] = struct{}{}
		losers[match[1]]++
	}

	winnersList := []int{}
	onlyOneLostList := []int{}

	for winner := range winners {
		if losers[winner] == 0 {
			winnersList = append(winnersList, winner)
		}
	}

	for loser, matchCount := range losers {
		if matchCount == 1 {
			onlyOneLostList = append(onlyOneLostList, loser)
		}
	}

	sort.Ints(winnersList)
	sort.Ints(onlyOneLostList)

	return [][]int{winnersList, onlyOneLostList}
}
