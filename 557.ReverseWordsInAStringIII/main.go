package main

func main() {

}

func reverseWords(s string) string {

	reversedWords := ""
	start, end := 0, 0

	for ; end < len(s); end++ {
		if string(s[end]) == " " {
			reversedWords += reverseWord(s[start:end]) + " "
			start = end + 1
		}
	}

	if start < end {
		reversedWords += reverseWord(s[start:])
	}

	return reversedWords
}

func reverseWord(s string) string {
	reversedWord := ""

	for i := len(s) - 1; i >= 0; i-- {
		reversedWord += string(s[i])
	}

	return reversedWord
}
