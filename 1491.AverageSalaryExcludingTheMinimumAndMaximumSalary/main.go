package main

import "fmt"

// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/description/

func main() {
	fmt.Println(average([]int{48000, 59000, 99000, 13000, 78000, 45000, 31000, 17000, 39000, 37000, 93000, 77000, 33000, 28000, 4000, 54000, 67000, 6000, 1000, 11000}))
}

func average(salary []int) float64 {

	min, max := salary[0], salary[1]

	sum := 0

	for i := range salary {
		sum += salary[i]
		if salary[i] < min {
			min = salary[i]
		}
		if salary[i] > max {
			max = salary[i]
		}
	}

	sum = sum - (min + max)

	return float64(sum) / float64(len(salary)-2)

}
