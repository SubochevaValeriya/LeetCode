package main

import (
	"strconv"
	"strings"
)

func main() {

}

func numDecodings(s string) int {
	number,  _ := strconv.Atoi(string(s[0]))
	if number == 0{
		return 0
	}
	dp := []int{}

	for i := 0; i < len(s); i++ {
		number,  _ := strconv.Atoi(string(s[0]))


		dp[i] =
	}

}
