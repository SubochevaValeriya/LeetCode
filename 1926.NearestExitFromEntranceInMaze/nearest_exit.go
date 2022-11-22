package main

import (
	"fmt"
)

type treeNode struct {
	row    int
	column int
	up     *treeNode
	down   *treeNode
	left   *treeNode
	right  *treeNode
}

func main() {

	mazeStr := [][]string{{"+", "+", ".", "+"}, {".", ".", ".", "+"}, {"+", "+", "+", "."}}
	mazeStr2 := [][]string{{"+", "+", "+"}, {".", ".", "."}, {"+", "+", "+"}}
	mazeStr3 := [][]string{{".", "+"}}
	//maze := [][]byte{}

	//for i := 0; i < len(mazeStr); i++ {
	//	buf := &bytes.Buffer{}
	//	gob.NewEncoder(buf).Encode(mazeStr[i])
	//	bs := buf.Bytes()
	//	fmt.Println(bs)
	//	maze = append(maze, bs)
	//	strs2 := []string{}
	//	gob.NewDecoder(buf).Decode(&strs2)
	//
	//}

	entrance := []int{1, 2}
	entrance2 := []int{1, 0}
	entrance3 := []int{0, 0}

	fmt.Println(nearestExitTest(mazeStr, entrance))
	fmt.Println(nearestExitTest(mazeStr2, entrance2))
	fmt.Println(nearestExitTest(mazeStr3, entrance3))
}

func nearestExit(maze [][]byte, entrance []int) int {
	row, column := entrance[0], entrance[1]
	stepsRight, stepsLeft, stepsUp, stepsDown := 0, 0, 0, 0

	//bfsQueue := [][]int{}
	bfs := map[[2]int]int{}
	howFar := 1
	notFoundWay := false

	for {
		// right
		if string(maze[row][column+stepsRight]) == "." {
			coordinatesCheck := [2]int{row, column}
			if _, ok := bfs[coordinatesCheck]; ok {
				coordinates := [2]int{row, column + stepsRight}
				bfs[coordinates] += 1
				if isExit(maze, row, column+howFar, entrance) {
					return bfs[coordinates]
				}
			}
		} else {
			notFoundWay = true
		}

		// left
		if string(maze[row][column-stepsLeft]) == "." {
			coordinatesCheck := [2]int{row, column}
			if _, ok := bfs[coordinatesCheck]; ok {
				coordinates := [2]int{row, column - stepsLeft}
				bfs[coordinates] += 1
				if isExit(maze, row, column-howFar, entrance) {
					return bfs[coordinates]
				}
			}
		} else {
			notFoundWay = true
		}

		// up
		if string(maze[row+stepsUp][column]) == "." {
			coordinatesCheck := [2]int{row, column}
			if _, ok := bfs[coordinatesCheck]; ok {
				coordinates := [2]int{row + stepsUp, column}
				bfs[coordinates] += 1
				if isExit(maze, row+stepsUp, column, entrance) {
					return bfs[coordinates]
				}
			}
		} else {
			notFoundWay = true
		}

		// down
		if string(maze[row-stepsDown][column]) == "." {
			coordinatesCheck := [2]int{row, column}
			if _, ok := bfs[coordinatesCheck]; ok {
				coordinates := [2]int{row - stepsDown, column}
				bfs[coordinates] += 1
				if isExit(maze, row-stepsUp, column, entrance) {
					return bfs[coordinates]
				}
			}
		} else {
			notFoundWay = true
		}

		if notFoundWay == true {
			return -1
		}

		notFoundWay = false

		if column+stepsRight < len(maze[1])-1 {
			stepsRight++
		}
		if column-stepsLeft > 0 {
			stepsLeft--
		}

		if row+stepsUp < len(maze)-1 {
			stepsRight++
		}
		if column-stepsDown > 0 {
			stepsLeft--
		}

	}
}

func isExit(maze [][]byte, i int, j int, entrance []int) bool {
	if i == 0 || i == len(maze)-1 || j == 0 || j == len(maze[i])-1 {
		if entrance[0] != i && entrance[1] != j {
			return true
		}
	}

	return false
}

