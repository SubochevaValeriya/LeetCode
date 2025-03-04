package main

import "fmt"

func main() {
	nums := []int{3, 2, 1}
	fmt.Println(findPeakElement(nums))
}

func findPeakElement(nums []int) int {
	for left, right := 0, len(nums)-1; left >= 0 && right < len(nums); {
		mid := left + (right-left)/2
		lowerNeighboursCount := 0

		if mid == 0 {
			lowerNeighboursCount++
		} else if nums[mid-1] < nums[mid] {
			lowerNeighboursCount++
		}

		if mid == len(nums)-1 {
			lowerNeighboursCount++
		} else if nums[mid+1] < nums[mid] {
			lowerNeighboursCount++
		}

		if lowerNeighboursCount == 2 {
			return mid
		}

		if mid == 0 {
			left++
			continue
		}

		if mid == len(nums)-1 {
			right--
			continue
		}

		if nums[mid-1] > nums[mid+1] {
			right--
		} else {
			left++
		}
	}

	return 0
}
