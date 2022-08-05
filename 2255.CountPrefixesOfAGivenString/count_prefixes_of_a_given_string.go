package main

func countPrefixes(words []string, s string) int {
	l := len(s)
	var count int
	for _, word := range words {
		if len(word) > l {
			break
		}
		for i := range word {
			if word[i] != s[i] {
				break
			}
			if i == len(word)-1 {
				count++
			}
		}
	}
	return count
}
