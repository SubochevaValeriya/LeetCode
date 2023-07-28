package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSpeedOnTime([]int{1, 1}, 1.0))

	fmt.Println(roundFloat(100000.0/6666667.0, 2))
	//	fmt.Println(minSpeedOnTime([]int{1, 3, 2}, 2.7))
	//fmt.Println(minSpeedOnTime([]int{1, 1}, 1.0))
}

func minSpeedOnTime(dist []int, hour float64) int {

	if len(dist)-1 > int(hour) {
		return -1
	}
	time := 0.0

	left, right := 1, 100000000
	speedTime := map[int]float64{}
	for left <= right {
		fmt.Println(speedTime)
		speed := left + (right-left)/2

		if speed >= 99999999 || speed < 1 {
			return -1
		}
		for i := 0; i < len(dist); i++ {
			add := float64(dist[i]) / float64(speed)
			if i != len(dist)-1 {
				addRounded := math.Round(add)
				if addRounded < add {
					add = addRounded + 1
				} else {
					add = addRounded
				}
			}
			time = time + roundFloat(add, 9)
			if time > hour {
				if lastTime, ok := speedTime[speed+1]; ok {
					if lastTime <= hour {
						return speed + 1
					}
				}
				speedTime[speed] = time
				left = speed
				time = 0
				break
			}
			if i == len(dist)-1 {
				if speed == 1 {
					return speed
				}
				speedTime[speed] = time
				if lastTime, ok := speedTime[speed-1]; ok {
					if lastTime > hour {
						return speed
					}
				}
				right = speed
				time = 0
			}
		}
	}

	return -1
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
