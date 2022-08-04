package main

import "fmt"

func main() {
	nums := []int{
		1, 0, 1, 1,
	}
	k := 1
	fmt.Println(containsNearbyDuplicate(nums, k))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	dict := map[int]int{}
	for i, value := range nums {
		//		fmt.Println(dict)
		if _, found := dict[value]; found {
			if i-dict[value] <= k {
				return true
			}
			//continue
		}
		dict[value] = i
	}

	return false
}
