package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
}

func singleNumber(nums []int) int {
	sort.Ints(nums)
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
			continue
		}
		if count != 3 {
			return nums[i-1]
		}
		count = 1
	}
	return nums[len(nums)-1]
}
