package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}
	fmt.Println(eraseOverlapIntervals(intervals))
}

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	fmt.Println(intervals)
	dp := make([]int, len(intervals))
	dp[0] = 0
	lastInterval := intervals[0]
	for i := 1; i < len(dp); i++ {
		dp[i] = dp[i-1]
		if intervals[i][0] < lastInterval[1] {
			dp[i] = dp[i-1] + 1
		} else {
			lastInterval = intervals[i]
		}
	}

	return dp[len(dp)-1]
}
