package main

import "fmt"

func FindDuplicate(nums []int) int {
	//for i := 0; i < len(nums)-1; i++ {
	//	for j := i + 1; j < len(nums); j++ {
	//		if nums[i] == nums[j] {
	//			return nums[i]
	//		}
	//	}
	//}
	//sort.Ints(nums)
	//target := 0
	//for i := 0; i < len(nums); i++ {
	//	target ^= nums[i]
	//	fmt.Println(nums, nums[i], target)
	//}
	//
	//return -1
	//m := map[int]int{}
	//for _, num := range nums {
	//	if _, ok := m[num]; ok {
	//		return num
	//	}
	//
	//	m[num]++
	//}
	//
	//return -1

	var target int

	for _, num := range nums {
		target ^= num
	}

	fmt.Println(nums, target)

	for i := 1; i < len(nums); i++ {
		target ^= i
		fmt.Println(nums, i, target)
	}

	return target
}

//var sum, target int

//target := 0
//
//for i := 0; i < len(nums); i++ {
//if i == 1 && target == 0 {
//	return nums[i]
//}
//
//if i >= 2 {
//	if target^nums[i] == target&nums[i-2] {
//		return nums[i]
//	}
//}
//
//target ^= nums[i]
//
//fmt.Println(nums, nums[i], target)

//}
//
//return 0
//for _, val := range nums {
//	if ta
//
//		sum += val
//	target ^= val
//	fmt.Println(nums, val, sum, target)
//}
//
//return (sum - target) / 2
//}
