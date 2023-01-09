package main

import "fmt"

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	position := []int{-1, -1}

	for right >= left {
		avg := left + (right-left)/2
		fmt.Println(avg, right, left)
		switch {
		case nums[avg] > target:
			right = avg - 1
		case nums[avg] < target:
			left = avg + 1
		case nums[avg] == target:
			if avg == 0 {
				position[0] = avg
			} else if nums[avg-1] != target {
				position[0] = avg
			} else {
				position[0] = startingPosition(nums[:avg+1], target)
			}
			fmt.Println(avg, right, left)
			if avg == len(nums)-1 {
				position[1] = avg
			} else if nums[avg+1] != target {
				position[1] = avg
			} else {
				position[1] = avg + endingPosition(nums[avg:], target)
			}
			return position
		}
	}
	return position
}

func endingPosition(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for right >= left {
		avg := left + (right-left)/2
		switch {
		case nums[avg] > target:
			right = avg
		case nums[avg] == target:
			if avg == len(nums)-1 {
				return avg
			}
			if nums[avg+1] != target {
				return avg
			}
			left = avg + 1
		}
	}
	return -1
}

func startingPosition(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for right >= left {
		avg := left + (right-left)/2
		switch {
		case nums[avg] < target:
			left = avg + 1
		case nums[avg] == target:
			if avg == 0 {
				return avg
			}
			if nums[avg-1] != target {
				return avg
			}
			right = avg - 1
		}
	}
	return -1
}