//
func nearestExitTestFirst(maze [][]string, entrance []int) int {
	row, column := entrance[0], entrance[1]
	stepsRight, stepsLeft, stepsUp, stepsDown := 0, 0, 0, 0
	bfs := map[[2]int]int{}
	notFoundWay, cantChange := true, true
	coordinates := [2]int{row, column}
	bfs[coordinates] = 0
	step := 0

	for {
		// right
		if stepsLeft == 0 && stepsRight == 0 && stepsUp == 0 && stepsDown == 0 {
			step = 0
		} else {
			step = 1
		}

		if string(maze[row][column+stepsRight]) == "." && stepsRight != 0 {
			coordinatesCheck := [2]int{row, column + stepsRight - step}
			if _, ok := bfs[coordinatesCheck]; ok {
				coordinates = [2]int{row, column + stepsRight}
				bfs[coordinates] = bfs[coordinatesCheck] + 1
				notFoundWay = false
			}
			if isExitTest(maze, row, column+stepsRight, entrance) {
				return bfs[coordinates]
			}
		}

		// left
		if string(maze[row][column-stepsLeft]) == "." && stepsLeft != 0 {
			coordinatesCheck := [2]int{row, column - stepsLeft + step}
			if _, ok := bfs[coordinatesCheck]; ok {
				coordinates = [2]int{row, column - stepsLeft}
				bfs[coordinates] = bfs[coordinatesCheck] + 1
				notFoundWay = false
				if isExitTest(maze, row, column-stepsLeft, entrance) {
					return bfs[coordinates]
				}
			}
		}

		// up
		if string(maze[row+stepsUp][column]) == "." && stepsUp != 0 {
			coordinatesCheck := [2]int{row + stepsUp - step, column}
			if _, ok := bfs[coordinatesCheck]; ok {
				coordinates = [2]int{row + stepsUp, column}
				bfs[coordinates] = bfs[coordinatesCheck] + 1
				notFoundWay = false
				if isExitTest(maze, row+stepsUp, column, entrance) {
					return bfs[coordinates]
				}
			}
		}

		// down
		if string(maze[row-stepsDown][column]) == "." && stepsDown != 0 {
			coordinatesCheck := [2]int{row - stepsDown + step, column}
			if _, ok := bfs[coordinatesCheck]; ok {
				coordinates = [2]int{row - stepsDown, column}
				bfs[coordinates] = bfs[coordinatesCheck] + 1
				notFoundWay = false
				if isExitTest(maze, row-stepsDown, column, entrance) {
					return bfs[coordinates]
				}
			}
		}

		if column+stepsRight < len(maze[0])-1 {
			stepsRight++
			cantChange = false
		}

		if column-stepsLeft > 0 {
			stepsLeft++
			cantChange = false
		}

		if row+stepsUp < len(maze)-1 {
			stepsUp++
			cantChange = false
		}

		if row-stepsDown > 0 {
			stepsDown++
			cantChange = false
		}

		// no exit condition
		if notFoundWay == true && cantChange == true {
			return -1
		}

		notFoundWay = true
		cantChange = true
	}
}

func nearestExitTest(maze [][]string, entrance []int) int {
	row, column := entrance[0], entrance[1]
	queue := [][2]int{}
	coordinates := [2]int{row, column}
	queue = append(queue, coordinates)
	visited := map[[2]int]int{}
	visited[coordinates]++
	ways := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for i := 0; i < len(queue); i++ {
		fmt.Println(row, column)
		row, column = queue[i][0], queue[i][1]
		for w := 0; w < len(ways); w++ {
			coordinates = [2]int{row + ways[w][0], column + ways[w][1]}
			if _, ok := visited[coordinates]; ok {
				continue
			}
			if row+ways[w][0] < len(maze) && row+ways[w][0] >= 0 && column+ways[w][1] >= 0 && column+ways[w][1] < len(maze[0]) {
				visited[coordinates] = visited[[2]int{row, column}] + 1
				if string(maze[row+ways[w][0]][column+ways[w][1]]) == "." {
					if isExitTest(maze, row+ways[w][0], column+ways[w][1], entrance) {
						return visited[coordinates] - 1
					}
					coordinates = [2]int{row + ways[w][0], column + ways[w][1]}
					queue = append(queue, coordinates)
				}
			}
		}
	}
	return -1
}

func isExitTest(maze [][]string, i int, j int, entrance []int) bool {
	if i == 0 || i == len(maze)-1 || j == 0 || j == len(maze[i])-1 {
		if entrance[0] == i && entrance[1] == j {
			return false
		} else {
			return true
		}
	}

	return false
}
