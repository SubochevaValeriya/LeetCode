package main

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbb"))
}
func numJewelsInStones(jewels string, stones string) int {
	quantityOfJewels := 0
	jewelsList := map[rune]bool{}
	for _, jewelry := range jewels {
		jewelsList[jewelry] = true
	}

	for _, stone := range stones {
		if jewelsList[stone] {
			quantityOfJewels++
		}
	}

	return quantityOfJewels
}
