package main

import (
	"fmt"
	"math"
)

func main() {

	s := "tryhard"
	k := 3

	fmt.Println(maxVowels(s, k))
}

func maxVowels(s string, k int) int {

	max := 0

	for i := 0; i < k; i++ {
		fmt.Println(string(s[i]))
		if isVowel(s[i]) {
			max++
		}
	}

	current := max

	for i := k; i < len(s); i++ {
		if isVowel(s[i-k]) {
			current--
		}
		if isVowel(s[i]) {
			current++
		}

		max = int(math.Max(float64(max), float64(current)))
	}

	return max
}

func isVowel(c uint8) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
