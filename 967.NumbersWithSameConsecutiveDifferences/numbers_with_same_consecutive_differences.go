package main

import (
	"fmt"
	"strconv"
)

func numsSameConsecDiff(n int, k int) []int {
	answer := []int{}
	for i := 1; i < 10; i++ {
		fmt.Println(answer)
		digit := strconv.Itoa(i)
		lastDigit := i
		for j := 0; j < 10 && j >= 0; j++ {
			//fmt.Println(i, j, lastDigit)
			if lastDigit+k == j || lastDigit-k == j {
				lastDigit = j
				digit += strconv.Itoa(j)
			}

			if len(digit) == n {
				digitInt, _ := strconv.Atoi(digit)
				answer = append(answer, digitInt)
				digit = strconv.Itoa(i)
				continue
			}
			if j+1 >= 10 && lastDigit-k >= 0 {
				j = -1
				break
			}
		}
	}
	return answer
}
