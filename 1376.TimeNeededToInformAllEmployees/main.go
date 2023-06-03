package main

import (
	"fmt"
)

func main() {
	//n := 15
	//headID := 0
	//manager := []int{-1, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6}
	//informTime := []int{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}

	n := 11
	headID := 4
	manager := []int{5, 9, 6, 10, -1, 8, 9, 1, 9, 3, 4}
	informTime := []int{0, 213, 0, 253, 686, 170, 975, 0, 261, 309, 337}
	fmt.Println(numOfMinutes(n, headID, manager, informTime))
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {

	m := map[int][]int{}
	for employee, man := range manager {
		if len(m[man]) == 0 {
			m[man] = []int{}
		}
		m[man] = append(m[man], employee)
	}

	return informTime[headID] + dfs(headID, m, informTime)
}

func dfs(head int, m map[int][]int, informTime []int) int {
	maxTimePerLevel := 0
	for _, man := range m[head] {
		time := dfs(man, m, informTime) + informTime[man]
		if time > maxTimePerLevel {
			maxTimePerLevel = time
		}
	}
	return maxTimePerLevel
}
