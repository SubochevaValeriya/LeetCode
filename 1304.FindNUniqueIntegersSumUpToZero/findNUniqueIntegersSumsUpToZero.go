package main

import "fmt"

func main() {
	fmt.Println(sumZero(5))
}

func sumZero(n int) []int {

	answer := []int{}

	if n%2 != 0 {
		answer = append(answer, 0)
	}

	for i := 1; i <= n/2; i++ {
		answer = append(answer, i, -i)
	}

	return answer
}
