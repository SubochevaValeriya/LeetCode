package main

import "fmt"

func main() {

}

func numberOfSubarrays(nums []int, k int) int {

	fmt.Println(atMostK(nums, k), atMostK(nums, k-1))
	return atMostK(nums, k) - atMostK(nums, k-1)
}

func atMostK(nums []int, k int) int {
	left, right := 0, 0
	odd := 0
	countOfSubarrays := 0

	for right < len(nums) {
		if nums[right]%2 == 1 {
			odd++
		}
		right++
		for odd > k {
			if nums[left]%2 == 1 {
				odd--
			}
			left++
		}
		countOfSubarrays = countOfSubarrays + right - left
	}
	return countOfSubarrays
}
