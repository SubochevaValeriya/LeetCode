package main

import "fmt"

func main() {

	nums := []int{3, 5, 0, 3, 4}
	fmt.Println(find132pattern(nums))
}

type potentialK struct {
	val int
	i   int
}

func find132pattern(nums []int) bool {
	stack := make([]potentialK, 0)
	minimum := nums[0]
	for i := 1; i < len(nums); i++ {
		potentialJ := nums[i]
		for len(stack) > 0 && potentialJ >= stack[len(stack)-1].val {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 && potentialJ < stack[len(stack)-1].val && potentialJ > stack[len(stack)-1].i {
			return true
		}
		stack = append(stack, potentialK{nums[i], minimum})
		if nums[i] < minimum {
			minimum = nums[i]
		}
	}
	return false
}
