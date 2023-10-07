package main

import (
	"fmt"
	"regexp"
)

func main() {
	r, err := regexp.Match("^http.*+$", []byte("g"))
	fmt.Println(r, err)
	rooms := [][]int{{1}, {2}, {3}, {}}
	fmt.Println(canVisitAllRooms(rooms))
}

func canVisitAllRooms(rooms [][]int) bool {
	keys := map[int]bool{0: true}
	visited := map[int]bool{}
	for i := 0; i < len(rooms); i++ {
		bfs(rooms, keys, visited, i, 0)
	}
	if len(visited) == len(rooms) {
		return true
	}

	return false
}

func bfs(rooms [][]int, keys, visited map[int]bool, i int, countOpen int) {
	if visited[i] {
		return
	}
	if !keys[i] {
		return
	} else {
		countOpen++
	}
	visited[i] = true
	for _, key := range rooms[i] {
		keys[key] = true
		bfs(rooms, keys, visited, key, 0)
	}

	return
}
