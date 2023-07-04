package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	h := 5
	w := 4
	horizontalCuts := []int{1, 2, 4}
	verticalCuts := []int{1, 3}
	fmt.Println(maxArea(h, w, horizontalCuts, verticalCuts))
}

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {

	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)

	maxH, maxW := maxEdge(h, horizontalCuts), maxEdge(w, verticalCuts)
	fmt.Println(maxH, maxW)
	return (maxH * maxW) % (int(math.Pow(10.0, 9.0) + 7))
}

func maxEdge(size int, cuts []int) int {
	max := 0
	for i := 0; i < len(cuts); i++ {
		if i == 0 {
			max = cuts[i]
			continue
		}
		if cuts[i]-cuts[i-1] > max {
			max = cuts[i] - cuts[i-1]
		}
	}
	if size-cuts[len(cuts)-1] > max {
		max = size - cuts[len(cuts)-1]
	}
	return max
}
