package main

import "fmt"

func main() {
	//numCourses := 20
	//prerequisites := [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}}

	//numCourses := 3
	//prerequisites := [][]int{{0, 1}, {0, 2}, {1, 2}}

	//numCourses := 4
	//prerequisites := [][]int{{2, 0}, {1, 0}, {3, 1}, {3, 2}, {1, 3}}

	numCourses := 8
	prerequisites := [][]int{{1, 0}, {2, 6}, {1, 7}, {6, 4}, {7, 0}, {0, 5}}

	fmt.Println(canFinish(numCourses, prerequisites))
}

func canFinish(numCourses int, prerequisites [][]int) bool {

	if numCourses <= 1 {
		return true
	}
	coursesNeeded := map[int][]int{}

	for _, prerequisite := range prerequisites {
		if prerequisite[0] == prerequisite[1] {
			return false
		}
		coursesNeeded[prerequisite[0]] = append(coursesNeeded[prerequisite[0]], prerequisite[1])
	}

	for course, _ := range coursesNeeded {
		canTake, _, _ := dfs(course, coursesNeeded, map[int]bool{}, map[int]bool{})
		if !canTake {
			return canTake
		}
	}
	return true
}

func dfs(course int, coursesNeeded map[int][]int, visited map[int]bool, courseTaken map[int]bool) (bool, map[int]bool, map[int]bool) {

	if visited[course] {
		return false, visited, courseTaken
	}

	visited[course] = true
	canTake := true

	for _, courseNeeded := range coursesNeeded[course] {
		if !courseTaken[courseNeeded] {
			if _, ok := coursesNeeded[courseNeeded]; !ok {
				continue
			}
			canTake, visited, courseTaken = dfs(courseNeeded, coursesNeeded, visited, courseTaken)
			if !canTake {
				return false, visited, courseTaken
			}
		}
	}
	courseTaken[course] = true
	return canTake, visited, courseTaken
}
