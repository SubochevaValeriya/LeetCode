package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))
}

func toGoatLatin(sentence string) string {
	var answer string
	sentenceByWords := strings.Split(sentence, " ")

	for i, word := range sentenceByWords {
		switch word[0] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			answer += word
		default:
			answer += word[1:] + string(word[0])
		}

		answer += "ma" + strings.Repeat("a", i+1) + " "
	}
	return answer[:len(answer)-1]
}
