package main

import "fmt"

func main() {
	nums := []int{0, 4, 0, 3, 2}
	k := 1
	fmt.Println(findMaxAverage(nums, k))
}

func findMaxAverage(nums []int, k int) float64 {
	left, right := 0, k-1
	var maxSum int
	for i := 0; i <= right; i++ {
		maxSum += nums[i]
	}
	sum := maxSum
	right++
	left++
	for ; right <= len(nums)-1; right++ {
		sum = sum - nums[left-1] + nums[right]
		if sum > maxSum {
			maxSum = sum
		}
		fmt.Println(left, right, sum)
		left++
	}

	return float64(maxSum) / float64(k)
}
