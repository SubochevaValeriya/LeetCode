package main

import "fmt"

func main() {
	n := 5
	queries := [][]int{{2, 4}, {0, 2}, {0, 4}}
	fmt.Println(shortestDistanceAfterQueries(n, queries))
}

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	shortestDistances := make([]int, len(queries))
	roads := map[int][]int{}
	for i := 0; i < n; i++ {
		if i != n-1 {
			roads[i] = []int{i + 1}
		}
	}
	for i, query := range queries {
		roads[query[0]] = append(roads[query[0]], query[1])
		shortestDistances[i] = minDistance(roads, n)
	}

	return shortestDistances
}

func minDistance(roads map[int][]int, n int) int {
	dp := make([]int, n)
	dp[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		shortestDistance := n
		for _, neighbour := range roads[i] {
			shortestDistance = min(dp[neighbour]+1, shortestDistance)
		}
		dp[i] = shortestDistance
	}

	return dp[0]
}

func bfs(roads map[int][]int, from, goal int) int {
	var shortestDistance = goal
	for _, to := range roads[from] {
		distanceCurrent := 1
		if to != goal {
			distanceCurrent = 1 + bfs(roads, to, goal)
		}
		if distanceCurrent < shortestDistance {
			shortestDistance = distanceCurrent
		}
	}

	return shortestDistance
}
