package main

import "fmt"

func main() {
	nums := []int{
		1, 2, 3, 1,
	}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	dict := map[int]int{}
	for _, value := range nums {
		if _, found := dict[value]; found {
			return true
		}
		dict[value] = 1
	}

	return false
}
