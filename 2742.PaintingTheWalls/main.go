package main

func main() {

}

func paintWalls(cost []int, time []int) int {
	dp := make([]int, len(cost))
	timePaidPainterUsed := make([]int, len(cost))

	dp[0] = cost[0]
	timePaidPainterUsed[0] = time[0]

	for i := 1; i < len(dp); i++ {
		if timePaidPainterUsed[i-1]-1 >= 0 {
			dp[i] = dp[i-1]
			timePaidPainterUsed[i] = timePaidPainterUsed[i-1] - 1
			continue
		}

	}


	1 1 4
	1 -1
	4-1 3 *1
	2 3 4
	2 -3
	4-3 = 1 *2
	cost time n


}
