package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}

func pivotIndex(nums []int) int {

	sumRight, sumLeft := 0, 0
	for i := 0; i < len(nums); i++ {
		sumRight += nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 {
			sumLeft += nums[i-1]
		}
		sumRight -= nums[i]
		if sumLeft == sumRight {
			return i
		}
	}
	return -1
}
