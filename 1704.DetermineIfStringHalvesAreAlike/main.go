package main

import "slices"

func main() {

}

func halvesAreAlike(s string) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	var vowelsCountA, vowelsCountB int

	for i, letter := range s {
		if slices.Contains(vowels, letter) {
			if i < len(s)/2 {
				vowelsCountA++
			} else {
				vowelsCountB++
			}
		}
	}

	return vowelsCountA == vowelsCountB
}
