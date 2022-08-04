package main

import (
	"fmt"
)

func MaxArea(height []int) int {
	start, end := 0, len(height)-1
	var area, maxArea int

	for end >= start {
		area = (end - start) * min(height[end], height[start])
		fmt.Println(height[end], height[start], end, start, area)
		if area > maxArea {
			maxArea = area
		}
		if height[end] > height[start] {
			start++
		} else {
			end--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
