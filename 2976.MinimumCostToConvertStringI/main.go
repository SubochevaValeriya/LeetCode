package main

import (
	"fmt"
	"maps"
)

func main() {
	source := "aababdbddc"
	target := "adcbbbcdba"
	original := []byte{'a', 'd', 'b', 'a', 'd', 'c', 'd', 'b'}
	changed := []byte{'b', 'a', 'd', 'c', 'c', 'a', 'b', 'a'}
	cost := []int{10, 6, 8, 3, 6, 10, 8, 6}
	fmt.Println(minimumCost(source, target, original, changed, cost))
}

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	graph := constructGraph(original, changed, cost)
	minCost := 0
	fmt.Println(graph)
	shortedPaths := map[byte]map[byte]int{}
	for i := range source {
		costed, ok := shortedPaths[source[i]][target[i]]
		if ok {
			minCost += costed
			continue
		}
		visited := map[byte]bool{}
		costed, graph = costShortestPath(graph, visited, source[i], target[i], 0)
		fmt.Println(costed)
		if costed == -1 {
			return -1
		}
		if _, ok := shortedPaths[source[i]]; !ok {
			shortedPaths[source[i]] = map[byte]int{}
		}
		shortedPaths[source[i]][target[i]] = costed
		graph[source[i]][target[i]] = costed
		minCost += costed
	}
	fmt.Println(shortedPaths)
	return int64(minCost)
}

func constructGraph(original []byte, changed []byte, cost []int) map[byte]map[byte]int {
	graph := map[byte]map[byte]int{}

	for i := range original {
		if _, ok := graph[original[i]]; !ok {
			graph[original[i]] = map[byte]int{}
		}
		if _, ok := graph[original[i]][changed[i]]; !ok {
			graph[original[i]][changed[i]] = cost[i]
		} else {
			graph[original[i]][changed[i]] = min(cost[i], graph[original[i]][changed[i]])
		}
	}

	return graph
}

func costShortestPath(graph map[byte]map[byte]int, visited map[byte]bool, currentLetter, goal byte, costed int) (int, map[byte]map[byte]int) {
	currentLetters := string(currentLetter)
	goals := string(goal)
	if visited[currentLetter] {
		return -1, graph
	}
	if currentLetter == goal {
		return costed, graph
	}
	visited[currentLetter] = true
	minCost := -1
	for letter, cost := range graph[currentLetter] {
		var costToGoal int
		letters := string(letter)
		visited2 := map[byte]bool{}
		maps.Copy(visited2, visited)
		fmt.Println(currentLetters, goals, letters)
		costToGoal, graph = costShortestPath(graph, visited2, letter, goal, cost)
		if costToGoal == -1 {
			continue
		}
		if minCost == -1 {
			minCost = costToGoal
		} else {
			minCost = min(minCost, costToGoal)
		}
	}
	if minCost == -1 {
		return -1, graph
	}

	if _, ok := graph[currentLetter][goal]; !ok {
		graph[currentLetter][goal] = costed + minCost
	} else {
		graph[currentLetter][goal] = min(graph[currentLetter][goal], costed+minCost)
	}

	return costed + minCost, graph
}
