package main

import "fmt"

func main() {
	s := "10"

	fmt.Println(maxScore(s))
}

func maxScore(s string) int {
	var answer, zeroSum, onesSum int
	zeroSums := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zeroSum++
			zeroSums[i] = zeroSum
		}
	}

	fmt.Println(zeroSums)

	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '1' {
			onesSum++
		}
		zeroSumPr := 0
		if i > 0 {
			zeroSumPr = zeroSums[i-1]
		}
		if zeroSumPr+onesSum > answer {
			answer = zeroSumPr + onesSum
		}
	}

	return answer
}
