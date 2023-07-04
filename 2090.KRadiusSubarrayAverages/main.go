package main

import "fmt"

func main() {
	nums := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
	k := 3
	fmt.Println(getAverages(nums, k))
}

func getAverages(nums []int, k int) []int {

	averages := make([]int, len(nums))
	var sum, n int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		n++
		if i < k || i+k >= len(nums) {
			averages[i] = -1
		}
		if n == k*2+1 {
			averages[i-k] = sum / n
			sum -= nums[i-k*2]
			n--
		}
	}
	return averages
}
