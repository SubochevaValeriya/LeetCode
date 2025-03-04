package main

func main() {

}

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	isOccupied := map[[2]int]bool{}
	guardedCells := map[[2]int]bool{}

	for _, position := range guards {
		isOccupied[[2]int{position[0], position[1]}] = true
	}

	for _, position := range walls {
		isOccupied[[2]int{position[0], position[1]}] = true
	}

	for _, position := range guards {
		for i := position[0] + 1; i < m; i++ {
			cell := [2]int{i, position[1]}
			if isOccupied[cell] {
				break
			}
			guardedCells = addGuardedCell(guardedCells, cell)
		}
		for i := position[0] - 1; i >= 0; i-- {
			cell := [2]int{i, position[1]}
			if isOccupied[cell] {
				break
			}
			guardedCells = addGuardedCell(guardedCells, cell)
		}
		for i := position[1] + 1; i < n; i++ {
			cell := [2]int{position[0], i}
			if isOccupied[cell] {
				break
			}
			guardedCells = addGuardedCell(guardedCells, cell)
		}
		for i := position[1] - 1; i >= 0; i-- {
			cell := [2]int{position[0], i}
			if isOccupied[cell] {
				break
			}
			guardedCells = addGuardedCell(guardedCells, cell)
		}

	}

	return m*n - len(guardedCells) - len(walls) - len(guards)
}

func addGuardedCell(guardedCells map[[2]int]bool, cell [2]int) map[[2]int]bool {
	if guardedCells[cell] {
		return guardedCells
	}

	guardedCells[cell] = true

	return guardedCells
}
