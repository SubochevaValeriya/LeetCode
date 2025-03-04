package main

import (
	"fmt"
	"math"
)

func main() {
	n := 43
	fmt.Println(numSquares(n))
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if isFloatInt(math.Sqrt(float64(i))) {
			dp[i] = 1
			continue
		}
		minNumber := dp[i-1] + 1
		for j := i - 1; j >= 1; j-- {
			count := i / j
			number := dp[j]*count + dp[i-j*count]
			if number < minNumber {
				minNumber = number
			}
			if i-j*2 > 0 {
				break
			}
		}
		dp[i] = minNumber
	}

	return dp[n]
}

func isFloatInt(floatValue float64) bool {
	return math.Mod(floatValue, 1.0) == 0
}
