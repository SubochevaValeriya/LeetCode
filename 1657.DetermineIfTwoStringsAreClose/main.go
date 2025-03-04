package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	word1 := "uau"
	word2 := "sxx"

	fmt.Println(closeStrings(word1, word2))
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	letters1 := strings.Split(word1, "")
	letters2 := strings.Split(word2, "")

	for _, letter := range letters1 {
		if !strings.Contains(word2, letter) {
			return false
		}
	}

	mCount1 := mSequenceCount(letters1)
	mCount2 := mSequenceCount(letters2)

	for sequence, count := range mCount1 {
		if mCount2[sequence] != count {
			fmt.Println(sequence, count)
			return false
		}
	}

	return true
}

func mSequenceCount(letters []string) map[int]int {
	sort.Strings(letters)
	m := map[int]int{}
	count := 1
	for i := 1; i < len(letters); i++ {
		if letters[i] != letters[i-1] {
			m[count]++
			count = 0
		}
		count++
	}

	m[count]++

	return m
}
