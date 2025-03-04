package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findTheWinner(5, 2))

}

func findTheWinner(n int, k int) int {
	players := make([]int, n)

	for i := range players {
		players[i] = i + 1
	}

	for start := 0; len(players) != 1; {
		nextLoserIn := int(math.Mod(float64(k), float64(len(players))))
		loser := start + nextLoserIn - 1

		if loser > len(players)-1 {
			loser = loser - len(players)
		}
		if loser < 0 {
			loser = len(players) + loser
		}

		if loser+1 > len(players)-1 {
			players = players[:loser]
		} else {
			players = append(players[:loser], players[loser+1:]...)
		}

		start = loser
	}

	return players[0]
}
