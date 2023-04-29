package main

import "fmt"

func main() {
	word1 := "abccd"
	word2 := "pqrs"
	fmt.Println(mergeAlternately(word1, word2))
}

func mergeAlternately(word1 string, word2 string) string {

	first, second := 0, 0

	newString := make([]byte, len(word1)+len(word2))
	for i := 0; i < len(newString); i++ {
		if first >= len(word1) {
			newString = append(newString[:i], word2[second:]...)
			break
		}
		if second >= len(word2) {
			newString = append(newString[:i], word1[first:]...)
			break
		}
		if i%2 == 0 {
			newString[i] = word1[first]
			first++
		} else {
			newString[i] = word2[second]
			second++
		}
	}
	return string(newString)
}
