package main

func main() {

}

func reversePrefix(word string, ch byte) string {
	for i := range word {
		if word[i] == ch {
			newPrefix := reverse(word[:i+1])
			if i == len(word)-1 {
				return newPrefix
			}

			return newPrefix + word[i+1:]
		}
	}

	return word
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
