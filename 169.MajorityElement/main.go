package main

import "fmt"

func main() {
	nums := []int{
		3, 2, 3,
	}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	dict := map[int]int{}
	counter := 0
	counterNum := 0
	for _, val := range nums {
		dict[val] += 1
		if counter < dict[val] {
			counter = dict[val]
			counterNum = val
		}
	}

	return counterNum
}
