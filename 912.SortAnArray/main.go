package main

import "fmt"

func main() {
	nums := []int{-4, 0, 7, 4, 9, -5, -1, 0, -7, -1}
	fmt.Println(sortArray(nums))
}

func sortArray(nums []int) []int {
	base, baseIndx := nums[0], 0
	left, right := 0, len(nums)-1

	// Move base element to the right
	nums[baseIndx], nums[right] = nums[right], nums[baseIndx]

	// Go through the slice and put smaller elements to the left part
	for i := 0; i < len(nums); i++ {
		if nums[i] < base {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}

	// Put base element between smaller and higher elements (after last smaller element)
	nums[left], nums[right] = nums[right], nums[left]

	// Recursively use quick sort algorithm for left and right parts (if their length > 1)
	if len(nums[:left]) > 1 {
		sortArray(nums[:left])
	}

	if len(nums[left+1:]) > 1 {
		sortArray(nums[left+1:])
	}

	return nums
}
