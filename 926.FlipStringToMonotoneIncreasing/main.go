package main

import (
	"fmt"
)

//https://leetcode.com/problems/flip-string-to-monotone-increasing/

func main() {
	//s := "0101100011"
	s := "11011"
	// s := "101010111001010000011101101110"
	//s := "00110"
	//s := "10011111110010111011"
	fmt.Println(minFlipsMonoIncr(s))

}

func minFlipsMonoIncr(s string) int {

	dpZero := make([]int, len(s)+1)
	dpOne := make([]int, len(s)+1)

	for i, digit := range s {
		dpOne[i+1] = dpOne[i]
		dpZero[i+1] = dpZero[i]
		if digit == '0' {
			dpOne[i+1] += 1
		} else {
			dpZero[i+1] += 1
		}
	}

	dpFlips := make([]int, len(s)+1)
	minFlips := len(s)

	for i := 1; i < len(dpFlips); i++ {
		dpFlips[i] = dpZero[i] + dpOne[len(dpOne)-1] - dpOne[i]
		if i == len(dpFlips)-1 {
			dpFlips[i] = min(dpZero[i], dpOne[i])
		}
		if dpFlips[i] < minFlips {
			minFlips = dpFlips[i]
		}
	}
	return minFlips
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
