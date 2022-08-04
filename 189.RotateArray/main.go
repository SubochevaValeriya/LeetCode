package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 8
	//7 6 5 4 3 2 1
	//1 7 6 5 4 3 2 1

	// nums = 1 2 3
	//0
	//1
	//2
	//3 - 2
	//= 1
	//2
	//2
	//3
	//1
	//0
	//0
	//0
	//2

	fmt.Println(rotate(nums, k))
}

//1 2 3
//
//5 / len
//
//
//3 1 2
//2 3 1
//1 2 3
//3 1 2
//2 3 1
//
//
//
//2 3 1
//
//1 2 3

func rotate(nums []int, k int) []int {

	lenArray := len(nums)
	rotatedArray := []int{}

	//1 3 4
	//
	//4
	//
	//4 1 3
	//3 4 1
	//1 3 4
	//4 1 3
	//

	k = k % lenArray

	//for i := lenArray-1; i;

	leftPart := nums[lenArray-k:]
	rightPart := nums[0 : lenArray-k]
	//copy(nums, leftPart)
	rotatedArray = append(rotatedArray, leftPart...)
	rotatedArray = append(rotatedArray, rightPart...)
	copy(nums, rotatedArray)
	return nums
}
