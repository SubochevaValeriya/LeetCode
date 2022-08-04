package main

import (
	"fmt"
)

func main() {
	nums := []int{-10000, -9999, -7, -5, 0, 0, 10000}
	nums2 := []int{11, 2, 4, 3, 10}
	fmt.Println(sortedSquares2(nums))
	fmt.Println(sorted(nums2))
}

func sortedSquares2(nums []int) []int {
	if len(nums) == 1 {
		nums[0] = nums[0] * nums[0]
		return nums
	}
	result := make([]int, len(nums))
	left := 0
	right := len(nums) - 1
	i := 1
	for left <= right {
		if nums[right]*nums[right] < nums[left]*nums[left] {
			result[len(nums)-i] = nums[left] * nums[left]
			left++
			i++
		} else if nums[right]*nums[right] > nums[left]*nums[left] {
			result[len(nums)-i] = nums[right] * nums[right]
			right--
			i++
		} else if right == left {
			result[len(nums)-i] = nums[left] * nums[left]
			right--
		} else {
			result[len(nums)-i] = nums[left] * nums[left]
			result[len(nums)-i-1] = nums[right] * nums[right]
			left++
			right--
			i += 2
		}
		fmt.Println(result)
	}
	return result
}

func sortedSquares(nums []int) []int {
	result := []int{}
	for _, value := range nums {
		valueSq := value * value
		result = append(result, valueSq)
	}
	return sorted(result)
}

func sorted(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	middlePart := []int{}
	leftPart := []int{}
	rightPart := []int{}
	basic := nums[0]
	for _, value := range nums {
		if value < basic {
			leftPart = append(leftPart, value)
		} else if value > basic {
			rightPart = append(rightPart, value)
		} else {
			middlePart = append(middlePart, value)
		}
	}

	leftPart = sorted(leftPart)
	rightPart = sorted(rightPart)

	leftPart = append(leftPart, middlePart...)
	leftPart = append(leftPart, rightPart...)

	return leftPart
}
