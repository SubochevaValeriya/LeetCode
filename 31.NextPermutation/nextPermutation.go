package main

import "sort"

func nextPermutation(nums []int) []int {
	var max int
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	for i, num := range nums {
		if num == max {
			if i == 0 {
				sort.Ints(nums)
				break
			}
			nums[i], nums[i-1] = nums[i-1], nums[i]
			break
		}
	}

	return nums
}
