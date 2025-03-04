package main

import "fmt"

func main() {
	text1 := "abcde"
	text2 := "ace"

	fmt.Println(longestCommonSubsequence(text1, text2))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1))
	for i := range dp {
		dp[i] = make([]int, len(text2))
	}

	for i := range dp {
		for j := range dp[i] {
			if text1[i] == text2[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
				continue
			}
			if i == 0 && j != 0 {
				dp[i][j] = dp[i][j-1]
			} else if j == 0 && i != 0 {
				dp[i][j] = dp[i-1][j]
			} else if j != 0 && i != 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1], dp[i][j-1])
			}
		}
	}

	for _, ints := range dp {
		fmt.Println(ints)
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

func shortLongTexts(text1, text2 string) (string, string) {
	if len(text1) < len(text2) {
		return text1, text2
	}

	return text2, text1
}
