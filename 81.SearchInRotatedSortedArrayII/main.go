package main

import (
	"fmt"
	"math"
	"regexp"
)

func main() {
	fmt.Println(roundTwoDecimalPlaces(123.456))
	fmt.Println(regexp.Match("^[1-9][0-9]*|[e]|[.][0-9]*|[e]|[.]$", []byte("0")))

}
func roundTwoDecimalPlaces(num float32) float64 {
	return math.Round(float64(num*100)) / 100
}

func is() bool {
	return 1 == 1 || 1 == 2
}

func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left <= right {
		if nums[left] >= nums[right] {
			if target == nums[left] {
				return true
			}
			if target == nums[right] {
				return true
			}
			if target > nums[right] {
				right--
			} else if target < nums[right] {
				left++
			}
			continue
		}
		medium := left + (right-left)/2

		if nums[medium] == target {
			return true
		}
		if nums[medium] > target {
			right--
		}
		if nums[medium] < target {
			left++
		}
	}

	return false
}
