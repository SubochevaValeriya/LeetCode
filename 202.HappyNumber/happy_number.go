package main

import (
	"strconv"
)

// https://leetcode.com/problems/happy-number/description/?envType=study-plan&id=level-2

func main() {

}

func isHappy(n int) bool {
	if n == 1 {
		return true
	}

	nStr := strconv.Itoa(n)
	var sum int

	history := map[int]int{}
	history[n]++

	for i := 0; i < len(nStr); i++ {
		digit, _ := strconv.Atoi(string(nStr[i]))
		sum += square(digit)
		if i == len(nStr)-1 {
			if sum == 1 {
				return true
			}
			if _, ok := history[sum]; ok {
				return false
			}
			history[sum]++
			nStr = strconv.Itoa(sum)
			i = -1
			sum = 0
		}
	}
	return false
}

func square(n int) int {
	return n * n
}
