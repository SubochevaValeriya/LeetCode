package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{51, 71, 17, 24, 42}
	fmt.Println(maxSum(nums))
}

func maxSum(nums []int) int {
	m := map[int]int{}
	maxSum := -1
	for _, num := range nums {
		maxDig := maxDigit(num)
		if m[maxDig] != 0 && m[maxDig]+num > maxSum {
			maxSum = m[maxDig] + num
		}
		if num > m[maxDig] {
			m[maxDig] = num
		}
	}

	return maxSum
}

func maxDigit(num int) int {
	numsStr := strconv.Itoa(num)
	max := 0
	for _, digit := range numsStr {
		if int(digit) > max {
			max = int(digit)
		}
	}

	return max
}
