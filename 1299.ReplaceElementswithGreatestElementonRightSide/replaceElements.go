package main

import "fmt"

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))
}

func replaceElements(arr []int) []int {
	maxRight := 0
	for i := len(arr) - 1; i >= 0; i-- {
		num := arr[i]
		if maxRight == 0 {
			arr[i] = -1
		} else {
			arr[i] = maxRight
		}
		if num > maxRight {
			maxRight = num
		}
	}

	return arr
}
