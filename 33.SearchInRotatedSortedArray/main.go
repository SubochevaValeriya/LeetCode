package main

import "fmt"

func main() {

	nums := []int{3, 1}
	target := 0

	fmt.Println(search(nums, target))

}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		if nums[left] >= nums[right] {
			if target == nums[left] {
				return left
			}
			if target == nums[right] {
				return right
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
			return medium
		}
		if nums[medium] > target {
			right--
		}
		if nums[medium] < target {
			left++
		}
	}

	return -1
}
