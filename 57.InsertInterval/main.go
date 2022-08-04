package main

import "fmt"

func main() {

	//intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	//newInterval := []int{
	//	4, 8,
	//}

	//intervals := [][]int{{1, 5}}
	//newInterval := []int{
	//	2, 3,
	//}

	//intervals := [][]int{{1, 5}}
	//newInterval := []int{
	//	5, 7,
	//}

	intervals := [][]int{{2, 5}, {6, 7}}
	newInterval := []int{
		0, 1,
	}

	fmt.Println(insert(intervals, newInterval))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	leftPart := [][]int{}
	rightPart := [][]int{}
	newIntervalNew := newInterval

	for i, val := range intervals {
		fmt.Println(val[0], val[1])
		if val[0] > newInterval[1] {
			rightPart = append(rightPart, intervals[i:]...)
			break
		} else if val[1] < newInterval[0] {
			leftPart = append(leftPart, val)

		}

		if val[1] >= newInterval[0] && val[0] < newInterval[0] {
			newIntervalNew[0] = val[0]
		}
		if val[1] >= newInterval[1] && val[0] <= newInterval[1] {
			newIntervalNew[1] = val[1]
		}
	}

	resultIntervals := [][]int{}
	resultIntervals = append(leftPart, newInterval)
	resultIntervals = append(resultIntervals, rightPart...)
	//for i, val := range intervals {
	//	if val[0] > newInterval[0] {
	//		if val[0] > newInterval[1] {
	//			resultIntervals = append(resultIntervals, intervals[0:i]...)
	//			resultIntervals = append(resultIntervals, newInterval)
	//			resultIntervals = append(resultIntervals, intervals[i:]...)
	//			break
	//		}
	//		if val[1] <= newInterval[1] {
	//			intervals[i][0] = newInterval[0]
	//			resultIntervals = intervals
	//			break
	//		}
	//	}
	//	if val[0] < newInterval[0] {
	//		val
	//	}
	//
	//}
	return resultIntervals
}
