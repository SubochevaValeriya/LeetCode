package main

import (
	"fmt"
)

func main() {
	//nums, target := values()
	nums := []int{1, 3, 5, 6}
	target := 2
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		if right == left && target == nums[right] {
			return right
		} else {
			middle := left + (right-left)/2
			if nums[middle] > target {
				right = middle - 1
			} else if nums[middle] < target {
				left = middle + 1
			} else if nums[middle] == target {
				return middle
			}
		}
	}
	return left
}
