package main

import "fmt"

func main() {
	fmt.Println(rearrangeCharacters("ilovecodingonleetcode", "code"))
}

func rearrangeCharacters(s string, target string) int {
	m := map[rune]int{}
	for _, letter := range s {
		m[letter]++
	}

	maxNumberOfCopies := 0
	for i := 0; i < len(target); i++ {
		letter := rune(target[i])
		if m[letter] <= 0 {
			return maxNumberOfCopies
		}
		m[letter]--
		if i == len(target)-1 {
			maxNumberOfCopies++
			i = -1
		}
	}
	return maxNumberOfCopies
}
