package main

import (
	"fmt"
)

func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}

	//target := 15
	//nums := []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8}
	fmt.Println(minSubArrayLen(target, nums))
}

func minSubArrayLen(target int, nums []int) int {
	var left, right int
	sum := nums[0]
	minLenght := 1000000
	for left <= right {
		if right-left+1 >= minLenght {
			sum -= nums[left]
			left++
		}
		if sum < target {
			if right == len(nums)-1 {
				break
			}
			right++
			sum += nums[right]
		}
		if sum >= target {
			if right-left+1 < minLenght {
				minLenght = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}

	if minLenght == 1000000 {
		return 0
	}
	return minLenght
}
