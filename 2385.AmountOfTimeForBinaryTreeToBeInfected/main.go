package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	list := []*TreeNode{root}
	graph := map[int][]int{}
	for i := 0; i < len(list); i++ {
		node := list[i]
		if node == nil {
			continue
		}
		if graph[node.Val] == nil {
			graph[node.Val] = []int{}
		}
		if node.Left != nil {
			list = append(list, node.Left)
			graph[node.Val] = append(graph[node.Val], node.Left.Val)
			graph[node.Left.Val] = []int{node.Val}
		}
		if node.Right != nil {
			list = append(list, node.Right)
			graph[node.Val] = append(graph[node.Val], node.Right.Val)
			graph[node.Right.Val] = []int{node.Val}
		}
	}

	distances := map[int]int{start: 0}

	maxDistance := 0
	for node := range graph {
		var distance int
		distance, distances = bfs(node, graph, distances, map[int]bool{}, start)
		if distance > maxDistance {
			maxDistance = distance
		}
	}

	return maxDistance
}

func bfs(node int, graph map[int][]int, distances map[int]int, visited map[int]bool, start int) (int, map[int]int) {
	visited[node] = true
	distance, ok := distances[node]
	if ok {
		return distance, distances
	}

	minDistance := len(graph) + 1
	for _, adjacentNode := range graph[node] {
		if visited[adjacentNode] {
			continue
		}
		var distanceFromStart int
		distanceFromStart, distances = bfs(adjacentNode, graph, distances, visited, start)
		if distanceFromStart == -1 {
			continue
		}
		distanceFromStart++
		if distanceFromStart < minDistance {
			minDistance = distanceFromStart
		}
	}
	if minDistance == len(graph)+1 {
		return -1, distances
	}

	distances[node] = minDistance
	return minDistance, distances
}
