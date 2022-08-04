package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}

	fmt.Println(moveZeroes(nums))
}
func moveZeroes(nums []int) []int {
	left, right := 0, 1
	lenArray := len(nums)
	//nonZeroArray := []int{}
	//zeroArray := []int{}
	//newArr := []int{}
	//for _, value := range nums {
	//	if value == 0 {
	//		zeroArray = append(zeroArray, value)
	//	} else {
	//		nonZeroArray = append(nonZeroArray, value)
	//	}
	//}
	//newArr = append(newArr, nonZeroArray...)
	//newArr = append(newArr, zeroArray...)
	//copy(nums, newArr)

	for right < lenArray && left < lenArray {
		if nums[left] == 0 {
			nums[left], nums[right] = nums[right], nums[left]
			right++
			if nums[left] != 0 {
				left++
			}
		} else {
			left++
			right++
		}
		fmt.Println(nums, left, right)
	}
	return nums
}
