package main

import "fmt"

func main() {
	n := 5
	lamps := [][]int{{0, 1}}
	queries := [][]int{{1, 0}, {1, 1}}
	fmt.Println(gridIllumination(n, lamps, queries))
}

func gridIllumination(n int, lamps [][]int, queries [][]int) []int {

	illuminated := map[[2]int]bool{}
	rowsIlluminated := map[int]int{}
	lampsM := map[[2]int]struct{}{}

	for _, lamp := range lamps {
		lampToTurnOn := [2]int{lamp[0], lamp[1]}
		lampsM[lampToTurnOn] = struct{}{}
		illuminated[lampToTurnOn] = 1
		rowsIlluminated[lamp[0]] = true
		illuminated = illuminateRows(n, illuminated, lamp[0])
		illuminated = illuminateColumns(n, illuminated, lamp[1])
		illuminated = illuminateDiagonal(n, illuminated, lamp[0], lamp[1])
	}
	fmt.Println(illuminated)
	answers := make([]int, len(queries))

	for i, query := range queries {
		lampToDetermine := [2]int{query[0], query[1]}
		answers[i] = illuminated[lampToDetermine]
		lampsM = turnOffAdjacent(n, lampsM, query[0], query[1])
		illuminated = map[[2]int]int{}
		for lamp := range lampsM {
			illuminated[lamp] = 1
			illuminated = illuminateRows(n, illuminated, lamp[0])
			illuminated = illuminateColumns(n, illuminated, lamp[1])
			illuminated = illuminateDiagonal(n, illuminated, lamp[0], lamp[1])
		}
	}

	return answers
}

func illuminateRows(n int, illuminated map[[2]int]int, row int) map[[2]int]int {
	for i := 0; i < n; i++ {
		lampToTurnOn := [2]int{row, i}
		illuminated[lampToTurnOn] = 1
	}
	return illuminated
}

func illuminateColumns(n int, illuminated map[[2]int]int, column int) map[[2]int]int {
	for i := 0; i < n; i++ {
		lampToTurnOn := [2]int{i, column}
		illuminated[lampToTurnOn] = 1
	}
	return illuminated
}

func illuminateDiagonal(n int, illuminated map[[2]int]int, row, column int) map[[2]int]int {
	for i := 0; row+i < n; i++ {
		lampToTurnOn := [2]int{row + i, column + i}
		illuminated[lampToTurnOn] = 1
	}
	for i := 0; row-i >= 0; i++ {
		lampToTurnOn := [2]int{row - i, column - i}
		illuminated[lampToTurnOn] = 1
	}
	for i := 0; row+i < n; i++ {
		lampToTurnOn := [2]int{row + i, column - i}
		illuminated[lampToTurnOn] = 1
	}

	for i := 0; row-i >= 0; i++ {
		lampToTurnOn := [2]int{row - i, column + i}
		illuminated[lampToTurnOn] = 1
	}

	return illuminated
}

func turnOffAdjacent(n int, lamps map[[2]int]struct{}, row, column int) map[[2]int]struct{} {
	for i := column - 1; i <= column+1; i++ {
		if i < 0 || i >= n {
			continue
		}
		lampToTurnOff := [2]int{row, i}
		delete(lamps, lampToTurnOff)
		if row-1 >= 0 {
			lampToTurnOff = [2]int{row - 1, i}
			delete(lamps, lampToTurnOff)
		}
		if row+1 < n {
			lampToTurnOff = [2]int{row + 1, i}
			delete(lamps, lampToTurnOff)
		}
	}

	return lamps
}
