package main

import "fmt"

func main() {
	digits := []int{
		2, 4, 9, 3, 9,
	}
	fmt.Println(plusOne(digits))
}
func plusOne(digits []int) []int {
	l := len(digits) - 1

	for i := l; i >= 0; i-- {
		if digits[i] < 9 && i == l {
			digits[i]++
			return digits
		} else if digits[i] == 10 || (digits[i] == 9 && i == l) {
			digits[i] = 0
			if i > 0 {
				digits[i-1]++
			}
		}
	}

	if digits[0] == 10 || digits[0] == 0 {
		digits[0] = 0
		digits = append([]int{1}, digits...)
	}

	return digits
}
