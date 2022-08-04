package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	initial := []int{3, 2, 3}
	target := 6
	fmt.Println(twoSumTwoFor(initial, target))
	fmt.Println(twoSumMap(initial, target))
}

func twoSumTwoFor(nums []int, target int) []int {
	var result int
	var answer []int
	for i, integer1 := range nums {
		if len(answer) == 2 {
			break
		} else {
			result = target - integer1
			for j, integer2 := range nums {
				if result == integer2 && i != j {
					answer = append(answer, i, j)
				}
			}
		}

	}
	return answer
}

func twoSumMap(nums []int, target int) []int {
	var result int
	m := make(map[int]int)
	var answer []int
	for i, integer1 := range nums {
		_, found := m[integer1]
		if found == true {
			answer = append(answer, i, m[integer1])
			return answer
		} else if found == false {
			result = target - integer1
			m[result] = i
		}
	}
	return nil
}

func values() ([]int, int) {
	initial := make([]int, 10e4)
	rand.Seed(time.Now().UnixNano())
	var max int = 10e9
	var min int = -10e9
	for i := range initial {
		initial[i] = rand.Intn(max-min) + min
	}
	target := rand.Intn(max-min) + min
	return initial, target
}
