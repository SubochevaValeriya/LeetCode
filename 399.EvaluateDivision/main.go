package main

import (
	"fmt"
)

func main() {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	fmt.Println(calcEquation(equations, values, queries))
}

type vertex struct {
	number string
	value  float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := buildGraph(equations, values)
	results := make([]float64, len(queries))
	for i, query := range queries {
		graph, results[i] = dfs(graph, map[string]bool{}, query[0], query[1], 1)
	}

	return results
}

func buildGraph(equations [][]string, values []float64) map[string][]vertex {
	graph := map[string][]vertex{}

	for i, equation := range equations {
		graph = buildRelation(graph, equation[0], equation[1], values[i])
		graph = buildRelation(graph, equation[1], equation[0], 1/values[i])
	}

	return graph
}

func buildRelation(graph map[string][]vertex, numerator, denominator string, value float64) map[string][]vertex {
	if graph[numerator] == nil {
		graph[numerator] = []vertex{}
	}
	graph[numerator] = append(graph[numerator], vertex{
		number: denominator,
		value:  value,
	})

	return graph
}

func dfs(graph map[string][]vertex, visited map[string]bool, numerator, denominator string, multiplier float64) (map[string][]vertex, float64) {
	if graph[numerator] == nil || graph[denominator] == nil || visited[numerator] {
		return graph, -1
	}
	visited[numerator] = true

	for _, relation := range graph[numerator] {
		if relation.number == denominator {
			return graph, multiplier * relation.value
		}
		graph, result := dfs(graph, visited, relation.number, denominator, relation.value*multiplier)
		if result != -1 {
			return graph, result
		}
	}
	return graph, -1
}
