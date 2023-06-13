package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums := []int{96, 44, 99, 25, 61, 84, 88, 18, 19, 33, 60, 86, 52, 19, 32, 47, 35, 50, 94, 17, 29, 98, 22, 21, 72, 100, 40, 84}
	//fmt.Println(minimalKSum(nums, 35))

	nums := []int{5, 6}
	fmt.Println(minimalKSum(nums, 6))
}

func minimalKSum(nums []int, k int) int64 {

	sort.Ints(nums)
	fmt.Println(nums)
	sum := 0
	left := 1
	for i := 0; i < len(nums); i++ {
		if i != 0 {
			left = nums[i-1] + 1
		}
		for j := left; j < nums[i]; j++ {
			sum += j
			k--
			if k == 0 {
				return int64(sum)
			}
		}
	}

	if k == 0 {
		return int64(sum)
	}

	for i := nums[len(nums)-1] + 1; ; i++ {
		sum += i
		k--
		if k == 0 {
			break
		}
	}
	//
	//for i := 1; i <= 100000000; i++ {
	//	if k == 0 {
	//		break
	//	}
	//	if len(nums) == 0 {
	//		sum += i
	//		k--
	//	} else {
	//		if nums[0] > i {
	//			sum += i
	//			k--
	//		} else {
	//			nums = nums[1:]
	//		}
	//	}
	//}
	//return int64(sum)
	return int64(sum)
}
