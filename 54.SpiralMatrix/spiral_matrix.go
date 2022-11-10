package main

import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	firstI := 0
	lastI := len(matrix) - 1
	firstJ := 0
	lastJ := len(matrix[0]) - 1
	for i := 0; i < len(matrix); {
		for j := 0; j < len(matrix[i]); {
			result = append(result, matrix[i][j])
			if len(result) == len(matrix[i])*len(matrix) {
				return result
			}

			fmt.Println(i, j)
			fmt.Println(matrix[i][j])

			switch {
			case i == firstI && j != lastJ:
				j++
				if j == firstJ {
					firstJ++
				}
			case j == lastJ:
				i++
				firstI++
			case i == lastI:
				j--
				lastJ--
			case j == firstJ:
				i--
				lastI--
			}
		}
	}
	return result
}
