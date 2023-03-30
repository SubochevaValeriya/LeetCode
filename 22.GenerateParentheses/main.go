package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {

	result := []string{}

	var backtrack func(current string, open, closed int)

	backtrack = func(current string, open, closed int) {
		if open == n && closed == n {
			result = append(result, current)
		}
		if open < n {
			backtrack(current+"(", open+1, closed)
		}
		if closed < open {
			backtrack(current+")", open, closed+1)
		}
	}

	backtrack("", 0, 0)

	return result
}
