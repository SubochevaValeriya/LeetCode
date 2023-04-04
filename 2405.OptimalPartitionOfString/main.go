package main

import "fmt"

func main() {

	s := "hdklqkcssgxlvehva"

	//s := "abacaba"

	fmt.Println(partitionString(s))
}

func partitionString(s string) int {

	m := map[rune]int{}

	count := 1

	for _, letter := range s {
		m[letter]++
		if m[letter] > 1 {
			fmt.Println(string(letter))
			m = map[rune]int{}
			m[letter]++
			count++
		}
	}

	return count
}
