package main

import "fmt"

func main() {
	//nums := []int{2, 3, 3, 2, 2, 4, 2, 3, 4}
	nums := []int{14, 12, 14, 14, 12, 14, 14, 12, 12, 12, 12, 14, 14, 12, 14, 14, 14, 12, 12}
	fmt.Println(minOperations(nums))
}

func minOperations(nums []int) int {
	mFrequency := map[int]int{}

	for _, num := range nums {
		mFrequency[num]++
	}

	minOp := 0

	for _, count := range mFrequency {
		for countThree := count / 3; countThree >= 0; countThree-- {
			if (count-3*countThree)%2 == 0 && count-3*countThree != 1 {
				minOp = minOp + countThree + (count-3*countThree)/2
				break
			}
			if countThree == 0 {
				return -1
			}
		}
	}

	return minOp
}
