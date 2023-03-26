package main

import (
	"strconv"
)

func numDecodings(s string) int {
	if string(s[0]) == "0" {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1

	for i := 1; i <= len(s); i++ {
		numPrev, _ := strconv.Atoi(string(s[i-1]))
		if numPrev != 0 {
			dp[i] = dp[i-1]
		}

		if i > 1 {
			num, _ := strconv.Atoi(s[i-2 : i])
			if num >= 10 && num <= 26 {
				dp[i] += dp[i-2]
			}
		}
	}
	return dp[len(dp)-1]
}
