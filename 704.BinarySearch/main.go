package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//nums, target := values()
	nums := []int{5}
	target := 5
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		if right == left && right == nums[right] {
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
	return -1
}

//targetRange := nums[left:right]
//if target == nums[middle] {
//return middle
//} else if target < nums[middle] {
//search(leftPart, target)
//} else if target > nums[middle] {
//search(rightPart, target)
//} else {
//return middle
//}
//}
//return middle
//}*/
//
func search(nums []int, target int) int {
	var middle, result int
	if len(nums) == 1 {
		middle = 0
		if nums[middle] != target {
			return -1
		} else {
			return middle
		}
	} else {
		middle = len(nums) / 2
		leftPart := nums[0:middle]
		rightPart := nums[middle:]
		if len(nums)-1 >= 0 {
			if target == nums[middle] {
				result = nums[middle]
				return result
			} else if target < nums[middle] {
				middle = search(leftPart, target)
				if middle != -1 {
					for i := range nums {
						if middle == nums[i] {
							result = i
							return result
						}
					}
				} else {
					return -1
				}
			} else if target > nums[middle] {
				middle = search(rightPart, target)
				if middle != -1 {
					for i := range nums {
						if middle == nums[i] {
							result = i
							return result
						}
					}
				} else {
					return -1
				}
			}
		} else {
			return -1
		}
	}
	return result
}

func values() ([]int, int) {
	initial := make([]int, 10e4)
	rand.Seed(time.Now().UnixNano())
	var max int = 10e4
	var min int = -10e4
	for i := range initial {
		initial[i] = rand.Intn(max-min) + min
	}
	target := rand.Intn(max-min) + min
	return initial, target
}
