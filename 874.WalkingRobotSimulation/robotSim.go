package main

import "fmt"

func main() {
	//fmt.Println(robotSim([]int{4, -1, 3}, [][]int{}))
	//fmt.Println(robotSim([]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}))
	fmt.Println(robotSim([]int{6, -1, -1, 6}, [][]int{}))
}

func robotSim(commands []int, obstacles [][]int) int {
	currentLocation := [2]int{0, 0}
	direction := 1
	maxDistance := 0
	obstaclesMap := map[[2]int]struct{}{}

	/*	directionMap := map[int]string{
		1: "north",
		2: "east",
		3: "south",
		4: "west",
	}*/

	for _, obstacle := range obstacles {
		obstacleForMap := [2]int{obstacle[0], obstacle[1]}
		obstaclesMap[obstacleForMap] = struct{}{}
	}

	for _, command := range commands {
		var movingX, movingY int
		switch command {
		case -1:
			switch {
			case direction+1 > 4:
				direction = 1
			default:
				direction += 1
			}
		case -2:
			switch {
			case direction-1 < 1:
				direction = 4
			default:
				direction -= 1
			}
		default:
			switch direction {
			case 1:
				movingY = 1
			case 2:
				movingX = 1
			case 3:
				movingY = -1
			case 4:
				movingX = -1
			}

			for i := command; i > 0; i-- {
				newLocation := [2]int{currentLocation[0] + 1*movingX, currentLocation[1] + 1*movingY}
				for {
					if _, blockedByObst := obstaclesMap[newLocation]; !blockedByObst {
						currentLocation = newLocation
						break
					}
					newLocation = [2]int{newLocation[0] - 1*movingX, newLocation[1] - 1*movingY}
				}
			}
			if sqrt(currentLocation[0])+sqrt(currentLocation[1]) > maxDistance {
				maxDistance = sqrt(currentLocation[0]) + sqrt(currentLocation[1])
			}
		}
	}
	return maxDistance
}

func sqrt(a int) int {
	return a * a
}
