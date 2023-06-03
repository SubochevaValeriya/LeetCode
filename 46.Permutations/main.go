package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {

	permutations := [][]int{}

	mUsed := map[int]bool{}

	return addNumber(nums, []int{}, mUsed, permutations)
}

func addNumber(nums []int, currentNums []int, mUsed map[int]bool, permutations [][]int) [][]int {

	for i, num := range nums {
		if !mUsed[i] {
			if len(append(currentNums, num)) == len(nums) {
				permutations = append(permutations, append(currentNums, num))
				return permutations
			}
			mUsed[i] = true
			permutations = addNumber(nums, append(currentNums, num), mUsed, permutations)
			mUsed[i] = false
		}
	}
	return permutations
}
