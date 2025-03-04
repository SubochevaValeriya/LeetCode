package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "0100"
	fmt.Println(toCamelCaseFromSnakeCase("MINOR_UPGRADE"))
	fmt.Println(minOperations(s))
}

func toCamelCaseFromSnakeCase(str string) string {
	str = strings.ToLower(str)
	words := strings.Split(str, "_")
	newStr := ""
	for _, word := range words {
		newStr += strings.ToUpper(word[:1]) + word[1:]
	}

	return newStr
}

func minOperations(s string) int {
	toChangeToOne := 0
	toChangetoZero := 0

	for i, num := range s {
		if i%2 == 0 && num != '1' {
			toChangeToOne++
		}

		if i%2 == 1 && num != '0' {
			toChangeToOne++
		}
	}

	for i, num := range s {
		if i%2 == 0 && num != '0' {
			toChangetoZero++
		}

		if i%2 == 1 && num != '1' {
			toChangetoZero++
		}
	}

	if toChangeToOne < toChangetoZero {
		return toChangeToOne
	}

	return toChangetoZero
}
