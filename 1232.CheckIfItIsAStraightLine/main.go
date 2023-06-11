package main

import (
	"fmt"
	"math"
)

//https://leetcode.com/problems/check-if-it-is-a-straight-line/description/

func main() {
	//	coordinates := [][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}
	coordinates := [][]int{{0, 0}, {0, 1}, {0, -1}}

	fmt.Println(checkStraightLine(coordinates))
}

func checkStraightLine(coordinates [][]int) bool {
	baseX, baseY := coordinates[0][0], coordinates[0][1]
	var baseSlope float64
	for i := 1; i < len(coordinates); i++ {
		x, y := coordinates[i][0], coordinates[i][1]
		if i == 1 {
			baseSlope = float64(y-baseY) / float64(x-baseX)
			continue
		}
		slope := float64(y-baseY) / float64(x-baseX)
		if (slope == math.Inf(-1) || slope == math.Inf(+1)) && (baseSlope == math.Inf(-1) || baseSlope == math.Inf(+1)) {
			continue
		}
		if slope != baseSlope {
			return false
		}
	}

	return true
}
