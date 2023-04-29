package main

import "fmt"

func main() {

	//	grid := [][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}
	grid := [][]int{{0, 0, 0, 1, 1, 1, 0, 1, 0, 0}, {1, 1, 0, 0, 0, 1, 0, 1, 1, 1}, {0, 0, 0, 1, 1, 1, 0, 1, 0, 0}, {0, 1, 1, 0, 0, 0, 1, 0, 1, 0}, {0, 1, 1, 1, 1, 1, 0, 0, 1, 0}, {0, 0, 1, 0, 1, 1, 1, 1, 0, 1}, {0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, {0, 0, 1, 0, 0, 1, 0, 1, 0, 1}, {1, 0, 1, 0, 1, 1, 0, 0, 0, 0}, {0, 0, 0, 0, 1, 1, 0, 0, 0, 1}}
	for _, ints := range grid {
		fmt.Println(ints)
	}
	enclaves := map[[2]int]bool{}
	visited := map[[2]int]struct{}{}
	fmt.Println(bfs(5, 2, grid, enclaves, visited))
	fmt.Println(numEnclaves(grid))
}
func numEnclaves(grid [][]int) int {
	enclaves := map[[2]int]bool{}
	visited := map[[2]int]struct{}{}
	answer := 0
	var isEnclave bool
	for i := range grid {
		for j := range grid[i] {
			_, ok := visited[[2]int{i, j}]
			if ok {
				if grid[i][j] == 1 && enclaves[[2]int{i, j}] {
					answer++
				}
				continue
			}
			if grid[i][j] == 0 {
				enclaves[[2]int{i, j}] = true
				continue
			}
			if i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[i])-1 {
				enclaves[[2]int{i, j}] = false
				continue
			}
			enclaves, visited, isEnclave = bfs(i, j, grid, enclaves, visited)
			if isEnclave {
				answer++
			}
		}
	}
	return answer
}

func bfs(i int, j int, grid [][]int, enclaves map[[2]int]bool, visited map[[2]int]struct{}) (map[[2]int]bool, map[[2]int]struct{}, bool) {
	isEnclaves, ok := enclaves[[2]int{i, j}]
	if ok {
		return enclaves, visited, isEnclaves
	}
	_, ok = visited[[2]int{i, j}]
	if ok {
		return enclaves, visited, true
	}
	visited[[2]int{i, j}] = struct{}{}

	returnFalseOrTrue := func(falseOrTrue bool) (map[[2]int]bool, map[[2]int]struct{}, bool) {
		enclaves[[2]int{i, j}] = falseOrTrue
		return enclaves, visited, falseOrTrue
	}

	if grid[i][j] == 0 {
		return returnFalseOrTrue(true)
	}
	if i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[i])-1 {
		return returnFalseOrTrue(false)
	}
	enclaves, visited, isEnclaves = bfs(i-1, j, grid, enclaves, visited)
	if !isEnclaves {
		return returnFalseOrTrue(false)
	}
	enclaves, visited, isEnclaves = bfs(i, j-1, grid, enclaves, visited)
	if !isEnclaves {
		return returnFalseOrTrue(false)
	}
	enclaves, visited, isEnclaves = bfs(i+1, j, grid, enclaves, visited)
	if !isEnclaves {
		return returnFalseOrTrue(false)
	}
	enclaves, visited, isEnclaves = bfs(i, j+1, grid, enclaves, visited)
	if !isEnclaves {
		return returnFalseOrTrue(false)
	}
	return returnFalseOrTrue(true)
}
