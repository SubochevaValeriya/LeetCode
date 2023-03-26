package main

import (
	"fmt"
	"sort"
)

func main() {

	n := 6
	connections := [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}

	fmt.Println(minReorder(n, connections))
}

func minReorder(n int, connections [][]int) int {

	graph := make([][]int, n)
	visited := map[int]bool{}
	var minReord int

	for _, connection := range connections {
		from := connection[0]
		to := connection[1]
		graph[from] = append(graph[from], -to)
		graph[to] = append(graph[to], from)
	}
	start := 0
	dfs(start, &minReord, graph, visited)
	fmt.Println(graph)
	return minReord
}

func dfs(start int, minReord *int, graph [][]int, visited map[int]bool) {
	visited[start] = true
	for _, v := range graph[start] {
		if visited[abs(v)] == false {
			if v < 0 {
				*minReord++
			}
			dfs(abs(v), minReord, graph, visited)
		}
	}
	fmt.Println(start, *minReord)
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func minReorderOld(n int, connections [][]int) int {

	graph := make([][]int, n)
	visited := map[int]bool{}
	sort.Slice(connections, func(i, j int) bool {
		if connections[i][0] == connections[j][0] {
			return connections[i][1] < connections[j][1]
		}
		return connections[i][0] < connections[j][0]
	})
	var minReord int

	for _, connection := range connections {
		from := connection[0]
		to := connection[1]
		if to != 0 {
			if from == 0 || visited[from] == true {
				fmt.Println(connection)
				minReord++
				to, from = from, to
			}
		}
		visited[from] = true
		graph[from] = append(graph[from], to)

	}
	fmt.Println(minReord)
	return minReord
}
