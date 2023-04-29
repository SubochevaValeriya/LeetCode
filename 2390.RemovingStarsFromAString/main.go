package main

import "fmt"

func main() {
	s := "leet**cod*e"
	fmt.Println(removeStars(s))
}

func removeStars(s string) string {
	stack := []rune{}
	for _, letter := range s {
		if letter != '*' {
			stack = append(stack, letter)
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}
