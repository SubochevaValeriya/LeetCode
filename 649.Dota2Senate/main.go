package main

import "fmt"

func main() {
	fmt.Println(predictPartyVictory("DRRD"))
}

func predictPartyVictory(senate string) string {
	senateRune := []rune(senate)
	lastString := ""
	for i := 0; i < len(senateRune); i++ {
		fmt.Println(i, string(senateRune[i]), senateRune)
		if senateRune[i] == 'R' {
			j := i + 1
			if i == len(senateRune)-1 {
				j = 0
			}
			for ; j < len(senateRune); j++ {
				if senateRune[j] == 'D' {
					senateRune = append(senateRune[:j], senateRune[j+1:]...)
					break
				}
				if j == len(senateRune)-1 && i != j {
					j = -1
				}
				if i == j {
					return "Radiant"
				}
			}
			lastString = "Radiant"
		} else if senateRune[i] == 'D' {
			j := i + 1
			if i == len(senateRune)-1 {
				j = 0
			}
			for ; j < len(senateRune); j++ {
				if senateRune[j] == 'R' {
					//	fmt.Println(i, j, senateRune[j], senateRune[i])
					senateRune = append(senateRune[:j], senateRune[j+1:]...)
					break
				}
				if j == len(senateRune)-1 && i != j {
					j = -1
				}
				if i == j {
					return "Dire"
				}
			}
			lastString = "Dire"
		}

		if i > len(senateRune)-2 {
			i = -1
		}
	}
	fmt.Println(string(senateRune))
	return lastString
}
