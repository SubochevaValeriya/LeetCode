package main

import (
	"fmt"
)

func main() {
	nums := []int{-10000, -9999, -7, -5, 0, 0, 10000}
	fmt.Println(sortedSquares(nums))
}

func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	i := len(nums) - 1
	result := make([]int, len(nums))
	for left <= right {
		if mod(nums[right]) < mod(nums[left]) {
			result[i] = nums[left] * nums[left]
			left++
		} else {
			result[i] = nums[right] * nums[right]
			right--
		}
		i--
	}
	return result
}

func mod(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
