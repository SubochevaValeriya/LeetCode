package main

import "fmt"

func main() {

	s := "y#fo##f"
	t := "y#f#o##f"

	fmt.Println(backspaceCompare(s, t))
}

func backspaceCompare(s string, t string) bool {
	stack1, stack2 := makeStack(s), makeStack(t)

	fmt.Println(string(stack1), string(stack2))
	if len(stack1) != len(stack2) {
		return false
	}
	for i := 0; i < len(stack1); i++ {
		if stack1[i] != stack2[i] {
			return false
		}
	}

	return true
}

func makeStack(s string) []rune {
	stack := []rune{}
	for _, word := range s {
		if string(word) != "#" {
			stack = append(stack, word)
			continue
		}

		if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}

	}

	return stack
}
