package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	//nums := []int{1, 2}
	//nums := []int{-1}
	fmt.Println(summaryRanges(nums))
}

func summaryRanges(nums []int) []string {
	result := []string{}

	if len(nums) == 0 {
		return result
	}

	if len(nums) == 1 {
		result = append(result, strconv.Itoa(nums[0]))
	}

	var start, end = nums[0], nums[0]
	var oneRange string

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] != nums[i]+1 {
			if start == end {
				oneRange = strconv.Itoa(start)
			} else {
				oneRange = fmt.Sprintf("%d->%d", start, end)
			}
			result = append(result, oneRange)
			start, end = nums[i+1], nums[i+1]
		} else {
			end++
		}

		if i == len(nums)-2 {
			if start == end {
				oneRange = strconv.Itoa(start)
			} else {
				oneRange = fmt.Sprintf("%d->%d", start, end)
			}
			result = append(result, oneRange)
		}
	}

	return result
}
