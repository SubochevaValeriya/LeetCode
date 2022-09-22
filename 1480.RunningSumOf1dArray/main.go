package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(runningSum(nums))
}

func runningSum(nums []int) []int {
	answer := []int{}
	var sum int
	for _, num := range nums {
		sum += num
		answer = append(answer, sum)
	}
	return answer
}
