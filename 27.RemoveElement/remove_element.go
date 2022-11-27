package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	val := 2
	fmt.Println(removeElement(nums, val))

}

func removeElement(nums []int, val int) int {
	//	k := 0
	left, right := 0, len(nums)-1
	for right >= left {
		if nums[left] == val {
			nums[left] = nums[right]
			right--
		} else {
			left++
		}
	}
	return left
}
