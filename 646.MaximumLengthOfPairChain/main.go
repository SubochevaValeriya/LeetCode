package main

import (
	"fmt"
	"sort"
)

func main() {
	//	pairs := [][]int{{1, 2}, {7, 8}, {3, 4}}
	pairs := [][]int{{3, 4}, {2, 3}, {1, 2}}
	fmt.Println(findLongestChain(pairs))
}

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	fmt.Println(pairs)
	chainLen := 1
	dp := make([]int, len(pairs))
	dp[0] = pairs[0][1]
	for i := 1; i < len(pairs); i++ {
		if pairs[i][0] > dp[i-1] {
			chainLen++
			dp[i] = pairs[i][1]
		} else {
			dp[i] = dp[i-1]
		}
	}

	return chainLen
}
