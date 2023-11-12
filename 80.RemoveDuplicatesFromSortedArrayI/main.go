package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 2, 2, 3, 3}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	var count, currentNum, toRemoveLeft int
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != currentNum {
			if count > 2 {
				nums = append(nums[:toRemoveLeft], nums[i:]...)
				i = toRemoveLeft
			}
			currentNum = nums[i]
			count = 1
			continue
		}
		count++
		if count == 3 {
			toRemoveLeft = i
		}
		if i == len(nums)-1 {
			if count > 2 {
				nums = nums[:toRemoveLeft]
			}
		}
	}

	fmt.Println(nums)
	return len(nums)
}
