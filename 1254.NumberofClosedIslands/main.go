package main

import "fmt"

func main() {

	//grid := [][]int{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}}
	grid := [][]int{{0, 0, 1, 1, 0, 1, 0, 0, 1, 0}, {1, 1, 0, 1, 1, 0, 1, 1, 1, 0}, {1, 0, 1, 1, 1, 0, 0, 1, 1, 0}, {0, 1, 1, 0, 0, 0, 0, 1, 0, 1}, {0, 0, 0, 0, 0, 0, 1, 1, 1, 0}, {0, 1, 0, 1, 0, 1, 0, 1, 1, 1}, {1, 0, 1, 0, 1, 1, 0, 0, 0, 1}, {1, 1, 1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 0, 0, 1, 0, 1, 0, 1}, {1, 1, 1, 0, 1, 1, 0, 1, 1, 0}}

	for _, ints := range grid {
		fmt.Println(ints)
	}

	fmt.Println(closedIsland(grid))
	//	closed := map[[2]int]bool{}
	//	visited := map[[2]int]struct{}{}
	//fmt.Println(bfs(4, 3, grid, closed, visited))
}

func closedIsland(grid [][]int) int {

	closed := map[[2]int]bool{}
	visited := map[[2]int]struct{}{}
	answer := 0
	var isClosed bool
	for i := range grid {
		for j := range grid[i] {
			_, ok := visited[[2]int{i, j}]
			if ok {
				continue
			}
			if grid[i][j] == 1 {
				closed[[2]int{i, j}] = true
				continue
			}
			if i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[i])-1 {
				closed[[2]int{i, j}] = false
				continue
			}
			closed, visited, isClosed = bfs(i, j, grid, closed, visited)
			if isClosed {
				answer++
			}
		}
	}
	return answer
}

func bfs(i int, j int, grid [][]int, closed map[[2]int]bool, visited map[[2]int]struct{}) (map[[2]int]bool, map[[2]int]struct{}, bool) {
	isClosed, ok := closed[[2]int{i, j}]
	if ok {
		return closed, visited, isClosed
	}
	_, ok = visited[[2]int{i, j}]
	if ok {
		return closed, visited, true
	}
	visited[[2]int{i, j}] = struct{}{}

	returnFalseOrTrue := func(falseOrTrue bool) (map[[2]int]bool, map[[2]int]struct{}, bool) {
		closed[[2]int{i, j}] = falseOrTrue
		return closed, visited, falseOrTrue
	}

	if grid[i][j] == 1 {
		return returnFalseOrTrue(true)
	}

	if i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[i])-1 {
		return returnFalseOrTrue(false)
	}
	closed, visited, isClosed = bfs(i-1, j, grid, closed, visited)
	if !isClosed {
		return returnFalseOrTrue(false)
	}
	closed, visited, isClosed = bfs(i, j-1, grid, closed, visited)
	if !isClosed {
		return returnFalseOrTrue(false)
	}
	closed, visited, isClosed = bfs(i+1, j, grid, closed, visited)
	if !isClosed {
		return returnFalseOrTrue(false)
	}
	closed, visited, isClosed = bfs(i, j+1, grid, closed, visited)
	if !isClosed {
		return returnFalseOrTrue(false)
	}

	return returnFalseOrTrue(true)
}

/*closed, isClosed = bfs(i, j-1, grid, closed)
if isClosed == false{
closed[[2]int{i, j}] = false
return closed, false
}
closed, isClosed = bfs(i, j+1, grid, closed)
if isClosed == false{
closed[[2]int{i, j}] = false
return closed, false
}
closed, isClosed = bfs(i-1, j, grid, closed)
if isClosed == false{
closed[[2]int{i, j}] = false
return closed, false
}
closed, isClosed = bfs(i+1, j, grid, closed)
if isClosed == false{
closed[[2]int{i, j}] = false
return closed, false
}*/
