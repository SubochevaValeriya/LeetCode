package main

func main() {

}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make(map[int]bool, n)
	count := 0
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		bfs(isConnected, visited, i)
		count++
	}
	return count
}

func bfs(isConnected [][]int, visited map[int]bool, i int) {
	if visited[i] {
		return
	}
	visited[i] = true
	for j := range isConnected[i] {
		if isConnected[i][j] == 1 {
			bfs(isConnected, visited, j)
		}
	}
}
