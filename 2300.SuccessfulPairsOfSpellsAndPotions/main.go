package main

import (
	"fmt"
	"sort"
)

func main() {

	//spells := []int{40, 11, 24, 28, 40, 22, 26, 38, 28, 10, 31, 16, 10, 37, 13, 21, 9, 22, 21, 18, 34, 2, 40, 40, 6, 16, 9, 14, 14, 15, 37, 15, 32, 4, 27, 20, 24, 12, 26, 39, 32, 39, 20, 19, 22, 33, 2, 22, 9, 18, 12, 5}
	//potions := []int{31, 40, 29, 19, 27, 16, 25, 8, 33, 25, 36, 21, 7, 27, 40, 24, 18, 26, 32, 25, 22, 21, 38, 22, 37, 34, 15, 36, 21, 22, 37, 14, 31, 20, 36, 27, 28, 32, 21, 26, 33, 37, 27, 39, 19, 36, 20, 23, 25, 39, 40}
	//success := int64(600)

	spells := []int{3, 1, 2}
	potions := []int{8, 5, 8}
	success := int64(16)

	fmt.Println(successfulPairs(spells, potions, success))
}

func successfulPairs(spells []int, potions []int, success int64) []int {

	answer := make([]int, len(spells))
	allSuccess := 10000000000
	sort.Ints(potions)

	for i, spell := range spells {
		var minPotion int
		if int(success)%spell != 0 {
			minPotion = int(success)/spell + 1
		} else {
			minPotion = int(success) / spell
		}
		if potions[len(potions)-1] < minPotion {
			answer[i] = 0
			continue
		}
		if spell >= allSuccess {
			answer[i] = len(potions)
			continue
		}
		answer[i] = len(potions) - sort.SearchInts(potions, minPotion)

		if answer[i] == len(potions) {
			allSuccess = spell
		}
	}

	return answer
}

func successfulPairs2(spells []int, potions []int, success int64) []int {

	answer := make([]int, len(spells))
	allSuccess := 100000
	allNotSuccess := 0
	for i, spell := range spells {
		fmt.Println(answer)
		if spell > allSuccess {
			answer[i] = len(potions)
			continue
		}
		if spell < allNotSuccess {
			answer[i] = 0
			continue
		}

		for _, potion := range potions {
			if spell*potion >= int(success) {
				answer[i]++
			}
		}
		if answer[i] == 0 {
			allNotSuccess = spell
		}
		if answer[i] == len(potions) {
			allSuccess = spell
		}
	}

	return answer
}
