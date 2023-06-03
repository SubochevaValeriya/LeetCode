package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {

	result := make([]int, len(nums))
	multiplier := 1

	for i := 0; i < len(nums); i++ {
		if result[i] == 0 {
			result[i] = 1
		}
		result[i] *= multiplier
		multiplier *= nums[i]
	}

	multiplier = 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= multiplier
		multiplier *= nums[i]
	}

	return result
}
