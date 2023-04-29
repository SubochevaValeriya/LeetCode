package main

import "fmt"

func main() {

	fmt.Println(letterCombinations("23"))
}

var digitMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	var answer = []string{}
	if len(digits) == 0 {
		return answer
	}
	var currentCombination string

	return currentComb(digits, answer, currentCombination)
}

func currentComb(digits string, answer []string, currentCombination string) []string {
	if len(digits) == 0 {
		answer = append(answer, currentCombination)
		return answer
	}

	for _, letter := range digitMap[string(digits[0])] {
		currentCombination += string(letter)
		answer = currentComb(digits[1:], answer, currentCombination)
		currentCombination = currentCombination[:len(currentCombination)-1]
	}

	return answer
}
