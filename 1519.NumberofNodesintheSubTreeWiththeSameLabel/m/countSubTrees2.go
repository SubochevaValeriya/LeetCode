package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(countSubTrees(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, "abaedcd"))
	fmt.Println(countSubTrees(4, [][]int{{0, 1}, {1, 2}, {0, 3}}, "bbbb"))

	fmt.Println(countSubTrees(4, [][]int{{0, 2}, {0, 3}, {1, 2}}, "aeed"))
	fmt.Println(countSubTrees(5, [][]int{{0, 1}, {0, 2}, {1, 3}, {0, 4}}, "aabab"))
	fmt.Println(countSubTrees(25, [][]int{{4, 0}, {5, 4}, {12, 5}, {3, 12}, {18, 3}, {10, 18}, {8, 5}, {16, 8}, {14, 16}, {13, 16}, {9, 13}, {22, 9}, {2, 5}, {6, 2}, {1, 6}, {11, 1}, {15, 11}, {20, 11}, {7, 20}, {19, 1}, {17, 19}, {23, 19}, {24, 2}, {21, 24}}, "hcheiavadwjctaortvpsflssg"))
	http.ListenAndServe("localhost:8080", nil)

}

func countSubTrees(n int, edges [][]int, labels string) []int {
	answer := make([]int, n)

	edgesMap := map[int][]int{}
	vertexMap := map[int]map[string]int{}
	var parent, child int

	for _, edge := range edges {
		if edge[0] == 0 || edge[1] == 0 {
			switch {
			case edge[0] == 0:
				parent, child = edge[0], edge[1]
			case edge[1] == 0:
				parent, child = edge[1], edge[0]
			}
			edgesMap[parent] = append(edgesMap[parent], child)
			vertexMap[parent] = map[string]int{}
			vertexMap[child] = map[string]int{}
			continue
		}

		_, lastVertexRight := vertexMap[edge[1]]
		_, lastVertexLeft := vertexMap[edge[0]]

		switch {
		case lastVertexRight:
			parent, child = edge[1], edge[0]
			vertexMap[edge[0]] = map[string]int{}
		default:
			parent, child = edge[0], edge[1]
			vertexMap[edge[0]] = map[string]int{}
		}

		edgesMap[parent] = append(edgesMap[parent], child)

		if lastVertexLeft {
			vertexMap[edge[1]] = map[string]int{}
		}
	}
	fmt.Println(edgesMap)
	//labelsMap := make([]map[string]int, n)
	fmt.Println(vertexMap)
	vertexMap = DFS(labels, edgesMap, vertexMap, 0)

	for i, label := range labels {
		answer[i] = vertexMap[i][string(label)]
	}

	return answer
}

func DFS(labels string, edgesMap map[int][]int, vertexMap map[int]map[string]int, firstNode int) map[int]map[string]int {
	for _, node := range edgesMap[firstNode] {
		if len(vertexMap[node]) == 0 {
			vertexMap = DFS(labels, edgesMap, vertexMap, node)
		}
		for k, v := range vertexMap[node] {
			vertexMap[firstNode][k] += v
		}
	}
	vertexMap[firstNode][string(labels[firstNode])]++

	return vertexMap
}

//func DFS(label byte, edgesMap map[int][]int, firstNode int, labels string) int {
//	numberOfNodes := 0
//	for _, node := range edgesMap[firstNode] {
//		if labels[node] == label {
//			numberOfNodes++
//		}
//		numberOfNodes += DFS(label, edgesMap, node, labels)
//	}
//	return numberOfNodes
//}
