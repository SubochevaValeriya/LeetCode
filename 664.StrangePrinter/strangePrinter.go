package main

import (
	"fmt"
)

func main() {
	//fmt.Println(strangePrinter("tbgtgb"))
	//fmt.Println(strangePrinter("abcabc"))
	//fmt.Println(strangePrinter("aaabb"))
	fmt.Println(strangePrinter("abcabcabc"))
}

func strangePrinter(s string) int {

	sequences := []rune{}

	dp := make([]int, len(s))

	dp[0] = 1
	strangePrinterStep(0, s, sequences, dp, 0)
	fmt.Println(dp, "FF")
	return dp[len(s)-1]
}

func strangePrinterStep(startOfNewString int, s string, sequences []rune, dp []int, minTurns int) int {
	//fmt.Println(dp)
	//minTurns := 0
	for i, letter := range s {
		//	fmt.Println(string(letter), dp)
		//	fmt.Println(string(letter), startOfNewString, i, dp)
		if dp[startOfNewString+i] != 0 {
			continue
		}
		if i > 0 {
			if s[i] == s[i-1] {
				dp[startOfNewString+i] = dp[startOfNewString+i-1]
				continue
			}
		} else {
			minTurns++
		}
		//	fmt.Printf("ВНИМАНИЕ %d, %s, %d", i, string(letter), sequences)
		for j := 0; j < len(sequences); j++ {
			if sequences[j] == letter {

				sNew := make([]rune, len(sequences))
				copy(sNew, sequences)
				s1 := sNew[j+1:]
				s2 := sNew[:j+1]

				printNewCharacter := strangePrinterStep(startOfNewString+i+1, s[i+1:], s1, dp, minTurns)
				//if s == "baacdddaaddaaaaccbddbcabdaabdbbcdcbbbacbddcabcaaa" {
				//	fmt.Println(s1, s2)
				//}
				printOriginCharacter := strangePrinterStep(startOfNewString+i+1, s[i+1:], s2, dp, minTurns) - 1
				//if s == "baacdddaaddaaaaccbddbcabdaabdbbcdcbbbacbddcabcaaa" {
				//fmt.Println(printOriginCharacter, printNewCharacter)
				//fmt.Println(sequences)
				//}
				//fmt.Println(string(letter), startOfNewString, i, dp, "f")

				//fmt.Println(startOfNewString-1+i, startOfNewString, i)
				if printOriginCharacter <= printNewCharacter {
					//fmt.Println(string(letter), startOfNewString, i, dp, "f2")
					//fmt.Println(startOfNewString + i)
					dp[startOfNewString+i] = dp[startOfNewString+i-1]
					//fmt.Println(dp)
					return printOriginCharacter
				}
				dp[startOfNewString-1+i] = dp[startOfNewString-1+i-1] + 1
				fmt.Println(dp)
				minTurns++
				return printNewCharacter
			}
		}

		if len(sequences) > 0 {
			if sequences[len(sequences)-1] != letter {
				sequences = append(sequences, letter)
			}
		} else {
			sequences = append(sequences, letter)
		}

		if dp[startOfNewString+i] == 0 {
			if i == 0 && startOfNewString == 0 {
				dp[startOfNewString+i] = 1
			} else {
				if dp[startOfNewString+i-1] != 0 {
					dp[startOfNewString+i] = dp[startOfNewString+i-1] + 1
				} else {
					minTurns++
					dp[startOfNewString+i] = minTurns - 1
				}
			}
		}
		//sequences = append(sequences, letter)
	}

	//fmt.Println(dp)
	//fmt.Println(minTurns, "f")
	//dp[startOfNewString] = minTurns
	return minTurns
}
