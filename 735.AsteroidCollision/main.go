package main

import (
	"fmt"
	"math"
)

func main() {
	asteroids := []int{-2, -1, 1, -2}
	fmt.Println(asteroidCollision(asteroids))
}

func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for i := 0; i < len(asteroids); i++ {
		switch {
		case len(stack) == 0:
			stack = append(stack, asteroids[i])
		case isTheSameSign(stack[len(stack)-1], asteroids[i]) || stack[len(stack)-1] < 0:
			stack = append(stack, asteroids[i])
		case math.Abs(float64(asteroids[i])) > math.Abs(float64(stack[len(stack)-1])):
			stack = stack[:len(stack)-1]
			i--
		case math.Abs(float64(asteroids[i])) == math.Abs(float64(stack[len(stack)-1])):
			stack = stack[:len(stack)-1]
		}
	}

	return stack
}

func isTheSameSign(a, b int) bool {
	return (a >= 0 && b >= 0) || (a < 0 && b < 0)
}
