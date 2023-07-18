package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2

	fmt.Println(longestOnes(nums, k))
}

func longestOnes(nums []int, k int) int {
	canFlip := k
	max := 0
	current := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			current++
		} else {
			if canFlip == 0 {
				max = int(math.Max(float64(max), float64(current)))
				for j := i - current; j <= i; j++ {
					if nums[j] == 0 {
						canFlip++
						current = i - j - 1
						fmt.Println(current, "f", i, j)
						break
					}
				}
			}
			canFlip--
			current++
		}
	}

	return int(math.Max(float64(max), float64(current)))
}
