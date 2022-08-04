package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "0P"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	sByte := []byte(s)

	base := make([]byte, 0, len(sByte))

	for _, val := range sByte {
		if val >= '0' && val <= '9' || val >= 'a' && val <= 'z' {
			base = append(base, val)
		}
	}

	palindrome := make([]byte, 0, len(base))

	for i := len(base) - 1; i >= 0; i-- {
		palindrome = append(palindrome, base[i])
	}

	fmt.Println(string(base))

	fmt.Println(string(palindrome))

	if string(base) == string(palindrome) {
		return true
	}

	return false
}
